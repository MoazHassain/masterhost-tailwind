/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["**/*.{html, js}", "**/**/*.{html, js}", "**/**/**/*.{html, js}", "./index.html"],
  theme: {
    extend: {
      colors: {
        blue: {
          '1000': "#010c38",
        },
      },
      fontFamily: {
        roboto: ["Roboto", "sans-serif"],
      },
      container: {
        center: true,
        padding: "2rem",
        
      },
    },
  },
  plugins: [],
}
