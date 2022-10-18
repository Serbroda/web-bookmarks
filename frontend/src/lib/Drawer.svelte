<script lang="ts">
    import active from "svelte-spa-router/active";
    import { replace } from "svelte-spa-router";
    import logo from "../assets/logo.svg";
    import { groups } from "../stores/groups";
    import { Icon, Home } from "svelte-hero-icons";
    import { authService } from "../services/Services";
    import { user } from "../stores/auth";
    import { hashString } from "../utils/string";

    export let onCreateGroupClick: () => void;
    export let version = "";
</script>

<div class="drawer-side h-full">
    <label for="main-menu" class="drawer-overlay" />
    <aside class="flex flex-col bg-base-200 text-base-content w-80 h-full">
        <div
            class="sticky inset-x-0 top-0 z-10 w-full py-1 transition duration-200 ease-in-out border-b border-base-200 bg-base-200">
            <div class="mx-auto space-x-1 navbar max-w-none">
                <div class="flex items-center flex-none">
                    <a href="/" class="px-2 flex-0 btn btn-ghost md:px-4 nuxt-link-active" aria-label="Homepage">
                        <div class="inline-block text-3xl font-title text-primary flex">
                            <img src={logo} class="w-9 h-9 mr-3" alt="Logo" />
                            <span class="lowercase">rag</span><span class="uppercase text-base-content">bag</span>
                        </div>
                    </a>
                </div>
                <dd class="text-left font-mono opacity-50 linte-clamp-1" style="font-size: 0.6rem;">
                    {version}
                </dd>
            </div>
        </div>
        <div class=" h-full">
            <ul class="menu flex flex-col p-4 pt-2 compact">
                <li>
                    <a
                        href="/#/"
                        class="capitalize active:bg-base-200 active:text-primary"
                        use:active={{ path: "/", className: "bg-base-300" }}>
                        <Icon src={Home} class="h-5 w-5" />
                        Home
                    </a>
                </li>
            </ul>

            <ul class="menu flex flex-col p-4 pt-0 compact">
                <li class="menu-title -ml-2">
                    <span>
                        My Groups

                        <div class="tooltip tooltip-left float-right" data-tip="Add Group">
                            <button on:click={onCreateGroupClick} class="hover:text-primary"
                                ><svg
                                    class="w-6 h-6"
                                    fill="none"
                                    stroke="currentColor"
                                    viewBox="0 0 24 24"
                                    xmlns="http://www.w3.org/2000/svg"
                                    ><path
                                        stroke-linecap="round"
                                        stroke-linejoin="round"
                                        stroke-width="2"
                                        d="M12 6v6m0 0v6m0-6h6m-6 0H6" /></svg>
                            </button>
                        </div>
                    </span>
                </li>

                {#each $groups || [] as group}
                    <li>
                        <a
                            class="active:bg-base-200 active:text-primary"
                            href={`/#/groups/${group.id}`}
                            use:active={{ path: `/groups/${group.id}`, className: "bg-base-300" }}>
                            <span style="min-width: 20px">{group.icon}</span>
                            <span>{group.name}</span>
                        </a>
                    </li>
                {/each}
            </ul>
        </div>

        <footer class="sticky inset-x-0 bottom-0 bg-base-200 border-t p-2">
            <div class="dropdown dropdown-top">
                <label tabindex="0" class="btn btn-ghost btn-circle avatar">
                    <div class="w-10 rounded-full">
                        <img
                            src={`https://source.boringavatars.com/beam/120/${hashString(
                                $user?.username.toLowerCase(),
                            )}?colors=264653,f4a261,e76f51`} />
                    </div>
                </label>
                <ul
                    tabindex="0"
                    class="mt-3 p-2 shadow menu menu-compact dropdown-content bg-base-100 rounded-box w-52">
                    <li>
                        <a> Profile </a>
                    </li>
                    <li>
                        <a>Settings</a>
                    </li>
                    <li>
                        <button
                            on:click={async () => {
                                authService.logout();
                                await replace("/login");
                            }}>Logout</button>
                    </li>
                </ul>
            </div>
        </footer>
    </aside>
</div>
