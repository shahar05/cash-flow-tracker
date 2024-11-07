const strings = function () {

    function capitalizeFirstLetter(str) {
        if (!str) return '';  // Handle empty string
        return str.charAt(0).toUpperCase() + str.slice(1);
    }

    return {
        capitalizeFirstLetter
    }
}()

