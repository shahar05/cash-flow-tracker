

async function moveToExpensesGraph() {
    changeSection('id-graph-total-expenses');
    const res = await http.call('categories-graph');
    
    if (res.error) {
        console.error(res.error);
        return;
    }

    showExpensesGraph(res.data)

}


function showExpensesGraph(data){
    google.charts.load('current', {'packages':['corechart']});

    // Set a callback function to run when the package is loaded
    google.charts.setOnLoadCallback(() => {
        drawExpensesGraph(data);
    });
}

function drawExpensesGraph(dataObject) {
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
      var chart = new google.visualization.LineChart($$('id-graph-total-expenses-container'));
      chart.draw(data, options);
}


async function moveToPieByMonth() {
  changeSection('id-pie-chart-per-month')  
  const res = await http.call('monthly-analysis')
  if (res.error) {
    console.error(res.error);
    return 
  }
  google.charts.load('current', {'packages':['corechart']});
  google.charts.setOnLoadCallback( () =>  {
      drawMonthlyChart(res.data[0].categorySum)  
  } );
}


function drawMonthlyChart(data) {
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

  const chart = new google.visualization.PieChart($$('id-pie-chart-per-month-container'));


  // Add text in the center after drawing the chart
  google.visualization.events.addListener(chart, 'ready', function() {
  
      const container = document.getElementById('id-pie-chart-per-month-container');
      
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