<script lang="ts" context="module">
	import { RemoveImageById } from './api/api.svelte';

	export function RemoveImage(ids: string[], force: boolean) {
		ids.forEach((id: string): void => {
			RemoveImageById(id, force);
		});
	}
</script>
