// @ts-check
import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

export default defineConfig({
  site: 'https://thruqe.github.io',
  base: '/zpm',
  integrations: [
    starlight({
      title: 'ZPM',
      description: 'A small, sharp process manager for apps you want to keep alive.',
      social: [
        { icon: 'github', label: 'GitHub', href: 'https://github.com/thruqe/zpm' },
      ],
      customCss: ['./src/styles/custom.css'],
      sidebar: [
        {
          label: 'Getting Started',
          items: [
            { label: 'Overview', slug: 'index' },
            { label: 'Installation', slug: 'guides/installation' },
            { label: 'Quick Start', slug: 'guides/quickstart' },
          ],
        },
        {
          label: 'Guides',
          items: [
            { label: 'Clusters', slug: 'guides/clusters' },
            { label: 'Health Checks', slug: 'guides/health-checks' },
            { label: 'Logs & Rotation', slug: 'guides/logs' },
            { label: 'Ecosystem Config', slug: 'guides/ecosystem' },
            { label: 'Daemon & OS Boot', slug: 'guides/daemon' },
          ],
        },
        {
          label: 'Reference',
          items: [
            { label: 'CLI Reference', slug: 'reference/cli' },
            { label: 'Configuration Schema', slug: 'reference/configuration' },
          ],
        },
      ],
    }),
  ],
});