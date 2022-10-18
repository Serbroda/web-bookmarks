import { ApiService } from "./ApiService";
import type { GroupDto, CreateGroupDto } from "../models/GroupDto";
import type { CreateLinkDto, LinkDto } from "../models/LinkDto";

export class GroupService extends ApiService {
    constructor(baseUrl: string) {
        super(baseUrl);
    }

    async getGroups(order?: string, limit?: string): Promise<GroupDto[]> {
        let url = `/api/v1/groups`;
        if (order || limit) {
            url = `${url}?` + new URLSearchParams({
                order: order, 
                limit: limit
            }).toString()
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

    async getLinks(groupId: string): Promise<LinkDto[]> {
        const response = await this.get(`/api/v1/groups/${groupId}/links`);
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

    async getVersion(): Promise<string> {
        const response = await this.get(`/version`);
        return response.text();
    }
}
