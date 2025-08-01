export default defineNuxtConfig({
    compatibilityDate: '2025-07-15',
    devtools: {enabled: true},
    runtimeConfig: {
        public: {
            apiBaseUrl: process.env.API_DOMAIN || 'localhost',
            apiWsSchema: process.env.API_WS_SCHEMA || 'ws://',
            apiSchema: process.env.API_SCHEMA || 'http://',
        },
    },
    typescript: {
        strict: true,
    },
    vite: {
        server: {
            proxy: {
                '/api': {
                    target: `${process.env.API_SCHEMA}${process.env.API_DOMAIN}`,
                    changeOrigin: true,
                    rewrite: (path) => path.replace(/^\/api/, '')
                },
                // fixme: настроить редирект на /ws
            },
        },
    },
});