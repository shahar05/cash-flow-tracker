
function parseFile() {  
    const file = $$('id-file-input')?.files[0];
    const reader = new FileReader();
    reader.onload = onLoadFile;
    reader.readAsText(file);
    changeSection('id-split-to-categories-section')
}

async function onLoadFile(e) {
    const parseTrans = JSON.parse(e.target.result).result.transactions
    await getClassifyTrans(parseTrans);
    displayCurrentTransaction();
    displayCategories();
}
