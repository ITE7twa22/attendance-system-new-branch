// tailwind.config.js
module.exports = {
  content: [
    './index.html',
    './src/**/*.{vue,js,ts,jsx,tsx}',
  ],
  theme: {
    extend: {
      gridRow: {
        'span-16': 'span 16 / span 16',
      },
      fontFamily: {
        'mizan': ['Mizan AR+LT', 'sans-serif'], // Add your custom font here
      },
    },
  },
  plugins: [
    // Additional plugins if any
  ],
}