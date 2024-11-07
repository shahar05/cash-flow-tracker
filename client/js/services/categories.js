const categoryService = (()=>{

    async function getAllCategories() {
        const response = await getCategories();

        if (response.error) {
          console.error("displayCategories, Error: response.error");
          return null;
        }
      
        return response.data;
    }

    return {
        getAllCategories
    }
})()