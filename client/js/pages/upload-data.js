
function parseFile() {  
    const file = $$('id-file-input')?.files[0];
    const reader = new FileReader();
    reader.onload = onLoadFile;
    reader.readAsText(file);
    changeSection('id-split-to-categories-section')
}

function onLoadFile(e) {
    globalTransArray = JSON.parse(e.target.result).result.transactions;
    displayCurrentTransaction();
    displayCategories();
}
