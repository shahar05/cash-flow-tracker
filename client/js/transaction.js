async function attachTransaction(cat) {

    const body = {
        transaction:getCurrentTransaction(),
        category: cat
    };
    
    const response = await sendHttpRequest("attach-transaction", "POST", body);

    console.log(response);

    displayNextTransaction();
}