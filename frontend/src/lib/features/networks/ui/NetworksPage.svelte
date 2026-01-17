<script lang="ts">
	import { onMount } from 'svelte';
	import Card from '$lib/shared/ui/Card.svelte';
	import PageHeader from '$lib/shared/ui/PageHeader.svelte';
	import Badge from '$lib/shared/ui/Badge.svelte';
	import { formatId } from '$lib/shared/utils/format';
	import { getNetworkList, removeNetwork } from '../api';
	import type { ListNetworkBaseInfo } from '../types';

	let netList: ListNetworkBaseInfo = { items: [] };
	let loading = true;
	let selectedIds: string[] = [];

	$: canRemove = selectedIds.length > 0;

	async function refresh() {
		loading = true;
		selectedIds = [];
		try {
			netList = await getNetworkList();
		} catch {
			// requestApi already notifies
		} finally {
			loading = false;
		}
	}

	onMount(refresh);

	function toggleSelected(id: string, checked: boolean) {
		selectedIds = checked ? [...selectedIds, id] : selectedIds.filter((x) => x !== id);
	}

	function selectAll(checked: boolean) {
		selectedIds = checked
			? netList.items.filter((n) => !(n.docker_network || n.is_used)).map((n) => n.id)
			: [];
		const elements = document.querySelectorAll('input[name=checkbox-item]');
		Array.prototype.forEach.call(elements, (el: HTMLInputElement) => {
			if (!el.disabled) el.checked = checked;
		});
	}

	async function doRemove() {
		const ids = [...selectedIds];
		selectedIds = [];
		try {
			await Promise.all(ids.map((id) => removeNetwork({ id })));
		} catch {
			// requestApi already notifies
		} finally {
			await refresh();
		}
	}
</script>

<svelte:head>
	<title>Networks</title>
</svelte:head>

<Card>
	<PageHeader title="Networks" subtitle="Docker networks">
		<button type="button" onclick={refresh} disabled={loading}>Refresh</button>
		<button type="button" onclick={doRemove} disabled={!canRemove}>Remove</button>
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
					<th>Id</th>
					<th>Driver</th>
					<th>IPv6</th>
					<th>IPAM</th>
					<th>Subnet</th>
					<th>Gateway</th>
					<th>Attachable</th>
				</tr>
			</thead>
			<tbody>
				{#each netList.items as network (network.id)}
					<tr>
						<td>
							<input
								type="checkbox"
								name="checkbox-item"
								disabled={Boolean(network.docker_network || network.is_used)}
								checked={selectedIds.includes(network.id)}
								onchange={(e) => toggleSelected(network.id, (e.target as HTMLInputElement).checked)}
							/>
						</td>
						<td>
							<span>{network.name}</span>
							{#if !network.is_used && !network.docker_network}
								<Badge variant="warn">unused</Badge>
							{/if}
							{#if network.docker_network}
								<Badge variant="accent">docker</Badge>
							{/if}
						</td>
						<td>{formatId(network.id)}</td>
						<td>{network.driver}</td>
						<td>{network.enable_ipv6 ?? false}</td>
						<td>{network.ipam_driver}</td>
						<td>{network.subnet?.length ? network.subnet.join('\n') : '-'}</td>
						<td>{network.gateway?.length ? network.gateway.join('\n') : '-'}</td>
						<td>{network.attachable ?? false}</td>
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

	td {
		white-space: pre-line;
	}
</style>
