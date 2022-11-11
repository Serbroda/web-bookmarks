import create from "zustand";
import {combine} from "zustand/middleware";

export interface PreferencesState {
    theme: string;
}

export type Theme = "dark" | "light";

const applyTheme = (theme : Theme) => {
    console.log(theme)
    if (theme === "dark") {
        document.documentElement.classList.add("dark");
    } else {
        document.documentElement.classList.remove("dark");
    }
}

const usePreferences = create(
    combine({
        theme: "light"
    }, (set, get) => ({

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
);

export default usePreferences;