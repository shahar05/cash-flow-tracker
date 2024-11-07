const transactionService = (()=>{
    let globalTransArray = null;
    let globalTransArrayIndex = 0;


    function createTransaction(category, transaction) {
        if (!transaction) {
            transaction = getCurrentTransaction();
        }
        return {        
            external_id: transaction.trnIntId, // Main Transaction ID
            name: transaction.merchantName, // merchant name
            amount: transaction.amountForDisplay, // transaction amount
            date_str: transaction.trnPurchaseDate, // date string
            date: new Date(transaction.trnPurchaseDate), // Date object
            address: transaction.merchantAddress, // merchant address
            card_unique_id: transaction.cardUniqueId, // card unique ID
            category,                                // transaction category
            merchant_phone_no: transaction.merchantPhoneNo, // merchant phone number
            international_branch_id: transaction.internationalBranchID, // international branch ID
        };
    }

    async function filterTransactions(trans) {
        const res = await http.call("filter-transactions", "POST", {trans});
        if (res.error) {
            return res.error; 
        }
        globalTransArray = res.data;
        return null;
    }

    function getCurrentTransaction() {
        if (!globalTransArray || !globalTransArray.length) {
            console.error('getCurrentTransaction: globalTransArray is null');
            return null;
        }
        return globalTransArray[globalTransArrayIndex];
    }

    return {
        filterTransactions,
        getCurrentTransaction,
        createTransaction
    }

})()