<script lang="ts">
	import type { Snippet } from 'svelte';

	let { children }: { children?: Snippet } = $props();
</script>

<section class="card">
	{@render children?.()}
</section>

<style>
	.card {
		background: var(--bg-1);
		border: 1px solid var(--border);
		border-radius: 16px;
		box-shadow: var(--shadow);
		padding: 14px;
		overflow: auto;
	}
</style>
