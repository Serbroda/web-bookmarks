import type { BaseDto } from "./BaseDto";

export interface UserDto extends BaseDto<number> {
    username: string;
    name?: string;
    email?: string;
}

export interface ChangePasswordDto {
    oldPassword: string;
    newPassword: string;
}
