import { ApiService } from "./ApiService";

export class AppService extends ApiService {
    constructor(baseUrl: string) {
        super(baseUrl, true);
    }

    async getVersion(): Promise<string> {
        const response = await this.get(`/version`);
        return response.text();
    }
}
