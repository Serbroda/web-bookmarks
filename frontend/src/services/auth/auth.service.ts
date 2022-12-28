
export type AuthPasswordMethod = {
    username: string;
    password: string;
}

export type AuthOAuth2Method = {
    provider: string;
    code: string;
    codeVerifier: string;
    redirectUrl: string;
    createData?: any;
}

export type AuthMethod = AuthPasswordMethod | AuthOAuth2Method;

export const isOAuth2Method = (method: AuthMethod): method is AuthOAuth2Method => {
    return (method as AuthOAuth2Method).provider !== undefined;
}

export interface AuthService<T> {
    authenticate(auth: AuthMethod): Promise<T | undefined | null>;
    isAuthenticated(): boolean;
    authRefresh(): Promise<T | undefined | null>;
    user(): Promise<T | undefined | null>;
    logout(): void;
    onChange(listener: (authenticated: boolean, user: T | undefined | null) => void): void;
}
