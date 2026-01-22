import { derived, writable } from 'svelte/store';
import { ApiError, HttpError, ParseError, SchemaError } from '$lib/core/errors';

export type NotificationLevel = 'info' | 'success' | 'warn' | 'error';

export type Notification = {
	id: string;
	level: NotificationLevel;
	title: string;
	message: string;
	details?: string;
	createdAt: number;
	dismissedAt?: number;
	readAt?: number;
};

function id(): string {
	return `${Date.now().toString(36)}-${Math.random().toString(36).slice(2, 10)}`;
}

const notificationsStore = writable<Notification[]>([]);
export const notifications = { subscribe: notificationsStore.subscribe };

const centerOpenStore = writable(false);
export const notificationCenterOpen = { subscribe: centerOpenStore.subscribe };

export function toggleNotificationCenter(force?: boolean) {
	centerOpenStore.update((v) => (force === undefined ? !v : force));
}

export function notify(input: Omit<Notification, 'id' | 'createdAt'>) {
	const n: Notification = {
		id: id(),
		createdAt: Date.now(),
		...input
	};

	notificationsStore.update((list) => [n, ...list]);

	setTimeout(() => {
		dismissNotification(n.id);
	}, 15_000);

	return n.id;
}

export function dismissNotification(id: string) {
	notificationsStore.update((list) =>
		list.map((n) => (n.id === id && !n.dismissedAt ? { ...n, dismissedAt: Date.now() } : n))
	);
}

export function markRead(id: string) {
	notificationsStore.update((list) =>
		list.map((n) => (n.id === id && !n.readAt ? { ...n, readAt: Date.now() } : n))
	);
}

export function markAllRead() {
	const now = Date.now();
	notificationsStore.update((list) => list.map((n) => (n.readAt ? n : { ...n, readAt: now })));
}

export function clearAll() {
	notificationsStore.set([]);
}

export const activeToasts = derived(notificationsStore, (list) =>
	list.filter((n) => !n.dismissedAt).slice(0, 3)
);

export const unreadCount = derived(
	notificationsStore,
	(list) => list.filter((n) => !n.readAt).length
);

export function notifyError(err: unknown, title = 'Ошибка') {
	const { message, details } = formatError(err);
	notify({ level: 'error', title, message, details });
}

function formatError(err: unknown): { message: string; details?: string } {
	if (err instanceof ApiError) {
		return {
			message: err.message,
			details: `code: ${err.code}`
		};
	}

	if (err instanceof HttpError) {
		const parts = [err.status ? `HTTP ${err.status}` : 'Network error'];
		if (err.method) parts.push(err.method);
		if (err.url) parts.push(err.url);
		return {
			message: parts.join(' '),
			details: err.data ? safeStringify(err.data) : undefined
		};
	}

	if (err instanceof SchemaError) {
		return { message: 'Wrong server response', details: err.message };
	}

	if (err instanceof ParseError) {
		return { message: 'Ошибка обработки ответа сервера', details: err.message };
	}

	if (err instanceof Error) {
		return { message: err.message };
	}

	return { message: 'Неизвестная ошибка', details: safeStringify(err) };
}

function safeStringify(v: unknown): string {
	try {
		return JSON.stringify(v, null, 2);
	} catch {
		return String(v);
	}
}
