<script lang="ts">
	import { page } from '$app/state';
	import SidebarPage from '$lib/components/sidebar-page.svelte';
	import { AddFeed } from '$lib/wailsjs/go/app/FeedList';
	import { app } from '$lib/wailsjs/go/models';
	import { EventsOff, EventsOn } from '$lib/wailsjs/runtime/runtime';
	import { setContext } from 'svelte';
	import '../app.css';
	import type { LayoutProps } from './$types';

	let { children, data }: LayoutProps = $props();

	let feeds = $state(
		data.feeds.reduce<Record<string, app.Feed>>((acc, cur) => {
			acc[cur.title] = cur;
			return acc;
		}, {})
	);

	$effect(() => {
		EventsOn(app.Event.NEW_ITEM, (...data: Array<app.Item>) => {
			data.forEach((item) => feeds[item.feedURL].items.push(item));
		});
		return () => {
			EventsOff(app.Event.NEW_ITEM);
		};
	});

	let selected = $derived(
		page.url.hash !== '' ? feeds[decodeURIComponent(page.url.hash.substring(1))] : null
	);

	let feedMenuItems = $derived(
		Object.keys(feeds).map((feedTitle) => {
			return {
				title: feedTitle,
				url: `#${feedTitle}`,
				isActive: selected?.title === feedTitle,
			};
		})
	);

	setContext('selected', () => selected);

	const addNewFeed = async (url: string) => {
		const newFeed = await AddFeed(url);
		feeds[newFeed.title] = newFeed;
	};
</script>

<SidebarPage {feedMenuItems} {addNewFeed} pageTitle={selected?.title}>
	<main>
		<div>{@render children?.()}</div>
	</main>
</SidebarPage>
