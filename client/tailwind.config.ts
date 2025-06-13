import type { Config } from "tailwindcss";

export default {
  content: ["./app/**/{**,.client,.server}/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {
      fontFamily: {
        sans: [
          '"Inter"',
          "ui-sans-serif",
          "system-ui",
          "sans-serif",
          '"Apple Color Emoji"',
          '"Segoe UI Emoji"',
          '"Segoe UI Symbol"',
          '"Noto Color Emoji"',
        ],
      },
      colors: {
        gold: {
          100: '#f9f3e8',
          200: '#f6e7c6',
          400: '#e3c675',
          500: '#cfa93f',
          600: '#b2892e',
        },
      },
      dropShadow: {
        glow: '0 0 8px #e3c675',
      },
    },
  },
  plugins: [],
} satisfies Config;
