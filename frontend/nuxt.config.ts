// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	compatibilityDate: '2025-05-15',
	devtools: { enabled: true },
	runtimeConfig: {
		public: {
			apiBase: process.env.API_URL || 'http://localhost:8000/api'
		},
	},
	modules: [
		'@nuxt/eslint',
		'@nuxt/fonts',
		'@nuxt/icon',
		'@nuxt/image',
		'@nuxtjs/tailwindcss'
	]
})
