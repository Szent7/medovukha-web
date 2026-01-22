<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import Card from '$lib/shared/ui/Card.svelte';
	import PageHeader from '$lib/shared/ui/PageHeader.svelte';
	import Badge from '$lib/shared/ui/Badge.svelte';
	import { unixTimeFormat } from '$lib/shared/utils/format';
	import { wsUrl } from '$lib/core/http/client';
	import { notifyError } from '$lib/app/notifications/store';
	import {
		getContainerById,
		getContainerList,
		killContainerById,
		pauseContainerById,
		restartContainerById,
		removeContainerById,
		startContainerById,
		stopContainerById,
		unpauseContainerById
	} from '../api';
	import { ContainerEventSchema, type ListContainerBaseInfo } from '../types';

	let conList: ListContainerBaseInfo = { items: [] };
	let loading = true;
	let selectedIds: string[] = [];

	let socket: WebSocket | null = null;

	const actionToState: Record<string, string> = {
		start: 'running',
		restart: 'running',
		unpause: 'running',
		stop: 'exited',
		kill: 'exited',
		die: 'exited',
		pause: 'paused'
	};

	const actionsBlacklist = new Map<string, string[]>([
		['running', ['start']],
		['paused', ['start', 'stop', 'pause']],
		['exited', ['stop', 'kill', 'pause', 'unpause']]
	]);

	$: selectedStates = selectedIds
		.map((id) => conList.items.find((c) => c.id === id)?.state)
		.filter((s): s is string => Boolean(s));
	$: uniqueStates = Array.from(new Set(selectedStates));
	$: disabledActions = (() => {
		if (selectedIds.length === 0)
			return new Set(['start', 'stop', 'kill', 'restart', 'pause', 'unpause', 'remove']);
		const disabled = new Set<string>();
		for (const st of uniqueStates) {
			for (const act of actionsBlacklist.get(st) ?? []) disabled.add(act);
		}
		return disabled;
	})();

	async function refresh() {
		loading = true;
		selectedIds = [];
		try {
			conList = await getContainerList();
		} catch {
			// requestApi already notifies
		} finally {
			loading = false;
		}
	}

	function setContainerState(id: string, newState: string) {
		const idx = conList.items.findIndex((c) => c.id === id);
		if (idx === -1) return;
		const newItems = [...conList.items];
		newItems[idx] = { ...newItems[idx], state: newState };
		conList = { items: newItems };
	}

	function removeFromListById(id: string) {
		conList = { items: conList.items.filter((c) => c.id !== id) };
		selectedIds = selectedIds.filter((x) => x !== id);
	}

	function upsertIntoList(item: (typeof conList.items)[number]) {
		const idx = conList.items.findIndex((c) => c.id === item.id);
		if (idx === -1) {
			conList = { items: [item, ...conList.items] };
			return;
		}
		const newItems = [...conList.items];
		newItems[idx] = { ...newItems[idx], ...item };
		conList = { items: newItems };
	}

	function connectEvents() {
		try {
			socket = new WebSocket(wsUrl('/ws/containerEvents'));
			socket.addEventListener('message', (e) => {
				try {
					const parsed = ContainerEventSchema.parse(JSON.parse(String(e.data)));
					const action = parsed.action;
					const id = parsed.actor.id;

					if (['remove', 'delete', 'destroy'].includes(action)) {
						removeFromListById(id);
						return;
					}

					if (action === 'create') {
						getContainerById({ id })
							.then((item) => upsertIntoList(item))
							.catch(() => {
								// requestApi already notifies
							});
						return;
					}

					//if (['destroy'].includes(action)) return;
					const newState = actionToState[action];
					if (newState) setContainerState(id, newState);
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

	function toggleSelected(id: string, checked: boolean) {
		selectedIds = checked ? [...selectedIds, id] : selectedIds.filter((x) => x !== id);
	}

	function selectAll(checked: boolean) {
		selectedIds = checked ? conList.items.map((c) => c.id) : [];
		const elements = document.querySelectorAll('input[name=checkbox-item]');
		Array.prototype.forEach.call(elements, (el: HTMLInputElement) => {
			el.checked = checked;
		});
	}

	async function runAction(action: string) {
		const ids = [...selectedIds];
		selectedIds = [];
		const tasks = ids.map((id) => {
			const payload = { id };
			switch (action) {
				case 'start':
					return startContainerById(payload);
				case 'stop':
					return stopContainerById(payload);
				case 'kill':
					return killContainerById(payload);
				case 'restart':
					return restartContainerById(payload);
				case 'pause':
					return pauseContainerById(payload);
				case 'unpause':
					return unpauseContainerById(payload);
				case 'remove':
					return removeContainerById(payload);
				default:
					return Promise.resolve();
			}
		});
		try {
			await Promise.all(tasks);
		} catch {
			// requestApi already notifies
		} finally {
			// refresh()
		}
	}

	function stateBadgeVariant(state: string) {
		if (state === 'running') return 'success';
		if (state === 'paused') return 'warn';
		if (state === 'exited') return 'danger';
		return 'neutral';
	}
</script>

<svelte:head>
	<title>Containers</title>
</svelte:head>

<Card>
	<PageHeader title="Containers" subtitle="Docker containers">
		<button type="button" onclick={refresh} disabled={loading}>Refresh</button>
		<a href="/create-container-git" class="btn">Create from git</a>
		<button type="button" onclick={() => runAction('start')} disabled={disabledActions.has('start')}
			>Start</button
		>
		<button type="button" onclick={() => runAction('stop')} disabled={disabledActions.has('stop')}
			>Stop</button
		>
		<button type="button" onclick={() => runAction('kill')} disabled={disabledActions.has('kill')}
			>Kill</button
		>
		<button
			type="button"
			onclick={() => runAction('restart')}
			disabled={disabledActions.has('restart')}>Restart</button
		>
		<button type="button" onclick={() => runAction('pause')} disabled={disabledActions.has('pause')}
			>Pause</button
		>
		<button
			type="button"
			onclick={() => runAction('unpause')}
			disabled={disabledActions.has('unpause')}>Resume</button
		>
		<button
			type="button"
			onclick={() => runAction('remove')}
			disabled={disabledActions.has('remove')}>Remove</button
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
					<th>State</th>
					<th>Image</th>
					<th>Port bindings</th>
					<th>Created</th>
				</tr>
			</thead>
			<tbody>
				{#each conList.items as container (container.id)}
					<tr>
						<td>
							<input
								type="checkbox"
								name="checkbox-item"
								checked={selectedIds.includes(container.id)}
								onchange={(e) =>
									toggleSelected(container.id, (e.target as HTMLInputElement).checked)}
							/>
						</td>
						<td>{container.names[0]}</td>
						<td>
							<Badge variant={stateBadgeVariant(container.state)}>{container.state}</Badge>
						</td>
						<td>{container.image_name}</td>
						<td>
							{#if !container.ports || container.ports.length === 0}
								-
							{:else}
								{#each container.ports as port}
									{#if port?.public_port !== undefined && port?.private_port !== undefined}
										<div>{port.public_port}:{port.private_port}</div>
									{/if}
								{/each}
							{/if}
						</td>
						<td>{unixTimeFormat(container.created)}</td>
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
</style>
