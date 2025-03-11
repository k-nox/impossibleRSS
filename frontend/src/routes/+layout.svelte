<script lang="ts">
	import AppSidebar from '$lib/components/app-sidebar.svelte';
	import { SidebarProvider, SidebarTrigger } from '$lib/components/ui/sidebar';
	import { feedList } from '$lib/feeds.svelte';
	import { app } from '$lib/wailsjs/go/models';
	import { EventsOff, EventsOn } from '$lib/wailsjs/runtime/runtime';
	import '../app.css';
	import type { LayoutProps } from './$types';

	let { children, data }: LayoutProps = $props();
	feedList.feeds = data.feeds.map((feed) => {
		return {
			[feed.title]: feed,
		};
	});
	$effect(() => {
		EventsOn(app.Event.NEW_ITEM, (...data: any) => {
			const item = data[0];
			feedList.feeds[item.title].items[item.guid] = item;
		});
		return () => {
			EventsOff(app.Event.NEW_ITEM);
		};
	});

	let menuItems = $derived.by(() => {
		let items = [];
		for (let title in feedList.feeds) {
			items.push({
				title: feedList.feeds[title].title,
				url: `#${feedList.feeds[title].title}`,
			});
		}
		return items;
	});
</script>

<SidebarProvider>
	<AppSidebar {menuItems} />
	<main>
		<SidebarTrigger />
		{@render children?.()}
	</main>
</SidebarProvider>
