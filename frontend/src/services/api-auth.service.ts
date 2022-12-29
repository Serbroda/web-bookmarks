import {AuthMethod, AuthService, isOAuth2Method} from "./auth/auth.service";
import {AuthApi, TokenPairDto, UserDto} from "../gen";
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
        const refreshToken = this.store.refreshToken
        if (refreshToken === undefined || refreshToken === null) {
            throw new Error("Refresh token not set");
        }

        let response = await this.auth.refreshToken({
            inlineObject: {refreshToken}
        });
        this.applyToken(response);
        const u = await this.user();
        this.store.user = u;
        return u;
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
            this.applyToken(response);

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

    private applyToken(tokens: TokenPairDto) {
        if (tokens.accessToken) {
            this.store.accessToken = tokens.accessToken;
        }
        if (tokens.refreshToken) {
            this.store.refreshToken = tokens.refreshToken;
        }
    }
}