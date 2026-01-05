<script lang="ts" context="module">
	// export function SizeConvert(bytes: string): string {
	// 	let parsed: number = parseInt(bytes);
	// 	if (Number.isNaN(parsed)) {
	// 		return 'unknown';
	// 	}

	// 	if (parsed === 0) return '0 Bytes';

	// 	const k = 1024;
	// 	const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
	// 	const i = Math.floor(Math.log(parsed) / Math.log(k));
	// 	return parseFloat((parsed / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
	// }

	export function SizeConvert(bytes: number): string {
		if (bytes === 0) return '0 Bytes';

		const k = 1024;
		const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
		const i = Math.floor(Math.log(bytes) / Math.log(k));
		return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
	}
</script>
