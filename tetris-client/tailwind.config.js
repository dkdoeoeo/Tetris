/** @type {import('tailwindcss').Config} */
export default {
  mode: 'jit',
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        jersey: '"Jersey 10", sans-serif',
        'mono': ' "monospace", monospace',
      },
      fontSize: {
        
      }

    },
  },
  plugins: [],
}

