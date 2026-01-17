<script lang="ts">
	import '$lib/shared/styles/global.css';
	import AppShell from '$lib/app/layout/AppShell.svelte';
	import ToastStack from '$lib/app/notifications/ToastStack.svelte';
	import NotificationCenter from '$lib/app/notifications/NotificationCenter.svelte';
	let { children } = $props();
</script>

<AppShell>
	{@render children()}
</AppShell>

<ToastStack />
<NotificationCenter />
