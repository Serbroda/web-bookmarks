import create from "zustand";
import {Admin, Record} from "pocketbase";
import {AuthMethod} from "../services/auth.service";
import {authService} from "../services/config";

export type AuthenticationState = {
    user: Admin | Record | undefined | null;
    authenticated: boolean;
    init: () => void;
    authenticate: (auth: AuthMethod) => void;
    logout: () => void;
}

const useAuthentication = create<AuthenticationState>((set) => ({
    user: null,
    authenticated: false,
    init: () => {
        const authenticated = authService.isAuthenticated();
        if (!authenticated) {
            set({user: undefined, authenticated})
        } else {
            authService.user()
                .then((res) => set({user: res, authenticated}))
                .catch((err) => set({user: undefined, authenticated}));
        }
        authService.onChange((authenticated, user) => {
            set({user, authenticated})
        })
    },
    authenticate: (auth: AuthMethod) => {
        authService.authenticate(auth)
            .then((res) => set({user: res?.record, authenticated: res !== undefined && res !== null}))
            .catch((err) => set({user: null, authenticated: false}));

    },
    authRefresh: () => {
        authService.authRefresh()
            .then((res) => set({user: res?.record, authenticated: res !== undefined && res !== null}))
            .catch((err) => set({user: null, authenticated: false}));
    },
    logout: () => {
        authService.logout();
        set({user: null, authenticated: false})
    }
}))

export default useAuthentication;
