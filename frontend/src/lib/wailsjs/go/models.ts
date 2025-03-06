export namespace app {
	
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
	export class Item {
	    guid: string;
	    title: string;
	    authors: string[];
	    content: string;
	    description: string;
	    publishedDate: Date;
	    feedURL: string;
	
	    static createFrom(source: any = {}) {
	        return new Item(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.guid = source["guid"];
	        this.title = source["title"];
	        this.authors = source["authors"];
	        this.content = source["content"];
	        this.description = source["description"];
	        this.publishedDate = new Date(source["publishedDate"]);
	        this.feedURL = source["feedURL"];
	    }
	}

}

