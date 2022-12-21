import PocketBase from 'pocketbase'
import PocketbaseAuthService from "./pocketbase-auth.service";
import PocketbaseApiService from "./pocketbase-api.service";

const pb = new PocketBase('http://127.0.0.1:8090');

const authService = new PocketbaseAuthService(pb);
const apiService = new PocketbaseApiService(pb);

export default pb;
export { authService, apiService }
