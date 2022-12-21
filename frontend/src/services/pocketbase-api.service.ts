import PocketBase from "pocketbase";
import {SpaceService} from "./api.service";
import {NewSpace, Space, UpdateSpace} from "../models/space";
import {ID} from "../models/base";

export default class PocketbaseApiService implements SpaceService {
    private readonly spacesCollection = 'spaces';

    constructor(private pb: PocketBase) {
    }

    private getId(): string | null {
        if (!this.pb.authStore.model) {
            return null;
        }
        return this.pb.authStore.model.id;
    }

    async createSpace(body: NewSpace): Promise<Space> {
        const res = await this.pb.collection(this.spacesCollection).create<Space>({
            ...body,
            ...{owner_id: this.getId()}
        })
        return res;
    }

    deleteSpace(id: ID): Promise<void> {
        return Promise.resolve(undefined);
    }

    getSpace(id: ID): Promise<Space | undefined> {
        return Promise.resolve(undefined);
    }

    async getSpaces(): Promise<Space[]> {
        const res = await this.pb.collection(this.spacesCollection).getList<Space>();
        return res.items;
    }

    updateSpace(id: ID, body: UpdateSpace): Promise<Space> {
        throw new Error("Not implemented")
    }

}
