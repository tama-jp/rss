// import adapter from '@sveltejs/adapter-auto';
import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://kit.svelte.dev/docs/integrations#preprocessors
	// for more information about preprocessors
	preprocess: vitePreprocess(),

	kit: {
		// adapter-auto only supports some environments, see https://kit.svelte.dev/docs/adapter-auto for a list.
		// If your environment is not supported, or you settled on a specific environment, switch out the adapter.
		// See https://kit.svelte.dev/docs/adapters for more information about adapters.
		// adapter: adapter({
		// 	fallback: 'index.html'
		// })
		// 静的サイト用アダプターの設定
		adapter: adapter({
			// 出力先のディレクトリを指定
			// pages: '../internal/frameworks/routing/static',  // HTMLファイルの出力先
			// assets: '../internal/frameworks/routing/static',  // 静的アセットの出力先
			pages: '../static',  // HTMLファイルの出力先
			assets: '../static',  // 静的アセットの出力先

			fallback: 'index.html' // SPAのルーティングに使うために必要です
		}),

		// ページの事前レンダリング設定
		prerender: {
			// '*' はすべてのルートをプリレンダリングすることを意味します
			entries: ['*']
		}
	}
};

export default config;
