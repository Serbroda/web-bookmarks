import {FetchParams, Middleware, RequestContext, ResponseContext} from "../../gen";
import {AuthStore} from "../auth/auth.store";
import {authStore} from "../config";

const RETRY_HEADER_KEY = 'Retry';

export type RefreshTokenFn<T> = () => Promise<T>;

export interface TokenMiddlewareProps {
    authStore: AuthStore<any>;
    refreshTokenFn?: RefreshTokenFn<any>;
}

export default class TokenMiddleware implements Middleware {
    constructor(private props: TokenMiddlewareProps) {
    }

    public async pre(context: RequestContext): Promise<void | FetchParams> {
        const accessToken = this.props.authStore.accessToken;

        if (accessToken) {
            return {
                url: context.url,
                init: {
                    ...context.init,
                    headers: new Headers({
                        ...context.init.headers,
                        Authorization: `Bearer ${accessToken}`
                    }),
                },
            };
        } else {
            return {url: context.url, init: context.init};
        }
    }

    public async post(context: ResponseContext): Promise<void | Response> {
        if (context.response.status === 401 || context.response.status === 403) {
            const headers = new Headers(context.init.headers);

            const retry = headers.get(RETRY_HEADER_KEY);
            if (this.props.refreshTokenFn && (retry === undefined || retry === null || retry !== "true")) {
                try {
                    const accessToken = await this.props.refreshTokenFn();
                    if (accessToken) {
                        return context.fetch(context.url, {
                            ...context.init,
                            headers: new Headers({
                                ...context.init.headers,
                                Authorization: `Bearer ${accessToken}`,
                                RETRY_HEADER_KEY: "true"
                            }),
                        })
                    }
                } catch (err) {
                    console.error("Failed to refresh token", err);
                }
                authStore.clear();
            }
        }
        return Promise.resolve(context.response);
    }
}

