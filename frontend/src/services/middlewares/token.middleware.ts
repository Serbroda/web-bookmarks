import { FetchParams, Middleware, RequestContext, ResponseContext } from "../../gen";
import {AuthStore} from "../auth/auth.store";

export default class TokenMiddleware implements Middleware {

    constructor(private authStore: AuthStore<any>) {
    }

    public async pre(context: RequestContext): Promise<void | FetchParams> {
        const accessToken = this.authStore.accessToken;

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

