// attach-transaction.js

var AttachTransactionPage = (()=>{

    async function onInit(object) {
        startDisplayTransactions(object?.unfilteredTrans);
        displayCategories();
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

    async function startDisplayTransactions(unfilteredTrans) {
        if (!unfilteredTrans) {
            console.error(`startDisplayTransactions: Didn't received any unfilteredTrans`); // TODO: Handle Error
            return; 
        }

        const err = await transactionService.filterTransactions(unfilteredTrans);
        if (err) {            
            console.error(err); // TODO: Handle Error
            return;
        }

        const currentTransaction = transactionService.getCurrentTransaction();
        displayTransaction(currentTransaction);
    }

    async function attachTransaction(name,id) {
        const cat = {name, id};
        const body = transactionService.createTransaction(cat);    
        await http.call("transactions", "POST", body); // Attach the Transaction

        // Display Next Transaction
        const nextTrans = transactionService.getNextTransaction()
        displayTransaction(nextTrans);
    }

    function displayTransaction(trans) {
        if (!trans) { // End of transactions move to Pie chart Page
            // Add some Animation
            NavService.changeSection(null, 'id-view-analyses');
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

