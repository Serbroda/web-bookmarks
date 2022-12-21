import {ID} from "./base";

export type Page = {
    id: ID;
    name: string;
    description?: string;
}

export type NewPage = {
    name: string;
    description?: string;
}

export type UpdatePage = {
    spaceId?: ID;
    name?: string;
    description?: string;
};
