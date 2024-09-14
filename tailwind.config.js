/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./views/**/*.html'],
  theme: {
    extend: {
      container: {
        center: true,
        padding: "16px"
      },
      fontFamily: {
        firaCode: "'Fira Code', monospace",
      },
      colors: {
        bgWhite: '#EEEEEE',
        darkColor: '#222831',
        semiLight: '#76ABAE',
      }
    },
  },
  plugins: [],
}

