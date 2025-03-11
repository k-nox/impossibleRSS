export namespace app {
	
	export enum Event {
	    REFRESH_ERROR = "RefreshError",
	    NEW_ITEM = "NewItem",
	}
	export class Item {
	    guid: string;
	    title: string;
	    authors: string[];
	    content: string;
	    description: string;
	    publishedDate?: Date;
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
	export class Feed {
	    title: string;
	    description: string;
	    link: string;
	    items: Record<string, Item>;
	
	    static createFrom(source: any = {}) {
	        return new Feed(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.description = source["description"];
	        this.link = source["link"];
	        this.items = this.convertValues(source["items"], Item, true);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

