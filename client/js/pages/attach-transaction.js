// attach-transaction.js

var AttachTransactionPage = (()=>{

    async function onInit(object) {
        const err = await transactionService.filterTransactions(object.unfilteredTarns);
        if (err) {            
            console.error(err); // Handle it
            return;
        }
        displayCurrentTransaction();
        displayCategories();
    }

    async function attachTransaction(name,id) {
        const cat = {name, id};
        const body = transactionService.createTransaction(cat);    
        const response = await http.call("transactions", "POST", body);
        console.log(response);
        displayNextTransaction();
    }

    function incrementTransIndex() {
        if (globalTransArrayIndex === globalTransArray.length - 1) {            
            console.error('displayCurrentTransaction: globalTransArrayIndex >= globalTransArray.length');
            return true;
        }
        globalTransArrayIndex++;
        return false;
    }

    function displayNextTransaction() {
        const finishedTransactions = incrementTransIndex();
        if (finishedTransactions) {
            // Move to Next Page: showAnalysesPage();
            return;
        }
        displayCurrentTransaction();
    }

    async function displayCategories() {
        const categories = await categoryService.getAllCategories();

        let catStr = '';
        for (const cat of categories) {
            catStr += `
                 <div  onclick="AttachTransactionPage.attachTransaction('${cat.name}','${cat.id}')" class="c-card c-pointer" id="id-category-${cat.id}" >
                    ${cat.name}
                </div>
            `;
        }
        $$('id-category-list-container').innerHTML = catStr;
    }

    function displayCurrentTransaction() {
        const trans = transactionService.getCurrentTransaction();

        if (!trans) {
            console.error("displayCurrentTransaction: Error getting trans"); // Handle Error!            
            return;
        }

        $$('id-trans-merchant-name').innerHTML = trans.merchantName;
        $$('id-trans-display-amount').innerHTML = trans.amountForDisplay;
        $$('id-trans-debt-date').innerHTML = trans.trnPurchaseDate;
    }

    return {
        onInit,
        attachTransaction
    }

})()

