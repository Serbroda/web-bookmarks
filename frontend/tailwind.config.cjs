/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['./src/**/*.{html,js,svelte,ts}'],
    theme: {
        extend: {},
    },
    daisyui: {
        themes: [
          {
            ragbaglight: {
                "primary": "#F7494A",
                "secondary": "#00447A",
                "accent": "#1C7777",
                "neutral": "#191D24",
                "base-100": "#ffffff",
                "info": "#3ABFF8",
                "success": "#36D399",
                "warning": "#FBBD23",
                "error": "#F87272",
            },
            ragbagdark: {
                "primary": "#F7494A",
                "secondary": "#00447A",
                "accent": "#1C7777",
                "neutral": "#191D24",
                "base-100": "#171212",
                "info": "#3ABFF8",
                "success": "#36D399",
                "warning": "#FBBD23",
                "error": "#F87272",
            },
          },
        ],
      },
    plugins: [
        require('@tailwindcss/forms'),
        require('@tailwindcss/line-clamp'),
        require('daisyui'),
    ],
}
