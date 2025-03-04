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

	const feeds = [
		{
			title: 'lux@web',
			url: '#',
		},
		{
			title: 'testing',
			url: '#',
		},
	];
</script>

<Sidebar>
	<SidebarContent>
		<SidebarGroupLabel>Feeds</SidebarGroupLabel>
		<SidebarGroupContent>
			<SidebarMenu>
				{#each feeds as feed (feed.title)}
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
