import { TOKEN_KEY } from "./AuthService";
import { HEADER_APPLICATION_JSON } from "../consts/rest";

export interface RequestProps extends RequestInit {
    ignoreAuth?: boolean;
}

export class ApiService {
    constructor(public readonly baseUrl: string, public readonly ignoreAuth: boolean = false) {
    }

    get(url: string, init?: RequestProps): Promise<Response> {
        return this.fetch(url, init);
    }

    post(url: string, body: any, init?: RequestProps): Promise<Response> {
        const ri = {
            ...{
                method: "POST",
                headers: HEADER_APPLICATION_JSON,
                body: JSON.stringify(body),
            }, ...init,
        };
        return this.fetch(url, ri);
    }

    put(url: string, body: any, init?: RequestProps): Promise<Response> {
        const ri = {
            ...{
                method: "PUT",
                headers: HEADER_APPLICATION_JSON,
                body: JSON.stringify(body),
            }, ...init,
        };
        return this.fetch(url, ri);
    }

    patch(url: string, body: any, init?: RequestProps): Promise<Response> {
        const ri = {
            ...{
                method: "PATCH",
                headers: HEADER_APPLICATION_JSON,
                body: JSON.stringify(body),
            }, ...init,
        };
        return this.fetch(url, ri);
    }

    delete(url: string, init?: RequestProps): Promise<Response> {
        const ri = {
            ...{
                method: "DELETE",
                headers: HEADER_APPLICATION_JSON,
            }, ...init,
        };
        return this.fetch(url, ri);
    }

    fetch(url: string, init?: RequestProps): Promise<Response> {
        const ignoreAuth = init?.ignoreAuth !== undefined ? init.ignoreAuth === true : this.ignoreAuth;

        if (!ignoreAuth) {
            const token = localStorage.getItem(TOKEN_KEY);
            if (token) {
                if (!init) {
                    init = {
                        method: "GET",
                    };
                }
                init.headers = { ...init.headers, ...{ Authorization: `Bearer ${token}` } };
            }
        }

        const baseUrl = this.baseUrl.replace(/\/\s*$/, "");

        let absolute = new URL(`${baseUrl}${url}`, window.location.href)
        return fetch(absolute.href, init).then(async (res) => {
            if (res.status === 401 && !this.ignoreAuth) {
                localStorage.removeItem(TOKEN_KEY);
            }
            return res;
        });
    }

    isResponseOk(response: Response): boolean {
        return response.status >= 200 && response.status <= 299;
    }

}
