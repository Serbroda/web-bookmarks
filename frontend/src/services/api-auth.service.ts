import {AuthMethod, AuthService, isOAuth2Method} from "./auth/auth.service";
import {AuthApi, UserDto} from "../gen";
import {AuthListener, AuthStore} from "./auth/auth.store";

export class ApiAuthService implements AuthService<UserDto> {
    listeners: AuthListener<UserDto>[] = [];

    constructor(private auth: AuthApi, private store: AuthStore<UserDto>) {
        store.onChange((authenticated, user) => {
            for (const l of this.listeners) {
                l(authenticated, user);
            }
        })
    }

    async authRefresh(): Promise<UserDto | undefined | null> {
        throw new Error("OAuth2 is not implemented");
    }

    async authenticate(auth: AuthMethod): Promise<UserDto | undefined | null> {
        if (isOAuth2Method(auth)) {
            throw new Error("OAuth2 is not implemented");
        } else {
            const response = await this.auth.login({
                loginDto: {
                    username: auth.username,
                    password: auth.password
                }
            });
            if (response.accessToken) {
                this.store.accessToken = response.accessToken;
            }
            if (response.refreshToken) {
                this.store.refreshToken = response.refreshToken;
            }
            const u = await this.user();
            this.store.user = u;
            return u;
        }
    }

    isAuthenticated(): boolean {
        return this.store.authenticated;
    }

    logout(): void {
        this.store.clear();
    }

    onChange(listener: AuthListener<UserDto>): void {
        if (!this.listeners.includes(listener)) {
            this.listeners.push(listener);
        }
    }

    async user(): Promise<UserDto | undefined | null> {
        return null;
    }
}