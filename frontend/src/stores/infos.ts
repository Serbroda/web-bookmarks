import { appService } from "../services/Services";
import { readable } from "svelte/store";

export const version = readable<string>("", (set) => {
    appService.getVersion().then(set);
});
