import { AuthApi, Configuration, ConfigurationParameters, GroupsApi, LinksApi, SpacesApi } from "../gen";
import TokenMiddleware from "./token.middleware";

const { VITE_BACKEND_BASE_URL } = import.meta.env;

const configParams: ConfigurationParameters = {
    basePath: `${VITE_BACKEND_BASE_URL}/api/v1`,
    middleware: [new TokenMiddleware()],
};

const apiConfig = new Configuration(configParams);

const authApi = new AuthApi(apiConfig);
const spacesApi = new SpacesApi(apiConfig);
const groupsApi = new GroupsApi(apiConfig);
const linksApi = new LinksApi(apiConfig);

export { authApi, spacesApi, groupsApi, linksApi };
