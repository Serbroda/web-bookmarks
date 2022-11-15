import { FetchParams, Middleware, RequestContext, ResponseContext } from "../gen";

const ACCESS_TOKEN = "access_token";

const getToken = (): string | null => {
    return localStorage.getItem(ACCESS_TOKEN);
}

const setToken = (token: string | null) => {
    if(token) {
        localStorage.setItem(ACCESS_TOKEN, token);
    } else {
        localStorage.removeItem(ACCESS_TOKEN);
    }
}

export default class TokenMiddleware implements Middleware {

    public async pre(context: RequestContext): Promise<void | FetchParams> {
        const accessToken = getToken();

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

export {getToken, setToken}