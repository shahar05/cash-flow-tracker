var globalTransArray = null;
var globalTransArrayIndex = 0;

async function getClassifyTrans(fullTrans) {
    const res = await http.call("filter-transactions", "POST", {trans:fullTrans});
    if (res.error) throw Error(res.error);
    globalTransArray = res.data
}

async function attachTransaction(catName,catID) {
    const cat = {name: catName,id: catID}
    const body = createNewTransaction(getCurrentTransaction(), cat)

    const response = await http.call("transactions", "POST", body);
    console.log(response);
    displayNextTransaction();
}

function displayCurrentTransaction() {
    if (!globalTransArray || !globalTransArray.length) {
        console.error('displayCurrentTransaction: globalTransArray is null');
        showAnalysesPage();
        return;
    }

    const trans = getCurrentTransaction();
    $$('id-trans-merchant-name').innerHTML = trans.merchantName;
    $$('id-trans-display-amount').innerHTML = trans.amountForDisplay;
    $$('id-trans-debt-date').innerHTML = trans.trnPurchaseDate;

}

function getAnalyses() {
    return http.call('analysis')
}

 function showAnalysesPage() {
    const pageId = 'id-view-analyses-section';
    changeSection(pageId, async ()=> {
        const res = await getAnalyses();             
        google.charts.load('current', {'packages':['corechart']});
        google.charts.setOnLoadCallback( () =>  {
            drawChart(res.data)  
        } );
    });
}

function makeArray(data) {    
    const arr = [['Category', 'Sum']];
    for (const item of data) {
        arr.push([item.name, item.sum])
    }
    return arr;
}

function drawChart(data) {
    const totalSum = data.reduce((total, item) => total + item.sum, 0);
    data = google.visualization.arrayToDataTable(makeArray(data));

    const options = { 
        pieHole: 0.4, // For a donut chart effect
        pieSliceText: 'value',
        width: 1200,
        height: 1200,
        chartArea: {
            top: 0,    // Adjust top margin (lower value moves the entire chart area up)
            left: 0,    // Adjust left margin
            width: '100%',
            height: '80%' // Increase this for a larger pie chart
        },
        legend: {
            position: 'right',  // Place the legend beside the pie chart
            alignment: 'center', // Vertically align the legend in the center
        },
        pieSliceTextStyle: {
            fontSize: 12,  // Control text size
        }
    };

    const chart = new google.visualization.PieChart($$('id-view-analyses-container'));


    // Add text in the center after drawing the chart
    google.visualization.events.addListener(chart, 'ready', function() {
    
        const container = document.getElementById('id-view-analyses-container');
        
        // Remove any existing text to avoid multiple text elements
        const existingText = document.getElementById('centerText');
        if (existingText) {
            existingText.remove();
        }

        // Create a text element or use innerHTML to place the text in the center
        const centerText = document.createElement('div');
        centerText.id = 'centerText';  // Unique ID for the text element
        centerText.style.position = 'absolute';
        

        centerText.style.left = '45%';
        centerText.style.top = '50%';

        centerText.style.fontSize = '20px';
        centerText.style.fontWeight = 'bold';
        centerText.innerHTML = 'Total<br>' + Math.floor(totalSum);

        // Append text to the chart container
        container.appendChild(centerText);
    });
    

    chart.draw(data, options);


    // Add click event listener
    google.visualization.events.addListener(chart, 'select', function() {
        var selectedItem = chart.getSelection()[0];
        if (selectedItem) {
            const category = data.getValue(selectedItem.row, 0); // Get category
            const sum = data.getValue(selectedItem.row, 1); // Get value (sum)
            clickedOnPieChartItem(category, sum);
        }
    });
}


// function drawChart(data) {

//     data = google.visualization.arrayToDataTable(makeArray(data));

//     const options = { 
//         pieHole: 0.4, // For a donut chart effect
//         pieSliceText: 'value',
//         width: 1200,
//         height: 1200,
//         chartArea: {
//             top: 0,    // Adjust top margin (lower value moves the entire chart area up)
//             left: 150,    // Adjust left margin
//             width: '100%',
//             height: '80%' // Increase this for a larger pie chart
//         },
//         legend: {
//             position: 'right',  // Place the legend beside the pie chart
//             alignment: 'center', // Vertically align the legend in the center

//         },
//         pieSliceTextStyle: {
//             fontSize: 12,  // Control text size
//         }
//     };

//     const chart = new google.visualization.PieChart($$('id-view-analyses-container'));

//     chart.draw(data, options);

//     // Add click event listener
//     google.visualization.events.addListener(chart, 'select', function() {
//         var selectedItem = chart.getSelection()[0];
//         if (selectedItem) {
//             const category = data.getValue(selectedItem.row, 0); // Get category
//             const sum = data.getValue(selectedItem.row, 1); // Get value (sum)
//             clickedOnPieChartItem(category, sum);
//         }
//     });

