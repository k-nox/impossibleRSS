export namespace app {
	
	export class Feed {
	    Title: string;
	    Description: string;
	    Link: string;
	
	    static createFrom(source: any = {}) {
	        return new Feed(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Title = source["Title"];
	        this.Description = source["Description"];
	        this.Link = source["Link"];
	    }
	}

}

