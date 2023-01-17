import {AuthMethod, AuthService, isOAuth2Method} from "./auth/auth.service";
import {AuthApi, AuthLoginPostRequest, ControllersLoginResponse, ModelsUser} from "../gen";
import {AuthListener, AuthStore} from "./auth/auth.store";

export class ApiAuthService implements AuthService<ModelsUser> {
    listeners: AuthListener<ModelsUser>[] = [];

    constructor(private auth: AuthApi, private store: AuthStore<ModelsUser>) {
        store.onChange((authenticated, user) => {
            for (const l of this.listeners) {
                l(authenticated, user);
            }
        })
    }

    async authRefresh(): Promise<ModelsUser | undefined | null> {
        const refreshToken = this.store.refreshToken
        if (refreshToken === undefined || refreshToken === null) {
            throw new Error("Refresh token not set");
        }

        const response = await this.auth.authRefreshTokenPost({
            refreshtoken: {refreshToken}
        });
        this.applyToken(response);
        return response.user;
    }

    async authenticate(auth: AuthMethod): Promise<ModelsUser | undefined | null> {
        if (isOAuth2Method(auth)) {
            throw new Error("OAuth2 is not implemented");
        } else {
            const response = await this.auth.authLoginPost({
                login: {
                    username: auth.username,
                    password: auth.password
                }
            });
            this.applyToken(response);
            return response.user;
        }
    }

    isAuthenticated(): boolean {
        return this.store.authenticated;
    }

    logout(): void {
        this.store.clear();
    }

    onChange(listener: AuthListener<ModelsUser>): void {
        if (!this.listeners.includes(listener)) {
            this.listeners.push(listener);
        }
    }

    async user(): Promise<ModelsUser | undefined | null> {
        return this.store.user;
    }

    private applyToken(tokens: ControllersLoginResponse) {
        if (tokens.accessToken) {
            this.store.accessToken = tokens.accessToken;
        }
        if (tokens.refreshToken) {
            this.store.refreshToken = tokens.refreshToken;
        }
        if (tokens.user) {
            this.store.user = tokens.user;
        }
    }
}