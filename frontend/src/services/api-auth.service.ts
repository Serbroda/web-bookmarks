import {AuthMethod, AuthService, isOAuth2Method} from "./auth.service";
import {AuthApi, UserDto} from "../gen";
import {getToken, REFRESH_TOKEN, setToken} from "./middlewares/token.middleware";

export class ApiAuthService implements AuthService<UserDto> {
    constructor(private auth: AuthApi) {
    }

    async authRefresh(): Promise<UserDto | undefined | null> {
        return Promise.resolve(undefined);
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
            if(response.accessToken) {
                setToken(response.accessToken!);
            }
            if (response.refreshToken) {
                localStorage.setItem(REFRESH_TOKEN, response.refreshToken);
            }
            return this.user();
        }
    }

    isAuthenticated(): boolean {
        const token = getToken();
        return token !== undefined && token !== null;
    }

    logout(): void {
        setToken(null);
        localStorage.removeItem(REFRESH_TOKEN);
    }

    onChange(listener: (authenticated: boolean, user: any) => void): void {
    }

    async user(): Promise<UserDto | undefined | null> {
        return Promise.resolve(undefined);
    }

}