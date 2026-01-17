<script lang="ts">
	import { onDestroy } from 'svelte';
	import Card from '$lib/shared/ui/Card.svelte';
	import PageHeader from '$lib/shared/ui/PageHeader.svelte';
	import { wsUrl } from '$lib/core/http/client';
	import { notifyError } from '$lib/app/notifications/store';
	import { createFromGit } from '../api';
	import type { DeployFromGitInfo } from '../types';
	import '@xterm/xterm/css/xterm.css';
	import { Terminal } from '@xterm/xterm';
	import { FitAddon } from '@xterm/addon-fit';

	let repoUrl = '';
	let overrideDockerfile = false;
	let dockerfileContent = '';
	let overrideCompose = false;
	let composeContent = '';
	let runCommand = '';

	let loading = false;

	let ws: WebSocket | null = null;
	let term: Terminal | null = null;
	let fitAddon: FitAddon | null = null;
	let terminalEl: HTMLDivElement | null = null;
	let wsReady = false;
	let terminalReady = false;

	function initTerminal() {
		term = new Terminal({
			cursorBlink: true,
			theme: { background: '#1e1e1e', foreground: '#f8f8f2' },
			rows: 22,
			fontSize: 12,
			fontFamily: 'monospace',
			disableStdin: true,
			scrollback: 5000,
			convertEol: true
		});

		fitAddon = new FitAddon();
		term.loadAddon(fitAddon);
		term.open(terminalEl!);
		requestAnimationFrame(() => {
			fitAddon!.fit();
			terminalReady = true;
		});
	}

	function connectWS(buildID: string) {
		ws = new WebSocket(wsUrl('/ws/buildLogs', { buildID }));

		ws.addEventListener('open', () => {
			wsReady = true;
		});

		ws.addEventListener('message', async (e) => {
			if (!term || !terminalReady) return;
			const raw = await wsMessageToString(e.data);
			try {
				const parsed = JSON.parse(raw) as { line?: unknown };
				const line = typeof parsed.line === 'string' ? parsed.line : raw;
				term.write(line);
			} catch {
				term.write(raw);
			}
		});

		ws.addEventListener('close', () => {
			loading = false;
		});

		ws.addEventListener('error', (err) => {
			notifyError(err, 'WebSocket error');
			loading = false;
		});
	}

	async function wsMessageToString(data: unknown): Promise<string> {
		if (typeof data === 'string') return data;
		if (data instanceof Blob) return await data.text();
		if (data instanceof ArrayBuffer) return new TextDecoder().decode(new Uint8Array(data));
		if (ArrayBuffer.isView(data)) return new TextDecoder().decode(data);
		return String(data);
	}

	async function handleSubmit(e: Event) {
		e.preventDefault();
		if (loading || !repoUrl.trim()) return;

		loading = true;
		wsReady = false;
		terminalReady = false;
		term?.clear();

		const payload: DeployFromGitInfo = {
			url: repoUrl,
			dockerfile: overrideDockerfile ? dockerfileContent : '',
			docker_compose: overrideCompose ? composeContent : '',
			docker_run: runCommand
		};

		try {
			const res = await createFromGit(payload);
			connectWS(res.build_id);
		} catch (err) {
			loading = false;
			notifyError(err, 'Не удалось запустить сборку');
		}
	}

	$: if (wsReady && terminalEl && term === null) {
		initTerminal();
	}

	onDestroy(() => {
		ws?.close();
		term?.dispose();
	});
</script>

<svelte:head>
	<title>Create from git</title>
</svelte:head>

<Card>
	<PageHeader title="Create from git" subtitle="Build and run from repository" />

	<form class="form" on:submit|preventDefault={handleSubmit}>
		<label for="repo-url">Repository URL</label>
		<input
			id="repo-url"
			type="url"
			bind:value={repoUrl}
			required
			placeholder="https://github.com/..."
		/>

		<label class="inline">
			<input type="checkbox" bind:checked={overrideDockerfile} />
			Redefine Dockerfile
		</label>
		{#if overrideDockerfile}
			<textarea
				bind:value={dockerfileContent}
				rows="10"
				required
				placeholder="Contents of the Dockerfile..."
			></textarea>
		{/if}

		<label class="inline">
			<input type="checkbox" bind:checked={overrideCompose} />
			Redefine docker-compose.yml
		</label>
		{#if overrideCompose}
			<textarea
				bind:value={composeContent}
				rows="10"
				required
				placeholder="Contents of the docker-compose.yml..."
			></textarea>
		{/if}

		<label for="run-cmd">docker run command</label>
		<input id="run-cmd" type="text" bind:value={runCommand} placeholder="docker run …" />

		<button type="submit" disabled={loading || !repoUrl.trim()}>
			{#if loading}Creating…{:else}Create container{/if}
		</button>
	</form>

	{#if wsReady}
		<div class="terminal-wrapper">
			<div bind:this={terminalEl} class="terminal"></div>
		</div>
	{/if}
</Card>

<style>
	.form {
		display: flex;
		flex-direction: column;
		gap: 12px;
		margin-top: 14px;
	}

	.inline {
		display: flex;
		gap: 10px;
		align-items: center;
		color: var(--fg-muted);
	}

	textarea {
		min-height: 150px;
		font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, 'Liberation Mono',
			'Courier New', monospace;
	}

	.terminal-wrapper {
		margin-top: 14px;
		border: 1px solid var(--border);
		border-radius: 12px;
		overflow: hidden;
		background: #1e1e1e;
		height: 320px;
	}
	.terminal {
		width: 100%;
		height: 100%;
	}
</style>
