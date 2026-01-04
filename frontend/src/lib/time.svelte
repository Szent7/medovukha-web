<script lang="ts" context="module">
	export function UnixTimeFormat(unixTimestamp: string): string {
		let parsed: number = parseInt(unixTimestamp);
		if (Number.isNaN(parsed)) {
			return 'unknown';
		}
		const date = new Date(parsed * 1000);
		return date.toLocaleString();
	}
</script>
