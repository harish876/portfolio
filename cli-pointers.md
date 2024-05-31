1. Remove unecessary config from the tailwind config.
```js

module.exports = {
    content: [
      // "../**/*.{html,templ,go,js}", //remove this comment
      "./views/**/*/*.{templ,html,js}",
    ],
    theme: {
      extend: {},
    },
    plugins: [
      require('@tailwindcss/forms'),
    ],
  }
```
2. Make the npm command as npm run dev and not npm run start

```json
  "scripts": {
    "dev:tailwind": "npx tailwindcss -i ./input.css -o ./assets/css/styles.css --watch",
    "dev:go": "air -c ./.air.toml",
    "dev": "concurrently \"npm run dev:tailwind\" \"npm run dev:go\""
  }

```

3. Make the not found component as default and display the path for the user to modify it
