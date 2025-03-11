import { Feeds } from '$lib/wailsjs/go/app/FeedList';
import type { LayoutLoad } from './$types';

export const prerender = true;
export const ssr = false;

export const load: LayoutLoad = async () => {
	return {
		feeds: await Feeds(),
	};
};
