export type AccessToken = string;
export type RefreshToken = string;

export type AuthListener<T> = ((authenticated: boolean, user: T | null | undefined) => void);

export interface AuthStoreProps {
    accessTokenKey: string;
    refreshTokenKey: string;
    userKey: string;
    listenStorage: boolean;
    informListenersTimeout: number;
    authenticatedValidator: (accessToken: AccessToken | null | undefined, user?: any) => boolean;
}

export class AuthStore<T> {
    _listeners: AuthListener<T>[] = [];
    _timeout: number | undefined;
    _props: AuthStoreProps;

    constructor(
        props?: Partial<AuthStoreProps>) {
        this._props = {
            ...{
                accessTokenKey: "access_token",
                refreshTokenKey: "refresh_token",
                userKey: "user",
                listenStorage: true,
                informListenersTimeout: 150,
                authenticatedValidator: (token, user) => token !== undefined && token !== null
            }, ...props
        };
        if (this._props.listenStorage) {
            const keys = [this._props.accessTokenKey, this._props.refreshTokenKey, this._props.userKey];
            addEventListener('storage', (e) => {
                if (e.key && keys.includes(e.key)) {
                    this.informListeners();
                }
            })
        }
    }

    public set accessToken(token: AccessToken | null | undefined) {
        this.setItem(this._props.accessTokenKey, token);
    }

    public set refreshToken(token: RefreshToken | null | undefined) {
        this.setItem(this._props.refreshTokenKey, token);
    }

    public get accessToken(): AccessToken | null | undefined {
        return localStorage.getItem(this._props.accessTokenKey);
    }

    public get refreshToken(): AccessToken | null | undefined {
        return localStorage.getItem(this._props.refreshTokenKey);
    }

    public set user(user: T | null | undefined) {
        if (user) {
            this.setItem(this._props.userKey, JSON.stringify(user))
        } else {
            this.setItem(this._props.userKey, undefined)
        }
    }

    public get user(): T | null | undefined {
        const userRaw = localStorage.getItem(this._props.userKey);
        if (userRaw) {
            return JSON.parse(userRaw);
        } else {
            return undefined;
        }
    }

    public get authenticated(): boolean {
        return this._props.authenticatedValidator(this.accessToken, this.user);
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
        }, this._props.informListenersTimeout);
    }
}
