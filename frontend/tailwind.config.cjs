/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: "jit",
  darkMode: "class",
  content: ["./index.html", "./src/**/*.{js,jsx,ts,tsx}"],
  theme: {},
  plugins: [require("@tailwindcss/forms"), require("@tailwindcss/line-clamp")],
};
