const allNavItems = document.querySelectorAll('#id-nav-items > a');

function changeSection(sectionID, callback) {
    const elem = $$(sectionID);
    const allSections = document.querySelectorAll("#id-main-container section");
    const currentNavItem = $$(sectionID + "-a");


    if (elem && allSections.length > 0) {
        allSections.forEach(sec => sec.classList.remove("c-active-section"));
        elem.classList.add("c-active-section");
        allNavItems.forEach(navItem => navItem.classList.toggle('c-active-nav-item', navItem === currentNavItem));
    }

    if (callback) {
        callback();
    }        
}

function toggleMenu(close) {
    const closedClass = 'c-menu-closed';
    const navbarClass = document.querySelector('.mmd-c-navbar').classList;
    (close !== undefined) ? navbarClass.toggle(closedClass, close) : navbarClass.toggle(closedClass);
}