
const colors = require('tailwindcss/colors')

/** @type {import('tailwindcss').Config} */
module.exports = {
    darkMode: 'selector',
    content: [
      "./components/**/*/*.{templ,html,js}",
      "./views/**/*/*.{templ,html,js}",
    ],
    theme: {
      extend: {
        colors:{
        primary: {...colors.blue , DEFAULT: colors.blue["500"]},
        secondary: {...colors.green , DEFAULT: colors.green["400"]},
        terminal: {
          base:"#202224",
          hover: colors.black
        },
        error:{
          base:colors.red["600"]
        }
      }
      },
    },
    plugins: [
      require('@tailwindcss/forms'),
      // require('daisyui'),
    ],
  }
  
  
