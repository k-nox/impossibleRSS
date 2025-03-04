import { Urls } from '$lib/wailsjs/go/feed/FeedList';
import type { PageLoad } from './$types';

export const load: PageLoad = async () => {
	return {
		feedUrls: await Urls(),
	};
};
