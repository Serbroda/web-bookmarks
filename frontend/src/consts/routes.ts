import Home from "../routes/HomeRoute.svelte";
import About from "../routes/GroupRoute.svelte";

export interface Route {
    route: string;
    name: string;
}

const routesItems: Route[] = [
    {
        route: "/",
        name: "Home",
    },
];

export { routesItems };
