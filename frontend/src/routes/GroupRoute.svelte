<script lang="ts">
    import { replace, params } from "svelte-spa-router";
    import { toasts } from "svelte-toasts";
    import MainContainer from "../lib/MainContainer.svelte";
    import { groupService } from "../services/Services";
    import { GroupVisibility, type GroupDto } from "../models/GroupDto";
    import GroupModal from "../lib/modals/GroupModal.svelte";
    import CreateLinkModal from "../lib/modals/CreateLinkModal.svelte";
    import LinkCard from "../lib/LinkCard.svelte";
    import type { LinkDto } from "../models/LinkDto";
    import ConfirmModal from "../lib/modals/ConfirmModal.svelte";
    import NavBar from "../lib/NavBar.svelte";
    import { querystring } from "svelte-spa-router";
    import { pushHistoryState, addQuery, removeQuery } from "../utils/url";
    import EditLinkModal from "../lib/modals/EditLinkModal.svelte";
    import ShareModal from "../lib/modals/ShareModal.svelte";
    import { Icon, Share, Eye, EyeOff } from "svelte-hero-icons";

    let id;

    let group: GroupDto | undefined;
    let links: LinkDto[] = [];

    let isCreateLinkModalOpen: boolean = false;
    let isEditLinkModalOpen: boolean = false;
    let isGroupModalOpen: boolean = false;
    let isLinkModalBusy: boolean = false;
    let isGroupShareModalOpen: boolean = false;

    let inputName: string = "";
    let inputUrl: string = "";
    let inputDescription: string = "";
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

    const updateGroup = async (icon: string, name: string, description?: string) => {
        const response = await groupService.updateGroup(id, {
            name,
            icon,
            description,
        });
        isGroupModalOpen = false;
        toasts.success("Group saved");

        group = await groupService.getGroup(id);
    };

    const changeGroupVisibility = async (id: string, visibility: GroupVisibility) => {
        const response = await groupService.changeGroupVisibility(id, visibility);
        group = await groupService.getGroup(id);
        toasts.success("Group saved");

        if (visibility == "public") {
            isGroupShareModalOpen = true;
        }
    };

    const deleteGroup = async () => {
        const response = await groupService.deleteGroup(id);
        await replace("/");
    };

    const deleteLink = async (link: LinkDto) => {
        try {
            isLinkModalBusy = true;

            const response = await groupService.deleteLink(link.id);

            const currentUrl = window.location.href;
            const url = removeQuery(currentUrl, "link");
            if (currentUrl !== url) {
                pushHistoryState(url);
            }

            isEditLinkModalOpen = false;
            toasts.success("Link deleted");

            links = await groupService.getLinks(id);
        } finally {
            isLinkModalBusy = false;
        }
    };
</script>

{#if group}
    <CreateLinkModal
        url={inputUrl}
        name={inputName}
        description={inputDescription}
        isOpen={isCreateLinkModalOpen}
        isBusy={isLinkModalBusy}
        selectedGroup={group}
        onClose={() => (isCreateLinkModalOpen = false)}
        onSave={async () => {
            isCreateLinkModalOpen = false;
            links = await groupService.getLinks(id);
        }} />

    <EditLinkModal
        isOpen={isEditLinkModalOpen}
        isBusy={isLinkModalBusy}
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
        }}
        onSave={async () => {
            isEditLinkModalOpen = false;
            links = await groupService.getLinks(id);
        }}
        onDelete={async () => {
            confirmModal.show({
                title: "Delete Link",
                message: "Do you really want to delete this Link?",
                buttons: [
                    {
                        content: "Delete Link",
                        onClick: async () => {
                            await deleteLink(selectedItem);
                            confirmModal.hide();
                        },
                        classes: "btn-error",
                    },
                    {
                        content: "Cancel",
                        onClick: () => confirmModal.hide(),
                    },
                ],
            });
        }} />

    <GroupModal
        isOpen={isGroupModalOpen}
        onClose={() => (isGroupModalOpen = false)}
        onSave={async (emoji, name, description) => await updateGroup(emoji, name, description)}
        selectedEmoji={group?.icon}
        name={group?.name}
        description={group?.description}
        mode="Edit" />

    <ShareModal
        isOpen={isGroupShareModalOpen}
        title={`Share Group "${group.icon} ${group.name}"`}
        link={`${window.location.origin}${window.location.pathname}#/public/groups/${group.id}`}
        onClose={() => (isGroupShareModalOpen = false)} />

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
                <button class="btn btn-ghost hidden md:flex" on:click={() => (isCreateLinkModalOpen = true)}>
                    <svg
                        class="w-6 h-6 mr-2"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                        xmlns="http://www.w3.org/2000/svg">
                        <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M12 9v3m0 0v3m0-3h3m-3 0H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    Add Link
                </button>

                {#if group.visibility === "public"}
                    <button class="btn btn-ghost hidden md:flex" on:click={() => (isGroupShareModalOpen = true)}>
                        <Icon src={Share} class="h-6 w-6" />
                    </button>
                {/if}

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
                        class="dropdown-content menu p-2 shadow bg-base-100 rounded-box w-52 text-base font-normal">
                        <li class="md:hidden">
                            <button on:click={() => (isCreateLinkModalOpen = true)}>
                                <svg
                                    class="w-6 h-6 mr-2"
                                    fill="none"
                                    stroke="currentColor"
                                    viewBox="0 0 24 24"
                                    xmlns="http://www.w3.org/2000/svg">
                                    <path
                                        stroke-linecap="round"
                                        stroke-linejoin="round"
                                        stroke-width="2"
                                        d="M12 9v3m0 0v3m0-3h3m-3 0H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z" />
                                </svg>
                                Add Link
                            </button>
                        </li>
                        <li>
                            <button on:click={() => (isGroupModalOpen = true)}>
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
                                        d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                                </svg>
                                Edit Group
                            </button>
                        </li>
                        {#if group.visibility === "private"}
                            <li>
                                <button
                                    on:click={async () =>
                                        await changeGroupVisibility(group.id, GroupVisibility.PUBLIC)}>
                                    <Icon src={Eye} class="h-6 w-6" />
                                    Make public
                                </button>
                            </li>
                        {/if}
                        {#if group.visibility === "public"}
                            <li>
                                <button
                                    on:click={async () =>
                                        await changeGroupVisibility(group.id, GroupVisibility.PRIVATE)}>
                                    <Icon src={EyeOff} class="h-6 w-6" />
                                    Make private
                                </button>
                            </li>
                            <li class="md:hidden">
                                <button on:click={() => (isGroupShareModalOpen = true)}>
                                    <Icon src={Share} class="h-6 w-6" />
                                    Share Group
                                </button>
                            </li>
                        {/if}
                        <li>
                            <button
                                on:click={async () => {
                                    confirmModal.show({
                                        title: "Delete Group",
                                        message: "Do you really want to delete this Group?",
                                        buttons: [
                                            {
                                                content: "Delete",
                                                onClick: async () => {
                                                    await deleteGroup();

                                                    confirmModal.hide();
                                                    isEditLinkModalOpen = false;
                                                    toasts.success("Group deleted");
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
                                Delete Group
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
