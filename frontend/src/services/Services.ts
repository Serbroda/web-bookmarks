import { AuthService } from "./AuthService";
import { GroupService } from "./GroupService";

const { VITE_BACKEND_BASE_URL } = import.meta.env;

const baseUrl: string = VITE_BACKEND_BASE_URL || "/";

const authService = new AuthService(baseUrl);
const groupService = new GroupService(baseUrl);

export { authService, groupService };
