const baseUrl = "localhost:8080/";

async function sendHttpRequest(endpoint, method = 'GET', body = null, headers = {}) {

    // Define allowed methods
    const allowedMethods = ['GET', 'POST', 'PUT', 'DELETE'];

    // Validate the method
    if (!allowedMethods.includes(method.toUpperCase())) {
        throw new Error(`Invalid HTTP method: ${method}. Allowed methods are ${allowedMethods.join(', ')}.`);
    }

    const options = {
        method: method.toUpperCase(),
        headers: {
            'Content-Type': 'application/json',
            ...headers
        }
    };

    // Include body if method is POST or PUT
    if (body && (method === 'POST' || method === 'PUT')) {
        options.body = JSON.stringify(body);
    }

    try {
        const response = await fetch(baseUrl + endpoint, options);

        // Check if the response is okay (status in the range 200-299)
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        // Parse and return the JSON response
        return await response.json();
    } catch (error) {
        console.error('Error making HTTP call:', error);
        throw error; // Rethrow the error for further handling if necessary
    }
}

// // Example usage:
// // GET request
// makeHttpCall('https://api.example.com/data')
//     .then(data => console.log(data))
//     .catch(error => console.error(error));

// // POST request
// makeHttpCall('https://api.example.com/data', 'POST', { key: 'value' })
//     .then(data => console.log(data))
//     .catch(error => console.error(error));
