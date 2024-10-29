// Navigator class to manage page navigation
class Navigator {
    constructor() {
        this.page = null;
        this.allNavItems = document.querySelectorAll('#id-nav-items > a');
    }

    changeSection(sectionID) {
        const elem = $$(sectionID);
        const allSections = document.querySelectorAll("#id-main-container section");
        const currentNavItem = $$(sectionID + "-a");
        
        if (!elem || !allSections?.length || !currentNavItem) {
            console.error(`changeSection: Cant navigate to other section.  elem: ${elem} , allSections: ${allSections}, currentNavItem: ${currentNavItem}`);
            return;
        }
        
        // Remove all active sections
        allSections.forEach(sec => sec.classList.remove("c-active-section"));
        
        // Choose the Active Section
        elem.classList.add("c-active-section");

        // Highlighted Navbar item  
        allNavItems.forEach(navItem => navItem.classList.toggle('c-active-nav-item', navItem === currentNavItem));
        
    }
    

    // Method to navigate to a new page
    navigate(newPage) {
        // If there's a current page, we can clean up or do something
        if (this.page) {
            console.log(`Leaving ${this.page.constructor.name} page.`);
        }        
        
        // Set new page
        if (newPage) {
            this.page = newPage;   
        }        

        // Call the init function
        this.page.init(); 

        // Change the section using the pageId of the current page
        this.changeSection(this.page.pageId);
    }
}