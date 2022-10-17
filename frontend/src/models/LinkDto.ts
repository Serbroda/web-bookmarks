import type { BaseDto } from "./BaseDto";

export enum LinkVisibility {
    PRIVATE = "private",
    PUBLIC = "public",
}

export interface LinkDto extends BaseDto<string> {
    name: string;
    url: string;
    description?: string;
    visibility: LinkVisibility;
    groupId?: string;
}

export interface CreateLinkDto {
    name?: string;
    url: string;
    description?: string;
    visibility?: LinkVisibility;
    groupId?: string;
}
