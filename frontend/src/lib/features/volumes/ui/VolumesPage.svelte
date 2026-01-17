<script lang="ts">
	import { onMount } from 'svelte';
	import Card from '$lib/shared/ui/Card.svelte';
	import PageHeader from '$lib/shared/ui/PageHeader.svelte';
	import Badge from '$lib/shared/ui/Badge.svelte';
	import { getVolumeList, removeVolume } from '../api';
	import type { ListVolumeBaseInfo } from '../types';

	let volList: ListVolumeBaseInfo = { items: [] };
	let loading = true;
	let selectedNames: string[] = [];

	$: canRemove =
		selectedNames.length > 0 &&
		selectedNames.every((name) => {
			const item = volList.items.find((i) => i.name === name);
			return !item?.is_used;
		});
	$: canForceRemove = selectedNames.length > 0;

	async function refresh() {
		loading = true;
		selectedNames = [];
		try {
			volList = await getVolumeList();
		} catch {
			// requestApi already notifies
		} finally {
			loading = false;
		}
	}

	onMount(refresh);

	function toggleSelected(name: string, checked: boolean) {
		selectedNames = checked ? [...selectedNames, name] : selectedNames.filter((x) => x !== name);
	}

	function selectAll(checked: boolean) {
		selectedNames = checked ? volList.items.map((v) => v.name) : [];
		const elements = document.querySelectorAll('input[name=checkbox-item]');
		Array.prototype.forEach.call(elements, (el: HTMLInputElement) => {
			if (!el.disabled) el.checked = checked;
		});
	}

	async function doRemove(force: boolean) {
		const names = [...selectedNames];
		selectedNames = [];
		try {
			await Promise.all(names.map((id) => removeVolume({ id, force })));
		} catch {
			// requestApi already notifies
		} finally {
			await refresh();
		}
	}
</script>

<svelte:head>
	<title>Volumes</title>
</svelte:head>

<Card>
	<PageHeader title="Volumes" subtitle="Docker volumes">
		<button type="button" onclick={refresh} disabled={loading}>Refresh</button>
		<button type="button" onclick={() => doRemove(false)} disabled={!canRemove}>Remove</button>
		<button type="button" onclick={() => doRemove(true)} disabled={!canForceRemove}
			>Force remove</button
		>
	</PageHeader>

	{#if loading}
		<div class="muted">Loading…</div>
	{:else}
		<table>
			<thead>
				<tr>
					<th>
						<input
							type="checkbox"
							name="checkbox-item-all"
							onchange={(e) => selectAll((e.target as HTMLInputElement).checked)}
						/>
					</th>
					<th>Name</th>
					<th>Driver</th>
					<th>Mountpoint</th>
					<th>Created</th>
				</tr>
			</thead>
			<tbody>
				{#each volList.items as volume (volume.name)}
					<tr>
						<td>
							<input
								type="checkbox"
								name="checkbox-item"
								checked={selectedNames.includes(volume.name)}
								onchange={(e) =>
									toggleSelected(volume.name, (e.target as HTMLInputElement).checked)}
							/>
						</td>
						<td>
							<span>{volume.name}</span>
							{#if !volume.is_used}
								<Badge variant="warn">unused</Badge>
							{/if}
						</td>
						<td>{volume.driver}</td>
						<td>{volume.mountpoint}</td>
						<td>{volume.created}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	{/if}
</Card>

<style>
	.muted {
		color: var(--fg-muted);
	}

	td :global(.badge) {
		margin-left: 8px;
	}
</style>
