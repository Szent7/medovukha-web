<script lang="ts">
	import { activeToasts, dismissNotification, markRead } from './store';
</script>

<div class="toast-stack" aria-live="polite" aria-relevant="additions">
	{#each $activeToasts as n (n.id)}
		<div class="toast toast-{n.level}" role="status">
			<div class="toast-head">
				<div class="title">{n.title}</div>
				<button
					class="close"
					type="button"
					onclick={() => {
						markRead(n.id);
						dismissNotification(n.id);
					}}
					aria-label="Close"
				>
					×
				</button>
			</div>
			<div class="msg">{n.message}</div>
			{#if n.details}
				<details class="details">
					<summary>Details</summary>
					<pre>{n.details}</pre>
				</details>
			{/if}
		</div>
	{/each}
</div>

<style>
	.toast-stack {
		position: fixed;
		right: 16px;
		bottom: 16px;
		z-index: 1000;
		display: flex;
		flex-direction: column;
		gap: 10px;
		max-width: min(420px, calc(100vw - 32px));
	}

	.toast {
		background: var(--bg-2);
		border: 1px solid var(--border);
		border-left: 4px solid var(--accent);
		border-radius: 12px;
		padding: 10px 12px;
		box-shadow: var(--shadow);
	}

	.toast-error {
		border-left-color: var(--danger);
	}
	.toast-warn {
		border-left-color: var(--warn);
	}
	.toast-success {
		border-left-color: var(--success);
	}

	.toast-head {
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: 8px;
	}

	.title {
		font-weight: 700;
		font-size: 0.95rem;
	}

	.msg {
		margin-top: 4px;
		color: var(--fg-0);
		font-size: 0.95rem;
		word-break: break-word;
	}

	.close {
		all: unset;
		cursor: pointer;
		color: var(--fg-muted);
		font-size: 1.25rem;
		line-height: 1;
		padding: 0 2px;
	}
	.close:hover {
		color: var(--fg-0);
	}

	.details {
		margin-top: 8px;
		color: var(--fg-muted);
	}
	.details pre {
		margin: 6px 0 0;
		padding: 8px;
		border-radius: 10px;
		background: var(--bg-1);
		border: 1px solid var(--border);
		overflow: auto;
		max-height: 180px;
	}
</style>
