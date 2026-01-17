<script lang="ts">
	import Card from '$lib/shared/ui/Card.svelte';
	import PageHeader from '$lib/shared/ui/PageHeader.svelte';
	import Badge from '$lib/shared/ui/Badge.svelte';
</script>

<svelte:head>
	<title>Home</title>
</svelte:head>

<Card>
	<PageHeader title="Medovukha panel" subtitle="Alpha">
		<Badge variant="accent">v0.0.1</Badge>
	</PageHeader>
	<div class="text">Use the sidebar to navigate.</div>
</Card>

<style>
	.text {
		margin-top: 10px;
		color: var(--fg-muted);
	}
</style>
