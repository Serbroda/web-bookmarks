
export type Token = string;

export interface LoginRequest {
    username: string;
    password: string;
}

export interface AuthResponse {
    accessToken: Token;
    refreshToken: Token;
}

export interface IAuthService {
    login(login: LoginRequest): Promise<AuthResponse>;
    logout(): Promise<void>;
    isAuthenticated(): Promise<boolean>;
}

export class AuthService implements IAuthService {

    constructor(private baseUrl: string) {
    }
    
    async login(login: LoginRequest): Promise<AuthResponse> {
        const response = await fetch(`${this.baseUrl}/auth`, {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(login)
          })
        return response.json();
    }
    
    async logout(): Promise<void> {
        throw new Error("Method not implemented.");
    }

    async isAuthenticated(): Promise<boolean> {
        throw new Error("Method not implemented.");
    }
}