var globalTransArray = null;
var globalTransArrayIndex = 0;

async function attachTransaction(catName,catID) {
    const cat = {name: catName,id: catID}
    const body = transformTransaction(getCurrentTransaction(), cat)

    const response = await sendHttpRequest("transactions", "POST", body);

    debugger;
    console.log(response);

    displayNextTransaction();
}

function displayCurrentTransaction() {
    if (!globalTransArray || !globalTransArray.length) {
        console.error('displayCurrentTransaction: globalTransArray is null');
        return;
    }

    const trans = getCurrentTransaction();
    $$('id-trans-merchant-name').innerHTML = trans.merchantName;
    $$('id-trans-display-amount').innerHTML = trans.amountForDisplay;
    $$('id-trans-debt-date').innerHTML = trans.trnPurchaseDate;

    if (alreadyAttached(trans)) {
        console.log(`Attached: ${trans.merchantName}. CategoryID: ${alreadyAttached(trans)}`);
    }
}

function alreadyAttached(transaction) {
    return false//CategoriesMap[transaction.internationalBranchID];
}

function displayNextTransaction() {
    const finishedTransactions = incrementTransIndex();
    if (finishedTransactions) {
        changeSection('id-category-list-container');
        return;
    }
    displayCurrentTransaction();
}

function getCurrentTransaction() {
    return globalTransArray[globalTransArrayIndex];
}

function incrementTransIndex() {
    if (globalTransArrayIndex === globalTransArray.length - 1) {
        alert("no more array");
        console.error('displayCurrentTransaction: globalTransArrayIndex >= globalTransArray.length');
        return true;
    }
    globalTransArrayIndex++;
    return false;
}

function transformTransaction(transaction, category) {
    return {
        // id: transaction.trnIntId, // transaction ID
        // TODO: Add transaction.trnIntId to Table
        name: transaction.merchantName, // merchant name
        amount: transaction.amountForDisplay, // transaction amount
        date_str: transaction.trnPurchaseDate, // date string
        date: new Date(transaction.trnPurchaseDate), // Date object
        address: transaction.merchantAddress, // merchant address
        card_unique_id: transaction.cardUniqueId, // card unique ID
        category,                                // transaction category
        merchant_phone_no: transaction.merchantPhoneNo, // merchant phone number
        international_branch_id: transaction.internationalBranchID // international branch ID
    };
}