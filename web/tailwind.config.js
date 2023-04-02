/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx}",
    "./src/components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
    colors: {
      dark: '#1f2937',
      light: '#f5f5f5',
      primary: "#00a896",
      secondary: "#006d77",
      accent: "#ff6b6b",
    },
  },
  plugins: [],
}

