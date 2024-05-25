import adapter from '@sveltejs/adapter-node';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';
import { mdsvex } from 'mdsvex';
import rehypeAutolinkHeadings from 'rehype-autolink-headings';
import rehypeSlug from 'rehype-slug';

/** @type {import('@sveltejs/kit').Config} */
const config = {
  kit: {
    // adapter-auto only supports some environments, see https://kit.svelte.dev/docs/adapter-auto for a list.
    // If your environment is not supported or you settled on a specific environment, switch out the adapter.
    // See https://kit.svelte.dev/docs/adapters for more information about adapters.
    adapter: adapter(),
  },
  preprocess: [
    vitePreprocess(),
    mdsvex({
      // The default mdsvex extension is .svx; this overrides that.
      extensions: ['.svelte', '.svx', '.md'],

      // Adds IDs to headings, and anchor links to those IDs. Note: must stay in this order to work.
      rehypePlugins: [rehypeSlug, rehypeAutolinkHeadings],
      smartypants: {
        quotes: true,
        ellipses: true,
        backticks: true,
        dashes: 'oldschool',
      },
    }),
  ],
  extensions: ['.svelte', '.svx', '.md'],
};

export default config;
