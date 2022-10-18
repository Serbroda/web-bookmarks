import { writable } from "svelte/store";
import type { UserDto } from "../models/UserDto";

export const user = writable<UserDto>();
