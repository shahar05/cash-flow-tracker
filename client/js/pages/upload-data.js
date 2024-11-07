// upload-file.js

var UploadFilePage = (()=>{

    function parseFile() {          
        const reader = new FileReader();
        reader.onload = onLoadFile;
        reader.readAsText($$('id-file-input')?.files[0]);        
    }

    async function onLoadFile(e) {
        const unfilteredTrans = JSON.parse(e.target.result).result.transactions
        debugger;
        NavService.changeSection(null, 'id-attach-trans-with-category', {unfilteredTrans})          
    }

    function onInit() {
        
    }

    return {
        onInit,
        parseFile
    }

})()






