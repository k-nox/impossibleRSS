import type { app } from './wailsjs/go/models';

export interface FeedList {
	feeds: Record<string, app.Feed>;
	selected: app.Feed | null;
}

export const feedList: FeedList = $state({ feeds: {}, selected: null });
