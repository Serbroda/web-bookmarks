import type { GroupSubscriptionDto } from "../models/GroupSubscriptionDto";
import { writable } from "svelte/store";
import type { GroupDto } from "../models/GroupDto";

export const groups = writable<GroupDto[]>([]);
export const groupSubscriptions = writable<GroupSubscriptionDto[]>([]);