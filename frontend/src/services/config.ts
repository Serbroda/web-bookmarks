import {AuthApi, Configuration, SpacesApi} from "../gen";
import TokenMiddleware from "./middlewares/token.middleware";
import {AuthService} from "./auth.service";
import {ApiAuthService} from "./api-auth.service";

const { VITE_BACKEND_BASE_URL } = import.meta.env;

const basePath: string = VITE_BACKEND_BASE_URL || "/";

const config = new Configuration({
    basePath,
    middleware: [new TokenMiddleware()]
});

const authApi = new AuthApi(config);
const spacesApi = new SpacesApi(config);

const authService: AuthService<any> = new ApiAuthService(authApi);

export { authApi, spacesApi, authService };