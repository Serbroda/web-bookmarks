export type AccessToken = string;
export type RefreshToken = string;

export type AuthListener<T> = ((authenticated: boolean, user: T | null | undefined) => void);

export class AuthStore<T> {
    _listeners: AuthListener<T>[] = [];
    _timeout: number | undefined;

    constructor(
        private accessTokenKey: string = "access_token",
        private refreshTokenKey: string = "refresh_token") {
    }

    public set accessToken(token: AccessToken | null | undefined) {
        this.setItem(this.accessTokenKey, token);
    }

    public set refreshToken(token: RefreshToken | null | undefined) {
        this.setItem(this.refreshTokenKey, token);
    }

    public get accessToken(): AccessToken | null | undefined {
        return localStorage.getItem(this.accessTokenKey);
    }

    public get refreshToken(): AccessToken | null | undefined {
        return localStorage.getItem(this.refreshTokenKey);
    }

    public set user(user: T | null | undefined) {
        if (user) {
            this.setItem("_user", JSON.stringify(user))
        } else {
            this.setItem("_user", undefined)
        }
    }

    public get user(): T | null | undefined {
        const userRaw = localStorage.getItem("_user");
        if (userRaw) {
            return JSON.parse(userRaw);
        } else {
            return undefined;
        }
    }

    public get authenticated(): boolean {
        const t = this.accessToken;
        return t !== undefined && t !== null;
    }

    public clear() {
        this.accessToken = null;
        this.refreshToken = null;
        this.user = null;
    }

    public onChange(listener: (authenticated: boolean, user: T | undefined | null) => void) {
        if (!this._listeners.includes(listener)) {
            this._listeners.push(listener);
        }
    }

    private setItem(key: string, value: string | null | undefined): void {
        const currentValue = localStorage.getItem(key);
        if (value) {
            localStorage.setItem(key, value);
        } else {
            localStorage.removeItem(key);
        }

        if (currentValue !== value) {
            this.informListeners();
        }
    }

    private informListeners() {
        clearTimeout(this._timeout);
        this._timeout = setTimeout(() => {
            for (const l of this._listeners) {
                l(this.authenticated, this.user);
            }
        }, 150);
    }
}
