const analysisService = (()=>{

    // This analysis is a pie chart off all month 
    function getAnalyses() {
        return http.call('analysis')
    }

    return {
        getAnalyses
    }

})()