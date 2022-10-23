import type { BaseDto } from "./BaseDto";
import type { GroupDto } from "./GroupDto";

export interface GroupSubscriptionDto extends BaseDto<number> {
    group: GroupDto;
}
