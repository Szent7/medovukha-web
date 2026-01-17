<script lang="ts">
	import type { Snippet } from 'svelte';

	type Props = {
		children?: Snippet;
		variant?: 'neutral' | 'danger' | 'warn' | 'success' | 'accent';
	};

	let { children, variant = 'neutral' }: Props = $props();
</script>

<span class="badge badge-{variant}">
	{@render children?.()}
</span>

<style>
	.badge {
		display: inline-flex;
		align-items: center;
		gap: 6px;
		padding: 2px 8px;
		border-radius: 999px;
		font-weight: 800;
		font-size: 0.72rem;
		text-transform: uppercase;
		letter-spacing: 0.06em;
		border: 1px solid var(--border);
		background: var(--bg-2);
		color: var(--fg-0);
	}

	.badge-accent {
		border-color: var(--accent);
		background: color-mix(in srgb, var(--accent-weak) 65%, var(--bg-2));
	}
	.badge-danger {
		border-color: var(--danger);
	}
	.badge-warn {
		border-color: var(--warn);
	}
	.badge-success {
		border-color: var(--success);
	}
</style>
