import { authService } from "../services/Services";
import { readable, writable } from "svelte/store";
import type { UserDto } from "../models/UserDto";

export const user = readable<UserDto>({ username: undefined }, (set) => {
    authService.getUser().then(set);
});

export const token = writable<string | undefined>(undefined, (set) => {
    set(authService.getToken())
});

export const authenticated = writable<boolean>(false, (set) => {
    set(authService.isLoggedIn())
});
