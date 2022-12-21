import {ID} from "./base";

export type SpaceVisibility = 'PRIVATE' | 'MANAGED' | 'PUBLIC';

export type Space = {
    id: ID;
    name: string;
    description?: string;
    visibility: SpaceVisibility;
}


export type NewSpace = {
    name: string;
    description?: string;
    visibility: SpaceVisibility;
}

export type UpdateSpace = {
    name?: string;
    description?: string;
    visibility?: SpaceVisibility;
};
