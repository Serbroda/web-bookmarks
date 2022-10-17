import { writable } from "svelte/store";
import type { GroupDto } from "../models/GroupDto";

export const groups = writable<GroupDto[]>([]);
