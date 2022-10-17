import type { BaseDto } from "./BaseDto";

export enum GroupVisibility {
    PRIVATE = "private",
    PUBLIC = "public",
}

export interface GroupDto extends BaseDto<string> {
    icon: string;
    name: string;
    description: string;
    visibility: GroupVisibility;
}

export interface CreateGroupDto {
    icon: string;
    name: string;
    description?: string;
    visibility?: GroupVisibility;
}
