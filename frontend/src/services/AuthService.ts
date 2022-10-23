import { ApiService } from "./ApiService";
import { HEADER_APPLICATION_JSON, HEADER_APPLICATION_X_WWW_FORM_URLENCODED } from "../consts/rest";
import type { UserDto } from "../models/UserDto";
import {token as tokenStore, authenticated as authenticatedStore} from "../stores/auth"

const TOKEN_KEY = "access_token";

export class AuthService extends ApiService {
    constructor(baseUrl: string) {
        super(baseUrl, true);
    }

    async login(formData: FormData): Promise<Response>;
    async login(username: string, password: string): Promise<Response>;
    async login(formDataOrUsername: FormData | string, password?: string): Promise<Response> {
        let data;
        let headers;

        if (typeof formDataOrUsername === "string") {
            data = {
                username: formDataOrUsername,
                password,
            };
            headers = HEADER_APPLICATION_JSON;
        } else {
            // @ts-ignore
            data = new URLSearchParams(formDataOrUsername);
            headers = HEADER_APPLICATION_X_WWW_FORM_URLENCODED;
        }

        const response = await this.post(`/login`, data, { headers, body: data });

        if (!this.isResponseOk(response)) {
            throw new Error(`Login failed ${response.statusText}`);
        }

        const token = await response.text();
        this.setToken(token);
        return response;
    }

    async register(formData: FormData): Promise<Response>;
    async register(username: string, password: string, email: string): Promise<Response>;
    async register(formDataOrUsername: FormData | string, password?: string, email?: string): Promise<Response> {
        let data;
        let headers;

        if (typeof formDataOrUsername === "string") {
            data = {
                username: formDataOrUsername,
                email,
                password,
            };
            headers = HEADER_APPLICATION_JSON;
        } else {
            // @ts-ignore
            data = new URLSearchParams(formDataOrUsername);
            headers = HEADER_APPLICATION_X_WWW_FORM_URLENCODED;
        }

        return this.post(`/register`, data, { headers, body: data });
    }

    async changePassword(formData: FormData): Promise<Response>;
    async changePassword(oldPassword: string, newPassword: string): Promise<Response>;
    async changePassword(formDataOrOldPassword: FormData | string, newPassword?: string): Promise<Response> {
        let data;
        let headers;

        if (typeof formDataOrOldPassword === "string") {
            data = {
                oldPassword: formDataOrOldPassword,
                newPassword,
            };
            headers = HEADER_APPLICATION_JSON;
        } else {
            // @ts-ignore
            data = new URLSearchParams(formDataOrOldPassword);
            headers = HEADER_APPLICATION_X_WWW_FORM_URLENCODED;
        }

        return this.patch(`/api/v1/users/change_password`, data, { ignoreAuth: false, headers, body: data });
    }

    async getUser(): Promise<UserDto> {
        const response = await this.get(`/api/v1/users/me`, { ignoreAuth: false });

        if (response.status === 404) {
            return undefined;
        }
        return response.json();
    }

    logout() {
        localStorage.removeItem(TOKEN_KEY);
        tokenStore.set(undefined);
        authenticatedStore.set(this.isLoggedIn())
    }

    getToken(): string | undefined {
        return localStorage.getItem(TOKEN_KEY);
    }

    setToken(token: string) {
        localStorage.setItem(TOKEN_KEY, token);
        tokenStore.set(token);
        authenticatedStore.set(this.isLoggedIn())
    }

    isLoggedIn(): boolean {
        const token = this.getToken();
        return token !== undefined && token !== null && token.trim() !== "";
    }
}

export { TOKEN_KEY };
