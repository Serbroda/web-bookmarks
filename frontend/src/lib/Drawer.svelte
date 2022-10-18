<script lang="ts">
    import active from "svelte-spa-router/active";
    import logo from "../assets/logo.svg";
    import { routesItems } from "../consts/routes.js";
    import { groups } from "../stores/groups";

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
                {#each routesItems as r}
                    <li>
                        <a
                            href={"/#" + r.route}
                            class="capitalize active:bg-base-200 active:text-primary"
                            use:active={{ path: r.route, className: "bg-base-300" }}>
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                fill="none"
                                viewBox="0 0 24 24"
                                class="inline-block w-6 h-6 mr-2 stroke-current">
                                <path
                                    stroke-linecap="round"
                                    stroke-linejoin="round"
                                    stroke-width="2"
                                    d="M8 4H6a2 2 0 00-2 2v12a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-2m-4-1v8m0 0l3-3m-3 3L9 8m-5 5h2.586a1 1 0 01.707.293l2.414 2.414a1 1 0 00.707.293h3.172a1 1 0 00.707-.293l2.414-2.414a1 1 0 01.707-.293H20" />
                            </svg>
                            {r.name}
                        </a>
                    </li>
                {/each}
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

        <footer class="sticky inset-x-0 bottom-0 bg-base-200 border-t p-2" />
    </aside>
</div>
