import { ApiService } from "./ApiService";
import type { GroupDto, CreateGroupDto, GroupVisibility } from "../models/GroupDto";
import type { CreateLinkDto, LinkDto } from "../models/LinkDto";
import type { GroupSubscriptionDto } from "../models/GroupSubscriptionDto";

export class GroupService extends ApiService {
    constructor(baseUrl: string) {
        super(baseUrl);
    }

    async getGroups(): Promise<GroupDto[]> {
        const response = await this.get(`/api/v1/groups`);
        return response.json();
    }

    async getLatestGroups(order?: string, limit?: string): Promise<GroupDto[]> {
        let url = `/api/v1/groups/latest`;
        if (order || limit) {
            url =
                `${url}?` +
                new URLSearchParams({
                    order: order,
                    limit: limit,
                }).toString();
        }
        const response = await this.get(url);
        return response.json();
    }

    async getGroup(id: string): Promise<GroupDto | undefined> {
        const response = await this.get(`/api/v1/groups/${id}`);
        if (response.status === 404) {
            return undefined;
        }
        return response.json();
    }

    async createGroup(dto: CreateGroupDto): Promise<GroupDto> {
        const response = await this.post(`/api/v1/groups`, dto);
        return response.json();
    }

    async updateGroup(id: string, dto: CreateGroupDto): Promise<GroupDto> {
        const response = await this.patch(`/api/v1/groups/${id}`, dto);
        return response.json();
    }

    async deleteGroup(id: string): Promise<Response> {
        return this.delete(`/api/v1/groups/${id}`);
    }

    async getLinks(groupId: string, order?: string, limit?: string): Promise<LinkDto[]> {
        let url = `/api/v1/groups/${groupId}/links`;
        if (order || limit) {
            url =
                `${url}?` +
                new URLSearchParams({
                    order: order,
                    limit: limit,
                }).toString();
        }
        const response = await this.get(url);
        return response.json();
    }

    async getLatestLinks(order?: string, limit?: string): Promise<LinkDto[]> {
        let url = `/api/v1/links`;
        if (order || limit) {
            url =
                `${url}?` +
                new URLSearchParams({
                    order: order,
                    limit: limit,
                }).toString();
        }
        const response = await this.get(url);
        return response.json();
    }

    async createLink(groupId: string, dto: CreateLinkDto): Promise<LinkDto> {
        const response = await this.post(`/api/v1/groups/${groupId}/links`, dto);
        return response.json();
    }

    async updateLink(linkId: string, dto: CreateLinkDto): Promise<LinkDto> {
        const response = await this.patch(`/api/v1/links/${linkId}`, dto);
        return response.json();
    }

    async deleteLink(id: string): Promise<Response> {
        return this.delete(`/api/v1/links/${id}`);
    }

    async changeGroupVisibility(id: string, visibility: GroupVisibility): Promise<Response> {
        return this.put(`/api/v1/groups/${id}/visibility`, {
            visibility
        });
    }

    async getGroupSubscriptions(): Promise<GroupSubscriptionDto[]> {
        const response = await this.get(`/api/v1/groups/subscriptions`);
        return response.json();
    }

    async createGroupSubscription(groupId: string): Promise<GroupSubscriptionDto> {
        const response = await this.post(`/api/v1/groups/subscriptions/${groupId}`);
        return response.json();
    }

    async deleteGroupSubscription(groupId: string): Promise<Response> {
        return this.delete(`/api/v1/groups/subscriptions/${groupId}`);
    }

}
