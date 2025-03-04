export namespace feed {
	
	export class Feed {
	    title: string;
	    description: string;
	    link: string;
	
	    static createFrom(source: any = {}) {
	        return new Feed(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.description = source["description"];
	        this.link = source["link"];
	    }
	}

}

