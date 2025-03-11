<script lang="ts">
	import { page } from '$app/state';
	import { feedList } from '$lib/feeds.svelte';

	// import { Button } from '$lib/components/ui/button';
	// import type { SubmitFunction } from '@sveltejs/kit';
	// import { enhance } from '$app/forms';
	// import { Input } from '$lib/components/ui/input';
	// import { Label } from '$lib/components/ui/label';
	// import { AddFeed } from '$lib/wailsjs/go/app/FeedList';
	// import { feedList } from '$lib/feeds.svelte';
	// import {
	// 	Dialog,
	// 	DialogContent,
	// 	DialogDescription,
	// 	DialogFooter,
	// 	DialogHeader,
	// 	DialogTitle,
	// 	DialogTrigger,
	// } from '$lib/components/ui/dialog';

	// let url = $state('');

	// const addUrl: SubmitFunction = ({ cancel }) => {
	// 	cancel();
	// 	AddFeed(url);
	// 	url = '';
	// };
	const selectedFeedTitle = $derived(page.url.hash !== '' ? page.url.hash.substring(1) : '');
</script>

<!-- {#snippet addFormDialog()} -->
<!-- 	<Dialog> -->
<!-- 		<DialogTrigger>Add Feed</DialogTrigger> -->
<!-- 		<DialogContent> -->
<!-- 			<DialogHeader> -->
<!-- 				<DialogTitle>Add Feed</DialogTitle> -->
<!-- 				<DialogDescription>Add a new feed here. Click save when you're done.</DialogDescription> -->
<!-- 			</DialogHeader> -->
<!-- 			<form method="dialog" use:enhance={addUrl}> -->
<!-- 				<div class="grid gap-4 py-4"> -->
<!-- 					<div class="grid grid-cols-4 items-center gap-4"> -->
<!-- 						<Label for="url">URL</Label> -->
<!-- 						<Input bind:value={url} id="url" placeholder="https://example.com/index.xml" /> -->
<!-- 					</div> -->
<!-- 				</div> -->
<!-- 				<DialogFooter> -->
<!-- 					<Button variant="outline" id="cancel-add-url" onclick={() => (url = '')}>Cancel</Button> -->
<!-- 					<Button type="submit">Add</Button> -->
<!-- 				</DialogFooter> -->
<!-- 			</form> -->
<!-- 		</DialogContent> -->
<!-- 	</Dialog> -->
<!-- {/snippet} -->

<!-- {#if !feedList.feeds.length} -->
<!-- 	{@render addFormDialog()} -->
<!-- {:else} -->
<!-- 	<h1>Feeds</h1> -->
<!-- 	<ul> -->
<!-- 		{#each feedList.feeds as feed (feed.link)} -->
<!-- 			<li><a href="/feed#{feed.title}">{feed.link}</a></li> -->
<!-- 		{/each} -->
<!-- 	</ul> -->
<!-- {/if} -->
{#if selectedFeedTitle !== ''}
	<div>{@html feedList.feeds[selectedFeedTitle].items[0].content}</div>
{:else}
	<div>Add a feed!</div>
{/if}
