<script lang="ts">
	import type { app } from '$lib/wailsjs/go/models';
	import { getContext } from 'svelte';
	let selected: () => app.Feed = getContext('selected');
</script>

{#if selected()}
	{#each Object.keys(selected().items) as itemGuid (itemGuid)}
		<div class="prose">
			<h1>{selected().items[itemGuid].title}</h1>
			{@html selected().items[itemGuid].content}
		</div>
	{/each}
{:else}
	<div>Select a feed!</div>
{/if}
