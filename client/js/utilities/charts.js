const charts = (()=>{
    var LastCategory = null;

    function drawGraph(){

    }

    function drawPie(){

    }

    function makeArray(data) {    
        const arr = [['Category', 'Sum']];
        for (const item of data) {
            let sum = item.sum;
            if (sum > 0) {
                arr.push([item.name, sum])
            }
        }
        return arr;
    }

    function tempDrawPieChart(data) {
        google.charts.load('current', {'packages':['corechart']});
        google.charts.setOnLoadCallback( () =>  {
            const totalSum = data.reduce((total, item) => total + item.sum, 0);
            const totalCredit = data.reduce((total, item) => {
                if (item.sum < 0) {
                    total += item.sum;
                }
                return total;
            }, 0);
            console.log(totalCredit);
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
        });
    }

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

    return {
        tempDrawPieChart
    }

})()