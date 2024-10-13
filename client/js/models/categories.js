

function getCategories(){
   return sendHttpRequest("categories")
}

async function displayCategories() {
  let catStr = '';

  const response = await getCategories();

  if (!response.status) {
    console.error("displayCategories");
    return;
  }

  const categories = response.data;

  for (const cat of categories) {
      catStr += `
           <div  onclick="attachTransaction('${cat.name}','${cat.id}')" class="c-card c-pointer" id="id-category-${cat.id}" >
              ${cat.name}
          </div>
      `;
  }

  $$('id-category-list-container').innerHTML = catStr;
}
