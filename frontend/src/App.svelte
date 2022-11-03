<script lang="ts" context="module">
    let activeTheme = localStorage.getItem("theme");

    export const setTheme = (theme: string) => {
        const htmlElement = document.querySelector("html");
        htmlElement.dataset.theme = theme;

        activeTheme = theme;
        localStorage.setItem("theme", theme);
    };

    export const toggleTheme = () => {
        setTheme(activeTheme === "ragbaglight" ? "ragbagdark" : "ragbaglight");
    };
</script>

<script lang="ts">
    import { authService } from "./services/Services.js";
    import wrap from "svelte-spa-router/wrap";
    import Router, { replace } from "svelte-spa-router";
    import GroupRoute from "./routes/GroupRoute.svelte";
    import HomeRoute from "./routes/HomeRoute.svelte";
    import LoginRoute from "./routes/LoginRoute.svelte";
    import RegisterRoute from "./routes/RegisterRoute.svelte";
    import ProfileRoute from "./routes/ProfileRoute.svelte";
    import { authenticated } from "./stores/auth";
    import ExternalGroupRoute from "./routes/ExternalGroupRoute.svelte";
    import PublicGroupRoute from "./routes/PublicGroupRoute.svelte";

    setTheme(activeTheme || "light");

    const loginCondition = async (details: any): Promise<boolean> => {
        return authService.isLoggedIn();
    };

    const conditionsFailed = async (event) => {
        console.log("Fuck");
        await replace("/login");
    };

    /*authenticated.subscribe(async (val) => {
        if (!val) {
            await replace("/login");
        }
    });*/

    function routeLoaded(event) {
        const mainMenuElement: HTMLElement = document.getElementById("main-menu");
        if (mainMenuElement) {
            (<HTMLInputElement>mainMenuElement).checked = false;
        }
    }

    const wrapAuthenticated = (component: any): any => {
        return wrap({
            component: component,
            conditions: [loginCondition],
        });
    };

    const routes = {
        "/login": LoginRoute,
        "/register": RegisterRoute,
        "/": wrapAuthenticated(HomeRoute),
        "/groups/:groupId": wrapAuthenticated(GroupRoute),
        "/public/groups/:groupId": PublicGroupRoute,
        "/external/groups/:groupId": wrapAuthenticated(ExternalGroupRoute),
        "/profile": wrapAuthenticated(ProfileRoute),
    };
</script>

<Router {routes} on:conditionsFailed={conditionsFailed} on:routeLoaded={routeLoaded} />
