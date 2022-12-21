import {AuthMethod, AuthService, isOAuth2Method} from "./auth.service";
import PocketBase, {Admin, Record} from "pocketbase";

export const USERS_COLLECTION = 'users'

export default class PocketbaseAuthService implements AuthService<Admin | Record> {
    listeners: ((authenticated: boolean, user: Admin | Record | undefined | null) => void)[] = [];

    constructor(private pb: PocketBase) {
        pb.authStore.onChange(async (fn) => {
            const u = await this.user();
            const a = this.isAuthenticated();
            for (const listener of this.listeners) {
                listener(a, u);
            }
        })
    }
    async authenticate(auth: AuthMethod): Promise<Admin | Record> {
        if (isOAuth2Method(auth)) {
            const res = await this.pb.collection(USERS_COLLECTION)
                .authWithOAuth2(
                    auth.provider,
                    auth.code,
                    auth.codeVerifier,
                    auth.redirectUrl,
                    auth.createData
                );
            return res.record;
        } else {
            const res = await this.pb.collection(USERS_COLLECTION)
                .authWithPassword(
                    auth.username,
                    auth.password
                );
            return res.record;
        }
    }

    isAuthenticated(): boolean {
        const token = this.pb.authStore.token;
        const tokenSet = token !== undefined && token !== null && token !== ""
        return tokenSet && this.pb.authStore.isValid;
    }

    async authRefresh(): Promise<Admin | Record | undefined | null> {
        const res = await this.pb.collection(USERS_COLLECTION)
            .authRefresh();
        return res.record;
    }

    async user(): Promise<Admin | Record | undefined | null> {
        return Promise.resolve(this.pb.authStore.model);
    }

    logout(): void {
        this.pb.authStore.clear();
    }

    onChange(listener: (authenticated: boolean, user: (Admin | Record | undefined | null)) => void): void {
        if (!this.listeners.includes(listener)) {
            this.listeners.push(listener);
        }
    }

}
