// category-analysis.js

var CategoryAnalysisPage = (()=>{

    async function onInit() {
        const res = await analysisService.getAnalyses();    
        if (res.error) {
            console.error(res);
            return;
        }         
        charts.tempDrawPieChart(res.data);
    }

    return {
        onInit
    }

})()

