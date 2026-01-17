<script lang="ts" context="module">
	export function UnixTimeFormat(unixTimestamp: number): string {
		const date = new Date(unixTimestamp * 1000);
		return date.toLocaleString();
	}

	export function SizeConvert(bytes: number): string {
		if (bytes === 0) return '0 Bytes';

		const k = 1024;
		const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
		const i = Math.floor(Math.log(bytes) / Math.log(k));
		return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
	}

	export function FormatId(id: string): string {
		const PREFIX = 'sha256:';

		if (id.startsWith(PREFIX)) {
			id = id.slice(PREFIX.length);
		}

		if (id.length <= 16) {
			return id;
		}

		const start = id.slice(0, 12); // первые 8
		//const end = id.slice(-4); // последние 4

		//return `${start}...${end}`;
		return start;
	}
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

	// export function UnixTimeFormat(unixTimestamp: string): string {
	// 	let parsed: number = parseInt(unixTimestamp);
	// 	if (Number.isNaN(parsed)) {
	// 		return 'unknown';
	// 	}
	// 	const date = new Date(parsed * 1000);
	// 	return date.toLocaleString();
	// }
</script>
