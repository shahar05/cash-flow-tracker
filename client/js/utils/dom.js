function $$(elemId) {
    return document.getElementById(elemId);
}

// Collapse and open menu bar
function toggleMenu(close) {
    const closedClass = 'c-menu-closed';
    const navbarClass = document.querySelector('.mmd-c-navbar').classList;
    (close !== undefined) ? navbarClass.toggle(closedClass, close) : navbarClass.toggle(closedClass);
}