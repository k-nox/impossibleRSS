import { ParseURL } from '$lib/wailsjs/go/app/Parser';
import type { PageLoad } from './$types';

export const load: PageLoad = async () => {
	return {
		feed: await ParseURL('https://blog.luxatweb.dev/index.xml')
	};
};
