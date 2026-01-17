<script lang="ts">
	import { page } from '$app/stores';
	import logo from '@assets/logo_small.svg';
	import { NAV } from '$lib/app/nav/nav';
	import { toggleNotificationCenter, unreadCount } from '$lib/app/notifications/store';

	let { children } = $props();
</script>

<div class="app">
	<aside class="sidebar">
		<div class="brand">
			<img src={logo} alt="logo" class="logo" />
			<div class="name">Medovukha</div>
		</div>
		<nav class="nav">
			{#each NAV as item (item.href)}
				<a class="link {$page.url.pathname === item.href ? 'active' : ''}" href={item.href}
					>{item.label}</a
				>
			{/each}
		</nav>
	</aside>

	<div class="main">
		<header class="topbar">
			<div class="spacer"></div>
			<button class="notif" type="button" onclick={() => toggleNotificationCenter()}>
				<span class="bell">Notifications</span>
				{#if $unreadCount > 0}
					<span class="count">{$unreadCount}</span>
				{/if}
			</button>
		</header>
		<main class="content">
			{@render children?.()}
		</main>
	</div>
</div>

<style>
	.app {
		height: 100vh;
		width: 100vw;
		display: grid;
		grid-template-columns: 260px 1fr;
		background: var(--bg-0);
	}

	.sidebar {
		background: var(--bg-2);
		border-right: 1px solid var(--border);
		padding: 14px;
		display: flex;
		flex-direction: column;
		gap: 14px;
	}

	.brand {
		display: flex;
		align-items: center;
		gap: 10px;
		padding: 10px;
		border-radius: 14px;
		background: color-mix(in srgb, var(--bg-1) 75%, var(--bg-2));
		border: 1px solid var(--border);
	}
	.logo {
		width: 34px;
		height: 34px;
	}
	.name {
		font-weight: 900;
		letter-spacing: 0.02em;
		font-variant: small-caps;
	}

	.nav {
		display: flex;
		flex-direction: column;
		gap: 6px;
	}

	.link {
		padding: 10px 10px;
		border-radius: 12px;
		border: 1px solid transparent;
		background: transparent;
		color: var(--fg-0);
		font-weight: 700;
	}
	.link:hover {
		border-color: var(--accent);
		background: var(--accent-weak);
		color: var(--accent);
	}
	.link.active {
		border-color: var(--accent);
		background: color-mix(in srgb, var(--accent-weak) 60%, var(--bg-2));
	}

	.main {
		display: flex;
		flex-direction: column;
		min-width: 0;
	}

	.topbar {
		height: 56px;
		display: flex;
		align-items: center;
		justify-content: flex-end;
		padding: 0 14px;
		border-bottom: 1px solid var(--border);
		background: var(--bg-1);
	}

	.notif {
		display: inline-flex;
		align-items: center;
		gap: 8px;
	}

	.count {
		min-width: 20px;
		height: 20px;
		border-radius: 999px;
		background: var(--accent);
		color: #001314;
		display: inline-flex;
		align-items: center;
		justify-content: center;
		font-size: 0.75rem;
		font-weight: 900;
		padding: 0 6px;
	}

	.content {
		padding: 14px;
		overflow: auto;
		min-width: 0;
	}

	@media (max-width: 920px) {
		.app {
			grid-template-columns: 1fr;
		}
		.sidebar {
			display: none;
		}
	}
</style>
