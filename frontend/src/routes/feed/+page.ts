import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';
import { Urls } from '$lib/wailsjs/go/feed/FeedList';
import { ParseURL } from '$lib/wailsjs/go/feed/Parser';

export const load: PageLoad = async ({ url }) => {
	const id = url.searchParams.get('id');
	if (!id) {
		error(400, 'Bad Request');
	}

	const numId = parseInt(id);
	const urls = await Urls();

	return {
		feed: await ParseURL(urls[numId]),
	};
};
