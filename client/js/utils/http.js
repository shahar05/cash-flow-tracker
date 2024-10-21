var http = (() => {
    const baseUrl = "http://localhost:8080/";


    async function sendRequest(endpoint, method = 'GET', body = null, headers = {}) {
        const allowedMethods = ['GET', 'POST', 'PUT', 'DELETE'];

        // Validate the method
        if (!allowedMethods.includes(method.toUpperCase())) {
            return { error: `Invalid HTTP method: ${method}. Allowed methods are ${allowedMethods.join(', ')}.`, data: null };
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
                return { error: `HTTP error! status: ${response.status}`, data: null };
            }

            const jsonResponse = await response.json();

            // Validate that the response matches the expected backend structure
            if (typeof jsonResponse !== 'object' || jsonResponse === null) {
                return { error: 'Invalid response format received from the server', data: null };
            }

            // Check the status field in the backend's response structure
            if (!jsonResponse.status) {
                const errorMessage = jsonResponse.error ? jsonResponse.error.message : 'Unknown error occurred';
                const errorCode = jsonResponse.error ? jsonResponse.error.code : 500;
                return { error: `Error: ${errorMessage} (Code: ${errorCode})`, data: null };
            }

            // If status is true, return the data field
            return { error: null, data: jsonResponse.data };

        } catch (error) {
            // Catch any network or fetch-related errors
            console.error('Error making HTTP call:', error);
            return { error: error.message, data: null };
        }
    }


    return {
        sendRequest
    }
})();

