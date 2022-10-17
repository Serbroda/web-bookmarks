import { ApiService } from "./ApiService";
import { HEADER_APPLICATION_JSON, HEADER_APPLICATION_X_WWW_FORM_URLENCODED } from "../consts/rest";

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

        if(typeof formDataOrUsername === "string") {
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

        const response = await this.post(`/login`, data, {headers, body: data});

        if(!this.isResponseOk(response)) {
            throw new Error(`Login failed ${response.statusText}`)
        }

        const token = await response.text();
        localStorage.setItem(TOKEN_KEY, token);
        return response;
    }

    async register(formData: FormData): Promise<Response>;
    async register(username: string, password: string, email: string): Promise<Response>;
    async register(formDataOrUsername: FormData | string, password?: string, email?: string): Promise<Response> {
        let data;
        let headers;

        if(typeof formDataOrUsername === "string") {
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

        return this.post(`/register`, data, {headers, body: data});
    }

    logout() {
        localStorage.removeItem(TOKEN_KEY);
    }

    isLoggedIn(): boolean {
        const token = localStorage.getItem(TOKEN_KEY);
        return token !== undefined && token !== null && token.trim() !== "";
    }
}

export { TOKEN_KEY };
