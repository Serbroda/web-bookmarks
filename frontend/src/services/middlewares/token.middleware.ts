import { FetchParams, Middleware, RequestContext, ResponseContext } from "../../gen";

export type AccessToken = string;
export type RefreshToken = string;

const ACCESS_TOKEN = "access_token";
const REFRESH_TOKEN = "refresh_token";

const getAccessToken = (): AccessToken | null => {
    return localStorage.getItem(ACCESS_TOKEN);
}

const getRefreshToken = (): RefreshToken | null => {
    return localStorage.getItem(REFRESH_TOKEN);
}

const setAccessToken = (token: AccessToken | null) => {
    if(token) {
        localStorage.setItem(ACCESS_TOKEN, token);
    } else {
        localStorage.removeItem(ACCESS_TOKEN);
    }
}

const setRefreshToken = (token: RefreshToken | null) => {
    if(token) {
        localStorage.setItem(REFRESH_TOKEN, token);
    } else {
        localStorage.removeItem(REFRESH_TOKEN);
    }
}

export default class TokenMiddleware implements Middleware {

    public async pre(context: RequestContext): Promise<void | FetchParams> {
        const accessToken = getAccessToken();

        if(accessToken) {
            return {
                url: context.url,
                init: {
                    ...context.init,
                    headers: new Headers({
                        ...context.init.headers,
                        Authorization: `Bearer ${accessToken}`,
                    }),
                },
            };
        } else {
            return {url: context.url, init: context.init};
        }
    }

    public post(context: ResponseContext): Promise<void | Response> {
        return Promise.resolve(context.response);
    }
}

export {getAccessToken, setAccessToken, ACCESS_TOKEN, REFRESH_TOKEN}