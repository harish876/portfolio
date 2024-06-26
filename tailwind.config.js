
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
        fontFamily:{
          "geist": ["geist"]
        },
        colors:{
        primary: {...colors.blue , DEFAULT: colors.blue["500"]},
        secondary: {...colors.emerald , DEFAULT: colors.emerald["700"]},
        terminal: {
          base:" #171717",
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
      require('daisyui'),
      require('flowbite/plugin')
    ],
  }
  
  
