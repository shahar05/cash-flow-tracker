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
        
        globalTransArrayIndex = 0; // Reset index when new data is assigned
        globalTransArray = res.data;
        return null;
    }

    function getNextTransaction() {
        ++globalTransArrayIndex;   
        return getCurrentTransaction();
    }

    function getCurrentTransaction() {
        if (!globalTransArray?.length || globalTransArrayIndex === globalTransArray.length) {
            // End of Array
            return null;
        }
        
        if (globalTransArrayIndex > globalTransArray.length) {
            console.error("getCurrentTransaction: This not should happen - Logic Error");
            return null;
        }
        return globalTransArray[globalTransArrayIndex];
    }

    return {
        filterTransactions,
        getCurrentTransaction,
        getNextTransaction,
        createTransaction
    }

})()