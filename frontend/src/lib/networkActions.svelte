<script lang="ts" context="module">
	import { RemoveNetworkById } from './api/api.svelte';

	export function RemoveNetwork(ids: string[]) {
		ids.forEach((id: string): void => {
			RemoveNetworkById(id);
		});
	}
</script>
