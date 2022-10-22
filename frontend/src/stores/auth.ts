import { authService } from "../services/Services";
import { readable } from "svelte/store";
import type { UserDto } from "../models/UserDto";

export const user = readable<UserDto>({ username: "" }, (set) => {
    authService.getUser().then(set);
});
