/**
 * Tailwind config for the Go/templ site (standalone CLI v3, see Justfile `css`).
 *
 * @type {import('tailwindcss').Config}
 */
module.exports = {
  content: ['./web/**/*.templ', './web/**/*.go', './web/static/js/**/*.js'],
  theme: {
    extend: {},
  },
  plugins: [],
}
