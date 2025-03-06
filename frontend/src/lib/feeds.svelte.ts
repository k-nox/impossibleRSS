import type { app } from './wailsjs/go/models';

interface FeedList {
	feeds: Array<app.Feed>;
}
export const feedList: FeedList = $state({ feeds: [] });
