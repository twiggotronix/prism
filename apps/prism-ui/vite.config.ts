/// <reference types='vitest' />
import { nxViteTsPaths } from "@nx/vite/plugins/nx-tsconfig-paths.plugin";
import react from "@vitejs/plugin-react";
import { defineConfig, loadEnv } from "vite";

export default defineConfig(({ mode }) => {
    const env = loadEnv(mode, process.cwd(), "");
    return {
        root: __dirname,
        cacheDir: "../../node_modules/.vite/apps/prism",

        server: {
            port: 4201,
            host: "localhost",
        },

        preview: {
            port: 4300,
            host: "localhost",
        },

        define: {
            "process.env.BASE_API": JSON.stringify(env.BASE_API),
        },

        plugins: [react(), nxViteTsPaths()],

        // Uncomment this if you are using workers.
        // worker: {
        //  plugins: [ nxViteTsPaths() ],
        // },

        build: {
            outDir: "../../dist/apps/prism",
            reportCompressedSize: true,
            commonjsOptions: {
                transformMixedEsModules: true,
            },
        },

        test: {
            globals: true,
            cacheDir: "../../node_modules/.vitest",
            setupFiles: "./tests/vitest.config.ts",
            environment: "jsdom",
            include: ["src/**/*.{test,spec}.{js,mjs,cjs,ts,mts,cts,jsx,tsx}"],

            reporters: ["default"],
            coverage: {
                reportsDirectory: "../../coverage/apps/prism",
                provider: "v8",
            },
        },
    };
});
