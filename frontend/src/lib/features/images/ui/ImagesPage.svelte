<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import Card from '$lib/shared/ui/Card.svelte';
	import PageHeader from '$lib/shared/ui/PageHeader.svelte';
	import Badge from '$lib/shared/ui/Badge.svelte';
	import { formatId, sizeFormat, unixTimeFormat } from '$lib/shared/utils/format';
	import { wsUrl } from '$lib/core/http/client';
	import { notifyError } from '$lib/app/notifications/store';
	import { getImageById, getImageList, removeImage } from '../api';
	import { ImageEventSchema, type ListImageBaseInfo } from '../types';

	let imgList: ListImageBaseInfo = { items: [] };
	let loading = true;
	let selectedIds: string[] = [];

	let socket: WebSocket | null = null;

	$: canRemove =
		selectedIds.length > 0 &&
		selectedIds.every((id) => {
			const item = imgList.items.find((i) => i.id === id);
			return !item?.is_used;
		});
	$: canForceRemove = selectedIds.length > 0;

	async function refresh() {
		loading = true;
		selectedIds = [];
		try {
			imgList = await getImageList();
		} catch {
			// requestApi already notifies
		} finally {
			loading = false;
		}
	}

	onMount(() => {
		refresh();
		connectEvents();
	});

	function updateImageUsedById(id: string, isUsed: boolean) {
		const idx = imgList.items.findIndex((i) => i.id === id);
		if (idx === -1) return;
		const newItems = [...imgList.items];
		newItems[idx] = { ...newItems[idx], is_used: isUsed };
		imgList = { items: newItems };
	}

	function removeFromListById(id: string) {
		imgList = { items: imgList.items.filter((i) => i.id !== id) };
		selectedIds = selectedIds.filter((x) => x !== id);
	}

	function upsertIntoList(item: (typeof imgList.items)[number]) {
		const idx = imgList.items.findIndex((i) => i.id === item.id);
		if (idx === -1) {
			imgList = { items: [item, ...imgList.items] };
			return;
		}
		const newItems = [...imgList.items];
		newItems[idx] = { ...newItems[idx], ...item };
		imgList = { items: newItems };
	}

	function connectEvents() {
		try {
			socket = new WebSocket(wsUrl('/ws/imageEvents'));
			socket.addEventListener('message', (e) => {
				try {
					const parsed = ImageEventSchema.parse(JSON.parse(String(e.data)));
					const action = parsed.action;
					const id = parsed.actor.id;
					if (['remove', 'delete', 'destroy'].includes(action)) {
						removeFromListById(id);
						return;
					}

					if (action === 'create') {
						getImageById({ id })
							.then((item) => upsertIntoList(item))
							.catch(() => {
								// requestApi already notifies
							});
						return;
					}

					//if (['destroy'].includes(action)) return;
					if (['tag', 'untag'].includes(action)) return;
					if (action === 'prune') return;
					if (action === 'use') updateImageUsedById(id, true);
					if (action === 'unuse') updateImageUsedById(id, false);
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

	onDestroy(() => {
		socket?.close();
	});

	function toggleSelected(id: string, checked: boolean) {
		selectedIds = checked ? [...selectedIds, id] : selectedIds.filter((x) => x !== id);
	}

	function selectAll(checked: boolean) {
		selectedIds = checked ? imgList.items.map((i) => i.id) : [];
		const elements = document.querySelectorAll('input[name=checkbox-item]');
		Array.prototype.forEach.call(elements, (el: HTMLInputElement) => {
			el.checked = checked;
		});
	}

	async function doRemove(force: boolean) {
		const ids = [...selectedIds];
		selectedIds = [];
		try {
			await Promise.all(ids.map((id) => removeImage({ id, force })));
		} catch {
			// requestApi already notifies
		} finally {
			await refresh();
		}
	}
</script>

<svelte:head>
	<title>Images</title>
</svelte:head>

<Card>
	<PageHeader title="Images" subtitle="Docker images">
		<button type="button" onclick={refresh} disabled={loading}>Refresh</button>
		<button type="button" onclick={() => doRemove(false)} disabled={!canRemove}>Remove</button>
		<button type="button" onclick={() => doRemove(true)} disabled={!canForceRemove}>
			Force remove
		</button>
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
					<th>Tag</th>
					<th>Id</th>
					<th>Size</th>
					<th>Created</th>
				</tr>
			</thead>
			<tbody>
				{#each imgList.items as image (image.id)}
					<tr>
						<td>
							<input
								type="checkbox"
								name="checkbox-item"
								checked={selectedIds.includes(image.id)}
								onchange={(e) => toggleSelected(image.id, (e.target as HTMLInputElement).checked)}
							/>
						</td>
						<td>
							<span>{image.tags.join(', ')}</span>
							{#if !image.is_used}
								<Badge variant="warn">unused</Badge>
							{/if}
						</td>
						<td>{formatId(image.id)}</td>
						<td>{sizeFormat(image.size)}</td>
						<td>{unixTimeFormat(image.created)}</td>
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
