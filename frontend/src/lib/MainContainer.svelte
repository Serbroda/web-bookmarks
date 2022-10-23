<script lang="ts">
    import { onMount } from "svelte";
    import { authService, groupService } from "../services/Services";
    import { FlatToast, ToastContainer, toasts } from "svelte-toasts";
    import GroupModal from "./modals/GroupModal.svelte";
    import CreateGroupSubscriptionModal from "./modals/CreateGroupSubscriptionModal.svelte";
    import { groups, groupSubscriptions } from "../stores/groups";
    import Drawer from "./Drawer.svelte";
    import InfoDropdown from "./components/InfoDropdown.svelte";

    let isGroupModalOpen: boolean = false;
    let isCreateGroupSubscriptionModalOpen: boolean = false;

    onMount(async () => {
        await loadGroups();
        await loadGroupSubscriptions();
    });

    const loadGroups = async () => {
        if (!authService.isLoggedIn()) {
            return;
        }
        $groups = await groupService.getGroups();
    };

    const loadGroupSubscriptions = async () => {
        if (!authService.isLoggedIn()) {
            return;
        }
        $groupSubscriptions = await groupService.getGroupSubscriptions();
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

<CreateGroupSubscriptionModal
    isOpen={isCreateGroupSubscriptionModalOpen}
    onClose={() => (isCreateGroupSubscriptionModalOpen = false)}
    onSuccess={async (subscription) => {
        await loadGroupSubscriptions();
        toasts.success("Group subscribed");
        isCreateGroupSubscriptionModalOpen = false;
    }} />

<main>
    <div class="drawer drawer-mobile">
        <input id="main-menu" type="checkbox" class="drawer-toggle" />
        <main class="flex-grow block overflow-x-hidden bg-base-100 text-base-content drawer-content">
            <div id="main-content">
                <slot />

                <InfoDropdown classes="absolute bottom-8 right-8 hidden dropdown-top dropdown-end lg:block" />
            </div>
        </main>

        <Drawer
            onCreateGroupClick={() => (isGroupModalOpen = true)}
            onCreateGroupSubscriptionClick={() => (isCreateGroupSubscriptionModalOpen = true)} />
    </div>

    <ToastContainer placement="bottom-right" let:data>
        <FlatToast {data} />
    </ToastContainer>

    {#each [0, 10, 20, 30, 40, 50] as zIndex}
        <div id={`__modal-container-${zIndex}`} class={`z-${zIndex}`} />
    {/each}
</main>
