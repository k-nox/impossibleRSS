<script lang="ts">
	import AppSidebar from '$lib/components/app-sidebar.svelte';
	import { SidebarProvider, SidebarTrigger } from '$lib/components/ui/sidebar';
	import '../app.css';
	let { children } = $props();
</script>

<SidebarProvider>
	<AppSidebar />
	<main>
		<SidebarTrigger />
		{@render children?.()}
	</main>
</SidebarProvider>
