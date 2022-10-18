<script lang="ts">
    import active from "svelte-spa-router/active";
    import { routesItems } from "../consts/routes.js";
    import { onMount } from "svelte";
    import { authService, groupService } from "../services/Services";
    import { FlatToast, ToastContainer, toasts } from "svelte-toasts";
    import GroupModal from "./modals/GroupModal.svelte";
    import { groups } from "../stores/groups";
    import { user } from "../stores/auth";
    import logo from "../assets/logo.svg";
    import Drawer from "./Drawer.svelte";
    import NavBar from "./NavBar.svelte";

    let isGroupModalOpen: boolean = false;
    let version = "";

    onMount(async () => {
        $user = await authService.getUser();
        version = await groupService.getVersion();

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
            </div>
        </main>

        <Drawer {version} onCreateGroupClick={() => (isGroupModalOpen = true)} />
    </div>

    <ToastContainer placement="bottom-right" let:data>
        <FlatToast {data} />
    </ToastContainer>

    {#each [0, 10, 20, 30, 40, 50] as zIndex}
        <div id={`__modal-container-${zIndex}`} class={`z-${zIndex}`} />
    {/each}
</main>
