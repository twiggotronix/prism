const { createGlobPatternsForDependencies } = require("@nx/react/tailwind");
const { join } = require("path");

/** @type {import('tailwindcss').Config} */
const withMT = require("@material-tailwind/react/utils/withMT");

module.exports = withMT({
    content: [
        "../../node_modules/@material-tailwind/react/components/**/*.{js,ts,jsx,tsx}",
        "../../node_modules/@material-tailwind/react/theme/components/**/*.{js,ts,jsx,tsx}",
        join(
            __dirname,
            "{src,pages,components,app}/**/*!(*.stories|*.spec).{ts,tsx,html}",
        ),
        ...createGlobPatternsForDependencies(__dirname),
    ],
    theme: {
        fontFamily: {
            sans: ["Roboto", "sans-serif"],
            serif: ["Roboto Slab", "serif"],
            body: ["Roboto", "sans-serif"],
        },
        extend: {},
    },
    plugins: [
        require("tailwind-fontawesome")({
            family: "sharp",
        }),
    ],
});
