import create from "zustand";

export type Theme = "dark" | "light";

export type PreferencesState = {
    theme: string;
    init: () => void;
    setTheme : (theme: Theme) => void;
}

const applyTheme = (theme : Theme) => {
    if (theme === "dark") {
        document.documentElement.classList.add("dark");
    } else {
        document.documentElement.classList.remove("dark");
    }
}

const usePreferences = create<PreferencesState>((set) => ({
    theme: "light",
    init: () => {
        let theme: Theme | null = null;
        if(localStorage.theme === "dark" || !("theme" in localStorage) && window.matchMedia("(prefers-color-scheme: dark").matches) {
            theme = "dark"
        } else {
            theme = "light"
        }
        applyTheme(theme);
        set({theme})
    },

    setTheme: (theme: Theme) => {
        localStorage.setItem("theme", theme);
        applyTheme(theme);
        set({theme});
    }
  }))

export default usePreferences;