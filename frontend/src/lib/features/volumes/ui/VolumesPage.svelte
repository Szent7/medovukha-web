<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import Card from '$lib/shared/ui/Card.svelte';
	import PageHeader from '$lib/shared/ui/PageHeader.svelte';
	import Badge from '$lib/shared/ui/Badge.svelte';
	import { wsUrl } from '$lib/core/http/client';
	import { notifyError } from '$lib/app/notifications/store';
	import { getVolumeById, getVolumeList, removeVolume } from '../api';
	import { VolumeEventSchema, type ListVolumeBaseInfo } from '../types';

	let volList: ListVolumeBaseInfo = { items: [] };
	let loading = true;
	let selectedNames: string[] = [];

	let socket: WebSocket | null = null;

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

	function updateVolumeUsedByName(name: string, isUsed: boolean) {
		const idx = volList.items.findIndex((v) => v.name === name);
		if (idx === -1) return;
		const newItems = [...volList.items];
		newItems[idx] = { ...newItems[idx], is_used: isUsed };
		volList = { items: newItems };
	}

	function removeFromListByName(name: string) {
		volList = { items: volList.items.filter((v) => v.name !== name) };
		selectedNames = selectedNames.filter((x) => x !== name);
	}

	function upsertIntoList(item: (typeof volList.items)[number]) {
		const idx = volList.items.findIndex((v) => v.name === item.name);
		if (idx === -1) {
			volList = { items: [item, ...volList.items] };
			return;
		}
		const newItems = [...volList.items];
		newItems[idx] = { ...newItems[idx], ...item };
		volList = { items: newItems };
	}

	function connectEvents() {
		try {
			socket = new WebSocket(wsUrl('/ws/volumeEvents'));
			socket.addEventListener('message', (e) => {
				try {
					const parsed = VolumeEventSchema.parse(JSON.parse(String(e.data)));
					const action = parsed.action;
					const name = parsed.actor.id;
					if (['remove', 'delete', 'destroy'].includes(action)) {
						removeFromListByName(name);
						return;
					}

					if (action === 'create') {
						getVolumeById({ id: name })
							.then((item) => upsertIntoList(item))
							.catch(() => {
								// requestApi already notifies
							});
						return;
					}

					//if (['destroy'].includes(action)) return;
					if (action === 'use') updateVolumeUsedByName(name, true);
					if (action === 'unuse') updateVolumeUsedByName(name, false);
				} catch (err) {
					notifyError(err, 'Wrong WebSocket-message');
				}
			});
			socket.addEventListener('error', (e) => {
				notifyError(new Error(String(e)), 'WebSocket error');
			});
		} catch (err) {
			notifyError(err, 'WebSocket error');
		}
	}

	onMount(() => {
		refresh();
		connectEvents();
	});

	onDestroy(() => {
		socket?.close();
	});

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
