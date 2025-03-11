<script lang="ts">
	import * as Sidebar from '$lib/components/ui/sidebar';
	import * as Collapsible from '$lib/components/ui/collapsible';
	import ChevronDown from 'lucide-svelte/icons/chevron-down';
	import Plus from 'lucide-svelte/icons/plus';
	import Rss from 'lucide-svelte/icons/rss';
	import type { ComponentProps } from 'svelte';
	import NewFeedDialog from './new-feed-dialog.svelte';

	let {
		ref = $bindable(null),
		feedMenuItems,
		addNewFeed,
		...restProps
	}: ComponentProps<typeof Sidebar.Root> & {
		feedMenuItems: Array<{ title: string; url: string; isActive: boolean }>;
		addNewFeed: (url: string) => void;
	} = $props();

	let dialogIsOpen = $state(false);
</script>

<Sidebar.Root variant="floating" {...restProps}>
	<Sidebar.Header>
		<Sidebar.Menu>
			<Sidebar.MenuItem>
				<Sidebar.MenuButton size="lg">
					{#snippet child({ props })}
						<a href="##" {...props}>
							<div
								class="flex aspect-square size-8 items-center justify-center rounded-lg bg-sidebar-primary text-sidebar-primary-foreground"
							>
								<Rss />
							</div>
							<div class="flex flex-col gap-0.5 leading-none">
								<span class="font-semibold">impossibleRSS</span>
							</div>
						</a>
					{/snippet}
				</Sidebar.MenuButton>
			</Sidebar.MenuItem>
		</Sidebar.Menu>
	</Sidebar.Header>
	<Sidebar.Content>
		<Collapsible.Root open class="group/collapsible">
			<Sidebar.Group>
				<Sidebar.GroupLabel>
					{#snippet child({ props })}
						<Collapsible.Trigger {...props}>
							<ChevronDown
								class="transition-transform group-data-[state=open]/collapsible:rotate-180"
							/>
							Feeds
						</Collapsible.Trigger>
					{/snippet}
				</Sidebar.GroupLabel>
				<Sidebar.GroupAction title="Add Feed" onclick={() => (dialogIsOpen = !dialogIsOpen)}>
					<Plus /> <span class="sr-only">Add Feed</span>
					<NewFeedDialog bind:isOpen={dialogIsOpen} {addNewFeed} />
				</Sidebar.GroupAction>
				<Collapsible.Content>
					<Sidebar.GroupContent>
						<Sidebar.Menu class="gap-2">
							{#each feedMenuItems as feedMenuItem (feedMenuItem.title)}
								<Sidebar.MenuButton class="font-medium" isActive={feedMenuItem.isActive}>
									{#snippet child({ props })}
										<a href={feedMenuItem.url} {...props}>
											{feedMenuItem.title}
										</a>
									{/snippet}
								</Sidebar.MenuButton>
							{/each}
						</Sidebar.Menu>
					</Sidebar.GroupContent>
				</Collapsible.Content>
			</Sidebar.Group>
		</Collapsible.Root>
	</Sidebar.Content>
</Sidebar.Root>
