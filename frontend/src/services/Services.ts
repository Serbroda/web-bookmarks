import { AuthService } from "./AuthService";
import { GroupService } from "./GroupService";
import { AppService } from "./AppService";

const { VITE_BACKEND_BASE_URL } = import.meta.env;

const baseUrl: string = VITE_BACKEND_BASE_URL || "/";

const appService = new AppService(baseUrl);
const authService = new AuthService(baseUrl);
const groupService = new GroupService(baseUrl);

export { appService, authService, groupService };
