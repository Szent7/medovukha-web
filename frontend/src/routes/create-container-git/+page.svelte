<script lang="ts">
	import { onDestroy } from 'svelte';
	import FrameElement from '@templates/frameElement.svelte';
	import Sidebar from '@templates/sidebar.svelte';
	import Frame from '@templates/frame.svelte';
	import logo from '@assets/logo_small.svg';
	import { CreateFromGit } from '@/lib/api/api.svelte';
	import type { DeployFromGitInfo } from '@/lib/api/types.svelte';
	import { Terminal } from '@xterm/xterm';
	import { FitAddon } from '@xterm/addon-fit';

	let repoUrl = '';
	let overrideDockerfile = false;
	let dockerfileContent = '';
	let overrideCompose = false;
	let composeContent = '';
	let runCommand = '';

	let loading = false;
	let errorMsg = '';

	let ws: WebSocket | null = null;
	let term: Terminal | null = null;
	let fitAddon: FitAddon | null = null;
	let terminalEl: HTMLDivElement | null = null;
	let wsReady = false;
	let terminalReady = false;

	async function connectWS(buildID: string) {
		ws = new WebSocket(`ws://localhost:10015/ws/buildLogs?buildID=${buildID}`);

		ws.addEventListener('open', () => {
			console.log('WS connected');
			wsReady = true;
		});

		ws.addEventListener('message', (e) => {
			if (!term || !terminalReady) return;

			const raw = typeof e.data === 'string' ? e.data : new TextDecoder().decode(e.data);
			try {
				const { line } = JSON.parse(raw);

				if (typeof line === 'string') {
					term.write(line + '\r\n');
				}
			} catch {
				term.write(raw);
			}
		});

		ws.addEventListener('close', () => {
			console.log('WS closed');
			loading = false;
		});

		ws.addEventListener('error', (err) => {
			console.error('WS error', err);
			errorMsg = 'WebSocket error: ' + err;
			loading = false;
		});
	}

	async function handleSubmit(e: Event) {
		e.preventDefault();

		if (loading || !repoUrl.trim()) return;

		loading = true;
		errorMsg = '';

		const payload: DeployFromGitInfo = {
			url: repoUrl,
			dockerfile: overrideDockerfile ? dockerfileContent : '',
			docker_compose: overrideCompose ? composeContent : '',
			docker_run: runCommand
		};

		try {
			const response = await CreateFromGit(payload);
			if (!response) {
				throw new Error('No buildID returned from server');
			}

			connectWS(response.build_id);
		} finally {
		}
	}

	$: if (wsReady && terminalEl && term === null) {
		initTerminal();
	}

	$: if (term) {
		term.attachCustomKeyEventHandler(() => false);
	}

	function initTerminal() {
		term = new Terminal({
			cursorBlink: true,
			theme: { background: '#1e1e1e', foreground: '#f8f8f2' },
			rows: 24,
			cols: 80,
			fontSize: 12,
			fontFamily: 'monospace',
			cursorStyle: 'block',
			convertEol: true,
			scrollback: 5000,
			disableStdin: true
		});

		fitAddon = new FitAddon();
		term.loadAddon(fitAddon);
		term.open(terminalEl!);

		requestAnimationFrame(() => {
			fitAddon!.fit();
			terminalReady = true;
		});
	}

	onDestroy(() => {
		if (ws) ws.close();
		if (term) term.dispose();
	});
</script>

<svelte:head>
	<title>Create from git</title>
	<link rel="icon" type="image/svg+xml" href={logo} />
</svelte:head>

<Sidebar />
<Frame {content} />

{#snippet content()}
	<FrameElement
		--display="flex"
		--justify-content="space-between"
		--align-items="center"
		content={header}
	/>
	<FrameElement content={createFromGitForm} />
	{#if errorMsg}
		<FrameElement content={errorMsgData} />
	{/if}

	{#if wsReady}
		<FrameElement content={termWindows} />
	{/if}
{/snippet}

{#snippet termWindows()}
	<div class="terminal-wrapper" hidden={!wsReady}>
		<div bind:this={terminalEl} class="terminal"></div>
	</div>
{/snippet}

{#snippet header()}
	<img class="logoHeader" src={logo} alt="logo" />
	<h2>Create from git</h2>
	<div class="releaseLabel">Alpha v0.0.1</div>
{/snippet}

{#snippet errorMsgData()}
	<p>{errorMsg}</p>
{/snippet}

{#snippet createFromGitForm()}
	<form on:submit|preventDefault={handleSubmit} class="create-form">
		<label for="repo-url">repository URL</label>
		<input
			id="repo-url"
			type="url"
			bind:value={repoUrl}
			required
			placeholder="https://github.com/…"
		/>

		<label>
			<input type="checkbox" bind:checked={overrideDockerfile} />
			Redefine Dockerfile
		</label>
		{#if overrideDockerfile}
			<textarea
				bind:value={dockerfileContent}
				placeholder="Contents of the Dockerfile..."
				rows="10"
				required
			></textarea>
		{/if}

		<label>
			<input type="checkbox" bind:checked={overrideCompose} />
			Redefine docker‑compose.yml
		</label>
		{#if overrideCompose}
			<textarea
				bind:value={composeContent}
				placeholder="Contents of the docker‑compose.yml..."
				rows="10"
				required
			></textarea>
		{/if}

		<label for="run-cmd">docker run command</label>
		<input id="run-cmd" type="text" bind:value={runCommand} placeholder="docker run …" />

		<button type="submit" disabled={loading || !repoUrl.trim()}>
			{#if loading}
				Creating…
			{:else}
				Create container
			{/if}
		</button>
	</form>
{/snippet}

<style>
	h2 {
		font-size: 24pt;
	}

	.releaseLabel {
		background-color: #ffffff;
		color: #1c1c1c;
		padding: 5px;
		border-radius: 5px;
		font-size: 14pt;
		font-weight: bold;
	}

	.logoHeader {
		display: block;
		max-height: 100px;
	}

	.create-form {
		display: flex;
		flex-direction: column;
		gap: 12px;
	}

	.create-form label {
		font-weight: 500;
	}

	.create-form input,
	.create-form textarea {
		padding: 8px;
		font-size: 1rem;
		border: 1px solid #ccc;
		border-radius: 4px;
	}

	.create-form button {
		align-self: flex-start;
	}

	.terminal-wrapper {
		margin-top: 1rem;
		border: 1px solid #444;
		border-radius: 4px;
		overflow: hidden;
		background: #1e1e1e;
		height: 300px;
	}
	.terminal {
		width: 100%;
		height: 100%;
	}

	:global(.xterm-helper-textarea) {
		all: unset;
		position: absolute;
		opacity: 0;
		pointer-events: none;
	}
</style>
