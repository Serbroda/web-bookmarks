import {AuthApi, Configuration, SpacesApi, ModelsUser} from "../gen";
import TokenMiddleware from "./middlewares/token.middleware";
import {AuthService} from "./auth/auth.service";
import {ApiAuthService} from "./api-auth.service";
import {AuthStore} from "./auth/auth.store";

const {VITE_BACKEND_BASE_URL} = import.meta.env;

const basePath: string = VITE_BACKEND_BASE_URL || "/";

const authStore = new AuthStore<ModelsUser>();

const publicConfig = new Configuration({
    basePath,
});

const authApi = new AuthApi(publicConfig);
const authService: AuthService<ModelsUser> = new ApiAuthService(authApi, authStore);

const restrictedConfig = new Configuration({
    basePath,
    middleware: [new TokenMiddleware({
        authStore,
        refreshTokenFn: async () => {
            await authService.authRefresh();
            return authStore.accessToken;
        }
    })]
});

const spacesApi = new SpacesApi(restrictedConfig);

export {authStore, authApi, spacesApi, authService};
