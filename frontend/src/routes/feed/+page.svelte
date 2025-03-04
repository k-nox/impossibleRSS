<script lang="ts">
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
	console.log(data.feed.items);
</script>

<h1>{data.feed.title}</h1>
<p>{data.feed.description}</p>
<p>{data.feed.link}</p>
<article class="prose prose-slate">
	{@html data.feed.items[0].content}
</article>
