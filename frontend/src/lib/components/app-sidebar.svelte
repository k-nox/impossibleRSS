<script lang="ts">
	import {
		Sidebar,
		SidebarContent,
		SidebarGroupContent,
		SidebarGroupLabel,
		SidebarMenu,
		SidebarMenuButton,
		SidebarMenuItem,
	} from './ui/sidebar';

	import RSS from 'lucide-svelte/icons/rss';

	let { menuItems } = $props();
</script>

<Sidebar>
	<SidebarContent>
		<SidebarGroupLabel>Feeds</SidebarGroupLabel>
		<SidebarGroupContent>
			<SidebarMenu>
				{#each menuItems as feed (feed.title)}
					<SidebarMenuItem>
						<SidebarMenuButton>
							{#snippet child({ props })}
								<a href={feed.url} {...props}>
									<RSS />
									<span>{feed.title}</span>
								</a>
							{/snippet}
						</SidebarMenuButton>
					</SidebarMenuItem>
				{/each}
			</SidebarMenu>
		</SidebarGroupContent>
	</SidebarContent>
</Sidebar>
