// category-analysis.js

var CategoryAnalysisPage = (()=>{

    async function onInit() {
        const res = await analysisService.getAnalyses();    
        if (res.error) {
            console.error(res);
            return;
        }         
        google.charts.load('current', {'packages':['corechart']});
        google.charts.setOnLoadCallback( () =>  {
            drawChart(res.data)  
        } );
    }


    function drawChart(params) {
        
    }

    return {
        onInit
    }

})()

