
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
        primary: {DEFAULT: "#252434", accent: "#2d2c3c", dark:"#191828"},
        secondary: {
            ...colors.emerald , 
          DEFAULT: colors.emerald["700"], 
          normal: "#8ac2f0",
          insert: "#c7a0dc",
          visiual: ""

        },
        terminal: {
          base:"#1e1d2d",
          hover: "#252434"
        },
        error:{
          base: "#F18AA7"
        },
        syntax:{
         tag: "#d7c6a2",
         content:"#c8748f",
         class:"#91c399",
         templ:"#c9a4f5",
         function: "#88b2f7"
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
  
  
