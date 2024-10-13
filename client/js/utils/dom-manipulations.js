function changeSection(sectionID, callback) {
    const elem = $$(sectionID);
    const allSections = document.querySelectorAll("#id-main-container section");

    if (elem && allSections.length > 0) {
        allSections.forEach(sec => sec.classList.remove("c-active-section"));
        elem.classList.add("c-active-section");
    }

    if (callback) {
        callback();
    }        
}