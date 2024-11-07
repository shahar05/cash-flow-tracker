class Page {

    constructor(pageId) {
        this.pageId = pageId; // Initialize the pageId property
    }

    init() {
        console.log(`Initializing ${this.constructor.name} page.`);
    }
}


var page = new Page();