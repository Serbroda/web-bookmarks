<script lang="ts">
    import { replace, params } from "svelte-spa-router";
    import { toasts } from "svelte-toasts";
    import MainContainer from "../lib/MainContainer.svelte";
    import { groupService } from "../services/Services";
    import LinkCard from "../lib/LinkCard.svelte";
    import type { LinkDto } from "../models/LinkDto";
    import ConfirmModal from "../lib/modals/ConfirmModal.svelte";
    import NavBar from "../lib/NavBar.svelte";
    import { querystring } from "svelte-spa-router";
    import { pushHistoryState, addQuery, removeQuery } from "../utils/url";
    import ViewLinkModal from "../lib/modals/ViewLinkModal.svelte";
    import type { GroupDto } from "../models/GroupDto";

    let id;

    let group: GroupDto | undefined;
    let links: LinkDto[] = [];

    let isEditLinkModalOpen: boolean = false;

    let selectedItem: LinkDto | undefined;
    let confirmModal: ConfirmModal;

    $: if ($params?.groupId) {
        id = $params.groupId;

        Promise.all([groupService.getGroup(id), groupService.getLinks(id)]).then((data) => {
            group = data[0];
            links = data[1];
        });
    }

    $: if (links && $querystring) {
        let params = new URLSearchParams($querystring);
        const linkParam = params.get("link");
        if (linkParam) {
            const res = links.find((l) => l.id === linkParam);
            if (res) {
                selectedItem = res;
                setTimeout(() => {
                    isEditLinkModalOpen = true;
                }, 200);
            }
        }
    }

    const deleteGroupSubscription = async () => {
        const response = await groupService.deleteGroupSubscription(id);
        isEditLinkModalOpen = false;
        await replace("/");
    };
</script>

{#if group}
    <ViewLinkModal
        isOpen={isEditLinkModalOpen}
        link={selectedItem}
        onShow={() => {
            if (selectedItem) {
                const currentUrl = window.location.href;
                const url = addQuery(currentUrl, "link", selectedItem.id);
                if (currentUrl !== url) {
                    pushHistoryState(url);
                }
            }
        }}
        onClose={() => {
            isEditLinkModalOpen = false;

            const currentUrl = window.location.href;
            const url = removeQuery(currentUrl, "link");
            if (currentUrl !== url) {
                pushHistoryState(url);
            }
        }} />

    <ConfirmModal bind:this={confirmModal} />

    <MainContainer>
        <!-- svelte-ignore a11y-label-has-associated-control -->
        <NavBar>
            <svelte:fragment slot="navbar-start">
                <div class="flex flex-col ml-0 lg:ml-3">
                    <h2 class="text-xl lg:text-2xl font-bold flex">
                        <div class="w-[30px] lg:w-[40px]">{group?.icon}</div>
                        <span class="line-clamp-1">{group?.name}</span>
                    </h2>
                    <i class="text-xs leading-4 opacity-80 hidden lg:block ml-1">{group?.description}</i>
                </div>
            </svelte:fragment>

            <svelte:fragment slot="navbar-end">
                <div class="dropdown dropdown-end">
                    <label tabindex="0" class="btn btn-ghost btn-circle">
                        <svg
                            class="w-6 h-6"
                            fill="none"
                            stroke="currentColor"
                            viewBox="0 0 24 24"
                            xmlns="http://www.w3.org/2000/svg">
                            <path
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="2"
                                d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
                        </svg>
                    </label>
                    <ul
                        tabindex="0"
                        class="dropdown-content menu p-2 shadow bg-base-100 rounded-box w-64 text-base font-normal">
                        <li>
                            <button
                                on:click={async () => {
                                    confirmModal.show({
                                        title: "Delete Group Subscription",
                                        message: "Do you really want to delete this Group Subscription?",
                                        buttons: [
                                            {
                                                content: "Delete Group",
                                                onClick: async () => {
                                                    confirmModal.hide();

                                                    await deleteGroupSubscription();

                                                    toasts.success("Group Subscription deleted");
                                                },
                                                classes: "btn-error",
                                            },
                                            {
                                                content: "Cancel",
                                                onClick: () => confirmModal.hide(),
                                            },
                                        ],
                                    });
                                }}>
                                <svg
                                    class="w-6 h-6"
                                    fill="none"
                                    stroke="currentColor"
                                    viewBox="0 0 24 24"
                                    xmlns="http://www.w3.org/2000/svg">
                                    <path
                                        stroke-linecap="round"
                                        stroke-linejoin="round"
                                        stroke-width="2"
                                        d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                                </svg>
                                Delete Subscription
                            </button>
                        </li>
                    </ul>
                </div>
            </svelte:fragment>
        </NavBar>

        <div id="content" class="p-4">
            {#if links}
                <div class="flex flex-wrap">
                    {#each links || [] as i}
                        <LinkCard
                            item={i}
                            onClick={() => {
                                selectedItem = i;
                                isEditLinkModalOpen = true;
                            }} />
                    {/each}
                </div>
            {/if}
        </div>
    </MainContainer>
{/if}
