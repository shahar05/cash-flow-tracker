const NavService = (()=>{

    const allNavItems = document.querySelectorAll('#id-nav-items > a');
    const allSections = document.querySelectorAll("#id-main-container section");        

    const pageMap = {
        "id-upload-file": UploadFilePage,
        "id-attach-trans-with-category": AttachTransactionPage,
        "id-view-analyses": CategoryAnalysisPage,
        "id-reattach-category": ReattachTransactionPage,
        "id-graph-total-expenses": ExpensesGraphPage,
        "id-pie-chart-per-month": MonthlyPieChartPage,
        "id-category-drill-down": CategoryPieChartPage,
    }

    function changeSection(elem, elemId, ...infiniteArguments) {
        
        // Use the element ID if `elem` is provided
        if (elem) elemId = elem.id;
    
        // Get the current navigation item and corresponding section element
        const currentNavItem = $$(elemId);
        const sectionElement = $$(`${elemId}-section`);
    
        // Error handling for missing elements
        if (!sectionElement || !allSections?.length || !currentNavItem) {
            console.error(`changeSection: Cannot navigate to other section. elem: ${sectionElement}, allSections: ${allSections}, currentNavItem: ${currentNavItem}`);
            return;
        }
    
        // Remove "active" class from all sections and add it to the selected one
        allSections.forEach(sec => sec.classList.remove("c-active-section"));    
        sectionElement.classList.add("c-active-section");
    
        // Highlight the current navigation item
        allNavItems.forEach(navItem => navItem.classList.toggle('c-active-nav-item', navItem === currentNavItem));
        
        // Pass `infiniteArguments` as separate arguments to `onInit`
        pageMap[elemId].onInit(...infiniteArguments);
    }
    

    return {
        changeSection
    }

})()

