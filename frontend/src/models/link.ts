import {ID} from "./base";

export type Link = {
    id: ID;
    url: string;
    name: string;
    description?: string;
}

export type NewLink = {
    url: string;
    name: string;
    description?: string;
}

export type UpdateLink = {
    pageId?: ID;
    url?: string;
    name?: string;
    description?: string;
};
