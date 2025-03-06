<script lang="ts">
	import AppSidebar from '$lib/components/app-sidebar.svelte';
	import { SidebarProvider, SidebarTrigger } from '$lib/components/ui/sidebar';
	import { feedList } from '$lib/feeds.svelte';
	import '../app.css';
	import type { LayoutProps } from './$types';

	let { children, data }: LayoutProps = $props();
	feedList.feeds = [...data.feeds];
	let menuItems = $derived(
		feedList.feeds.map((feed) => {
			return {
				title: feed.title,
				url: `#${feed.title}`,
			};
		})
	);
</script>

<SidebarProvider>
	<AppSidebar {menuItems} />
	<main>
		<SidebarTrigger />
		{@render children?.()}
	</main>
</SidebarProvider>
