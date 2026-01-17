<script lang="ts" context="module">
	import { RemoveVolumeById } from './api/api.svelte';

	export function RemoveVolume(ids: string[], force: boolean) {
		ids.forEach((id: string): void => {
			RemoveVolumeById(id, force);
		});
	}
</script>
