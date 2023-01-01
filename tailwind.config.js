/** @type {import('tailwindcss').Config} */
function withOpacity(colorName) {
  return ({ opacityValue }) => {
    if (opacityValue != undefined) {
      return `rgba(var(${colorName}), ${opacityValue})`
    }
    return `rgba(var(${colorName}), 0.2)`
  }
}

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

      textColor: {
        skin: {
          base: "var(--color-text-base)",
          muted: "var(--color-text-muted)",
          "base-inverted": "var(--color-text-base-inverted)",
          "muted-inverted": "var(--color-text-muted-inverted)",
          message: "var(--color-text-message)",
          hover: "var(--color-text-hover)",
        },
      },
      backgroundColor: {
        skin: {
          base: "var(--color-fill-base)",
          muted: "var(--color-fill-muted)",
          amplify: "var(--color-fill-amplify)",
          "base-inverted": "var(--color-fill-base-inverted)",
          "muted-inverted": "var(--color-fill-muted-inverted)",
          "amplify-inverted": "var(--color-fill-amplify-inverted)",
          "fill-hover": "var(--color-fill-hover)",
          "sidebar-fill": "var(--color-sidebar-fill)",
          "sidebar-fill-active": "var(--color-sidebar-fill-active)",
          "button-base": "var(--color-button-base)",
          "button-base-hover": "var(--color-button-base-hover)",
          "button-inverted": "var(--color-button-inverted)",
          "button-inverted-hover": "var(--color-button-inverted-hover)",
        },
      },
      borderColor: {
        skin: {
          base: "var(--color-border-base)",
          muted: "var(--color-border-muted)",
        },
      },
      outlineColor: {
        skin: {
          base: "var(--color-border-base)",
          muted: "var(--color-border-muted)",
        },
      },
      fill: {
        skin: {
          base: "var(--color-text-base)",
          muted: "var(--color-text-muted)",
          hover: "var(--color-text-hover)",
          inverted: "var(--color-text-muted-inverted)",
          "inverted-hover": "var(--color-text-base-inverted)",
        },
      },
      stroke: {
        skin: {
          base: "var(--color-text-base)",
          muted: "var(--color-text-muted)",
          hover: "var(--color-text-hover)",
          inverted: "var(--color-text-muted-inverted)",
          "inverted-hover": "var(--color-text-base-inverted)",
        },
      },
      boxShadowColor: {
        skin: {
          base: withOpacity("--color-boxshadow-base"),
          muted: withOpacity("--color-boxshadow-muted"),
        }
      },
    },
  },
  plugins: [],
}
