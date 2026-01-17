<script lang="ts">
	import type { Snippet } from 'svelte';

	type Props = {
		title: string;
		subtitle?: string;
		children?: Snippet;
	};

	let { title, subtitle, children }: Props = $props();
</script>

<div class="header">
	<div class="titles">
		<h1>{title}</h1>
		{#if subtitle}
			<div class="subtitle">{subtitle}</div>
		{/if}
	</div>
	{#if children}
		<div class="actions">{@render children?.()}</div>
	{/if}
</div>

<style>
	.header {
		display: flex;
		align-items: flex-start;
		justify-content: space-between;
		gap: 12px;
	}

	h1 {
		margin: 0;
		font-size: 1.35rem;
		letter-spacing: 0.01em;
	}

	.subtitle {
		margin-top: 2px;
		color: var(--fg-muted);
		font-size: 0.95rem;
	}

	.actions {
		display: flex;
		gap: 10px;
		flex-wrap: wrap;
		justify-content: flex-end;
	}
</style>
