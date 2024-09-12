/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./views/**/*.html'],
  theme: {
    extend: {
      container: {
        center: true
      },
      fontFamily: {
        firaCode: "'Fira Code', monospace",
      }
    },
  },
  plugins: [],
}

