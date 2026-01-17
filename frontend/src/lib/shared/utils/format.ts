export function unixTimeFormat(unixSeconds: number): string {
	const date = new Date(unixSeconds * 1000);
	return date.toLocaleString();
}

export function sizeFormat(bytes: number): string {
	if (bytes === 0) return '0 Bytes';

	const k = 1024;
	const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
	const i = Math.floor(Math.log(bytes) / Math.log(k));
	return `${parseFloat((bytes / Math.pow(k, i)).toFixed(2))} ${sizes[i]}`;
}

export function formatId(id: string): string {
	const PREFIX = 'sha256:';
	if (id.startsWith(PREFIX)) id = id.slice(PREFIX.length);
	return id.length <= 16 ? id : id.slice(0, 12);
}