// }

var LastCategory = null;

//console.log('Clicked on:', category, 'with value:', sum);
async function clickedOnPieChartItem(category, sum) {
    
    if (!category) {
        category = LastCategory;
    }
    LastCategory = category;
    if (!category) {
        alert('wtf')
        return;
    }
    const resDetails = await http.call(`category-analysis?name=${category}`)
    if (resDetails.error) {
        console.error(resDetails.error);
        return;
    }
    
    const resGraph = await http.call(`category-graph?name=${category}`)
    if (resGraph.error) {
        console.error(resGraph.error);
        return;
    }

    // Make API CALL For get the GRaph
    changeSection('id-category-drill-down-section',()=>{   
             
        $$('id-category-drill-down-section-a').removeAttribute('disabled');
        $$('id-category-detailed-title').innerHTML = strings.capitalizeFirstLetter(category)
        $$('id-list-of-category').innerHTML = buildItemList(resDetails.data);
        showGraph(resGraph.data)
    });
}

function showGraph(data) {
    google.charts.load('current', {'packages':['corechart']});

    // Set a callback function to run when the package is loaded
    google.charts.setOnLoadCallback(() => {
      drawGraph(data);
    });
}

      // Function to convert your data object and draw the chart
      function drawGraph(dataObject) {
        // Convert the dataObject into the format that Google Charts expects
        var data = new google.visualization.DataTable();
        data.addColumn('date', 'Month'); // First column: Date
        data.addColumn('number', 'Total'); // Second column: Total

        // Loop through the dataObject and add rows to the chart data
        dataObject.forEach(item => {
          data.addRow([new Date(item.Month), item.Total]);
        });

        // Set chart options
        var options = {
          title: 'Monthly Totals',
          curveType: 'function', // Makes the line smooth
          legend: { position: 'bottom' },
          hAxis: {
            title: 'Month',
            format: 'MMM yyyy' // Custom format for dates
          },
          vAxis: {
            title: 'Total'
          }
        };

        // Create the chart and draw it in the 'chart_div' element
        var chart = new google.visualization.LineChart($$('id-graph-category'));
        chart.draw(data, options);
      }

function toggleCardDetails(card) {
    card.classList.toggle('open');
}

function buildItemList(transArray) {
    let res = "";
    for (const trans of transArray) {
        res += createCard(trans);
    }
    return res;
}

function createCard(trans) {

    return `
               <div class="card" onclick="toggleCardDetails(this)">
                        <div class="card-header">
                            <div class="card-title">
                                <h3>${trans.name}</h3>
                                <p>₪${trans.amount}</p>
                            </div>
                            <span class="arrow">▲</span>
                        </div>
                        <div class="card-details">
                      <div class="card-details">
                        <p><strong>External ID:</strong> ${trans.external_id}</p>
                        <p><strong>Date:</strong> ${trans.date}</p>
                        <p><strong>Address:</strong> ${trans.address}</p>
                        <p><strong>Card Unique ID:</strong> ${trans.card_unique_id}</p>
                        <p><strong>Merchant Phone No:</strong> ${trans.merchant_phone_no}</p>
                        <p><strong>International Branch ID:</strong> ${trans.international_branch_id}</p>
                     </div>

                        </div>
                </div>
         `
    
}

function displayNextTransaction() {
    const finishedTransactions = incrementTransIndex();
    if (finishedTransactions) {
        showAnalysesPage();
        return;
    }
    displayCurrentTransaction();
}

function getCurrentTransaction() {
    if (!globalTransArray || !globalTransArray.length) {
        console.error('getCurrentTransaction: globalTransArray is null');
        return;
    }
    return globalTransArray[globalTransArrayIndex];
}

function incrementTransIndex() {
    if (globalTransArrayIndex === globalTransArray.length - 1) {
        alert("no more array");
        console.error('displayCurrentTransaction: globalTransArrayIndex >= globalTransArray.length');
        return true;
    }
    globalTransArrayIndex++;
    return false;
}

function createNewTransaction(transaction, category) {
    return {        
        external_id: transaction.trnIntId, // Main Transaction ID
        name: transaction.merchantName, // merchant name
        amount: transaction.amountForDisplay, // transaction amount
        date_str: transaction.trnPurchaseDate, // date string
        date: new Date(transaction.trnPurchaseDate), // Date object
        address: transaction.merchantAddress, // merchant address
        card_unique_id: transaction.cardUniqueId, // card unique ID
        category,                                // transaction category
        merchant_phone_no: transaction.merchantPhoneNo, // merchant phone number
        international_branch_id: transaction.internationalBranchID, // international branch ID
    };
}