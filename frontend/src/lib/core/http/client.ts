import axios from 'axios';

export const api = axios.create({
	baseURL: '',
	withCredentials: false,
	headers: {
		Accept: 'application/json'
	}
});

export function wsUrl(path: string, params?: Record<string, string>): string {
	const base = new URL(window.location.origin);
	base.protocol = base.protocol === 'https:' ? 'wss:' : 'ws:';
	base.pathname = path.startsWith('/') ? path : `/${path}`;
	if (params) {
		for (const [k, v] of Object.entries(params)) base.searchParams.set(k, v);
	}
	return base.toString();
}
