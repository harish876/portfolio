
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
          "geist": ["geist"],
          "nvim": ["JetBrainsMono","Nerd Font","Monospace"],
        }, 
        colors:{
        primary: {DEFAULT: "#252434"},
        secondary: {...colors.emerald , DEFAULT: colors.emerald["700"]},
        terminal: {
          base:"#1e1d2d",
          hover: "#252434"
        },
        error:{
          base: "#F18AA7"
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
  
  
