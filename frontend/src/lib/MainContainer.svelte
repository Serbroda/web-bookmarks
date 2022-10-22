<script lang="ts">
    import { onMount } from "svelte";
    import { authService, groupService } from "../services/Services";
    import { FlatToast, ToastContainer, toasts } from "svelte-toasts";
    import GroupModal from "./modals/GroupModal.svelte";
    import { groups } from "../stores/groups";
    import { user } from "../stores/auth";
    import Drawer from "./Drawer.svelte";
    import { Icon, QuestionMarkCircle } from "svelte-hero-icons";
    import { toggleTheme } from "../App.svelte";

    let isGroupModalOpen: boolean = false;
    let version = "";

    onMount(async () => {
        $user = await authService.getUser();
        version = await groupService.getVersion();
        version = "0.0.1-alpha.2+1666452764";

        await loadGroups();
    });

    const loadGroups = async () => {
        if (!authService.isLoggedIn()) {
            return;
        }
        $groups = await groupService.getGroups();
    };

    const createGroup = async (icon: string, name: string, description?: string) => {
        if (!authService.isLoggedIn()) {
            return;
        }

        try {
            const response = await groupService.createGroup({
                icon,
                name,
                description,
            });
            await loadGroups();
        } catch (err) {
            console.error(err);
        }
    };
</script>

<GroupModal
    isOpen={isGroupModalOpen}
    onClose={() => (isGroupModalOpen = false)}
    mode="Add"
    onSave={async (icon, name, description) => {
        await createGroup(icon, name, description);
        toasts.success("Group saved");
        isGroupModalOpen = false;
    }} />

<main>
    <div class="drawer drawer-mobile">
        <input id="main-menu" type="checkbox" class="drawer-toggle" />
        <main class="flex-grow block overflow-x-hidden bg-base-100 text-base-content drawer-content">
            <div id="main-content">
                <slot />

                <div class="dropdown dropdown-top dropdown-end absolute bottom-8 right-8">
                    <!-- svelte-ignore a11y-label-has-associated-control -->
                    <label tabindex="0" class="btn btn-ghost btn-sm btn-circle">
                        <Icon src={QuestionMarkCircle} class="h-5 w-5" />
                    </label>

                    <div tabindex="0" class="card compact dropdown-content shadow bg-base-100 rounded-box w-64">
                        <div class="card-body">
                            <ul tabindex="0" class="menu p-2 px-0 rounded-box">
                                <li>
                                    <button on:click={toggleTheme}
                                        ><svg
                                            width="20"
                                            height="20"
                                            xmlns="http://www.w3.org/2000/svg"
                                            fill="none"
                                            viewBox="0 0 24 24"
                                            class="inline-block h-5 w-5 stroke-current md:h-6 md:w-6"
                                            ><path
                                                stroke-linecap="round"
                                                stroke-linejoin="round"
                                                stroke-width="2"
                                                d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zm0 0h12a2 2 0 002-2v-4a2 2 0 00-2-2h-2.343M11 7.343l1.657-1.657a2 2 0 012.828 0l2.829 2.829a2 2 0 010 2.828l-8.486 8.485M7 17h.01" /></svg>
                                        Change theme</button>
                                </li>
                            </ul>
                            <span class="text-center font-mono line-clamp-1 text-xs">
                                {version}
                            </span>
                        </div>
                    </div>
                </div>
            </div>
        </main>

        <Drawer onCreateGroupClick={() => (isGroupModalOpen = true)} />
    </div>

    <ToastContainer placement="bottom-right" let:data>
        <FlatToast {data} />
    </ToastContainer>

    {#each [0, 10, 20, 30, 40, 50] as zIndex}
        <div id={`__modal-container-${zIndex}`} class={`z-${zIndex}`} />
    {/each}
</main>
