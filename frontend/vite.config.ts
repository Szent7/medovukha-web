import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
	server: {
		proxy: {
			'/rest': { target: 'http://localhost:10015', changeOrigin: true },
			'/ping': { target: 'http://localhost:10015', changeOrigin: true },
			'/ws': { target: 'http://localhost:10015', ws: true, changeOrigin: true }
		}
	}
});
