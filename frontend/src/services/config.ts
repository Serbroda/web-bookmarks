import {AuthApi, Configuration, SpacesApi, UserDto} from "../gen";
import TokenMiddleware from "./middlewares/token.middleware";
import {AuthService} from "./auth/auth.service";
import {ApiAuthService} from "./api-auth.service";
import {AuthStore} from "./auth/auth.store";

const { VITE_BACKEND_BASE_URL } = import.meta.env;

const basePath: string = VITE_BACKEND_BASE_URL || "/";

const authStore = new AuthStore<UserDto>();

const config = new Configuration({
    basePath,
    middleware: [new TokenMiddleware(authStore)]
});

const authApi = new AuthApi(config);
const spacesApi = new SpacesApi(config);

const authService: AuthService<UserDto> = new ApiAuthService(authApi, authStore);

export { authStore, authApi, spacesApi, authService };
