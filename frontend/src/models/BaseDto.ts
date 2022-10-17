export interface BaseDto<T extends number | string> {
    id?: T;
    createdAt?: Date;
    updatedAt?: Date;
    deletedAt?: Date;
}
