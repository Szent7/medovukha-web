<script lang="ts">
	import {
		clearAll,
		markAllRead,
		markRead,
		notificationCenterOpen,
		notifications,
		toggleNotificationCenter
	} from './store';
</script>

{#if $notificationCenterOpen}
	<button
		class="backdrop"
		type="button"
		aria-label="Close notifications"
		onclick={() => toggleNotificationCenter(false)}
	></button>
	<aside class="panel" aria-label="Notifications">
		<header class="header">
			<div class="title">Notifications</div>
			<div class="actions">
				<button class="btn" type="button" onclick={markAllRead}>Mark all read</button>
				<button class="btn" type="button" onclick={clearAll}>Clear</button>
				<button class="icon" type="button" onclick={() => toggleNotificationCenter(false)}>
					×
				</button>
			</div>
		</header>

		<div class="list">
			{#if $notifications.length === 0}
				<div class="empty">No notifications</div>
			{:else}
				{#each $notifications as n (n.id)}
					<button
						class="item {n.readAt ? 'read' : 'unread'}"
						type="button"
						onclick={() => markRead(n.id)}
					>
						<div class="row">
							<span class="level level-{n.level}"></span>
							<span class="item-title">{n.title}</span>
							<span class="time">{new Date(n.createdAt).toLocaleTimeString()}</span>
						</div>
						<div class="item-msg">{n.message}</div>
						{#if n.details}
							<pre class="item-details">{n.details}</pre>
						{/if}
					</button>
				{/each}
			{/if}
		</div>
	</aside>
{/if}

<style>
	.backdrop {
		all: unset;
		position: fixed;
		inset: 0;
		background: rgba(0, 0, 0, 0.45);
		z-index: 900;
		cursor: pointer;
	}

	.panel {
		position: fixed;
		top: 12px;
		right: 12px;
		bottom: 12px;
		width: min(520px, calc(100vw - 24px));
		background: var(--bg-2);
		border: 1px solid var(--border);
		border-radius: 14px;
		box-shadow: var(--shadow);
		z-index: 950;
		display: flex;
		flex-direction: column;
	}

	.header {
		padding: 12px;
		border-bottom: 1px solid var(--border);
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: 10px;
	}

	.title {
		font-weight: 800;
	}

	.actions {
		display: flex;
		gap: 8px;
		align-items: center;
	}

	.btn {
		padding: 6px 10px;
		border-radius: 10px;
		border: 1px solid var(--border);
		background: var(--bg-1);
		color: var(--fg-0);
		cursor: pointer;
	}
	.btn:hover {
		border-color: var(--accent);
	}

	.icon {
		all: unset;
		cursor: pointer;
		padding: 0 4px;
		font-size: 1.4rem;
		color: var(--fg-muted);
	}
	.icon:hover {
		color: var(--fg-0);
	}

	.list {
		padding: 12px;
		overflow: auto;
		display: flex;
		flex-direction: column;
		gap: 10px;
	}

	.empty {
		color: var(--fg-muted);
		padding: 8px;
	}

	.item {
		text-align: left;
		border-radius: 12px;
		border: 1px solid var(--border);
		background: var(--bg-1);
		padding: 10px;
		cursor: pointer;
	}
	.item:hover {
		border-color: var(--accent);
	}

	.item.read {
		opacity: 0.75;
	}

	.row {
		display: flex;
		align-items: center;
		gap: 8px;
	}

	.level {
		width: 8px;
		height: 8px;
		border-radius: 999px;
		background: var(--accent);
	}
	.level-error {
		background: var(--danger);
	}
	.level-warn {
		background: var(--warn);
	}
	.level-success {
		background: var(--success);
	}

	.item-title {
		font-weight: 700;
		flex: 1;
	}

	.time {
		color: var(--fg-muted);
		font-size: 0.85rem;
	}

	.item-msg {
		margin-top: 4px;
		word-break: break-word;
	}

	.item-details {
		margin: 8px 0 0;
		padding: 8px;
		border-radius: 10px;
		background: var(--bg-0);
		border: 1px solid var(--border);
		max-height: 220px;
		overflow: auto;
		color: var(--fg-muted);
		white-space: pre-wrap;
	}
</style>
