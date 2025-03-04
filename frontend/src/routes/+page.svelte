<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import type { SubmitFunction } from '@sveltejs/kit';
	import { enhance } from '$app/forms';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();

	import { Card, CardContent, CardFooter, CardHeader, CardTitle } from '$lib/components/ui/card';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import { Add } from '$lib/wailsjs/go/feed/FeedList';

	let url = $state('');
	let feedUrls: Array<string> = $state([...data.feedUrls]);

	const addUrl: SubmitFunction = ({ cancel }) => {
		cancel();
		feedUrls.push(url);
		Add(url);
		url = '';
	};
</script>

{#snippet addFormCard()}
	<div class="flex min-h-screen flex-row items-center justify-center">
		<Card class="w-1/2">
			<CardHeader>
				<CardTitle>Add feed</CardTitle>
			</CardHeader>
			<form method="POST" use:enhance={addUrl}>
				<CardContent>
					<div class="grid w-full items-center gap-4">
						<div class="flex flex-col space-y-1.5">
							<Label for="url">URL</Label>
							<Input bind:value={url} id="url" placeholder="Feed URL" />
						</div>
					</div>
				</CardContent>
				<CardFooter class="flex justify-between">
					<Button variant="outline" id="cancel-add-url">Cancel</Button>
					<Button type="submit">Add</Button>
				</CardFooter>
			</form>
		</Card>
	</div>
{/snippet}

{#if !feedUrls.length}
	{@render addFormCard()}
{:else}
	<h1>Feeds</h1>
	<ul>
		{#each feedUrls as feedUrl, index (index)}
			<li>{feedUrl}</li>
		{/each}
	</ul>
{/if}
