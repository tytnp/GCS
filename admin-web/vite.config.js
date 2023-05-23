import {defineConfig} from 'vite'

import path from "path";
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import {NaiveUiResolver} from 'unplugin-vue-components/resolvers'
import {ViconsResolver} from './vicons_resolver.js'

export default defineConfig({
    build:{
        assetsDir: "admin-assets",
        outDir:"../admin-service/resources/admin-dist"
    },
    resolve: {
        alias: {
            "@": path.resolve(__dirname, "src"),
        }
    },
    plugins: [
        vue(),
        AutoImport({
            dts: true,
            include: [
                /\.[tj]sx?$/, // .ts, .tsx, .js, .jsx
                /\.vue$/, /\.vue\?vue/, // .vue
                /\.md$/, // .md
            ],
            imports: [
                "vue",
                "vue-router"
            ],
            resolvers: [NaiveUiResolver(), ViconsResolver()],
        }),
        Components({
            dts: true,
            resolvers: [NaiveUiResolver(), ViconsResolver()]
        })
    ],
    css: {
        preprocessorOptions: {
            less: {
                javascriptEnabled: true,
                additionalData: `@import "@/config/style/global.less";`,
            },
        },
    },
})
