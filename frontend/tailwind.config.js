/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        slate: {
          850: '#1e2634',
          950: '#0c1015',
        },
      },
    },
  },
  plugins: [],
}