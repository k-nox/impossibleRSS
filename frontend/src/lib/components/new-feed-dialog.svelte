<script lang="ts">
	import * as Dialog from './ui/dialog';
	import { Input } from './ui/input';
	import { Label } from './ui/label';
	import { Button } from './ui/button';

	let { isOpen = $bindable(), addNewFeed }: { isOpen: boolean; addNewFeed: (url: string) => void } =
		$props();

	let url = $state('');
</script>

<Dialog.Root bind:open={isOpen}>
	<!-- <Dialog.Trigger title="Add Feed"><Plus /> <span class="sr-only">Add Feed</span></Dialog.Trigger> -->
	<Dialog.Content class="sm:max-w-[425px]">
		<Dialog.Header>
			<Dialog.Title>Add New Feed</Dialog.Title>
			<Dialog.Description>Click save when you're done.</Dialog.Description>
		</Dialog.Header>
		<div>
			<Label for="url" class="text-right">URL</Label>
			<Input id="url" bind:value={url} placeholder="https://example.com/index.xml" />
		</div>
		<Dialog.Footer>
			<Button type="button" onclick={() => addNewFeed(url)}>Save</Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
