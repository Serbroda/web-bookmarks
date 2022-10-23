<script lang="ts">
    import { params } from "svelte-spa-router";
    import { groupService } from "../services/Services";
    import LinkCard from "../lib/LinkCard.svelte";
    import logo from "../assets/logo.svg";
    import type { LinkDto } from "../models/LinkDto";
    import NavBar from "../lib/NavBar.svelte";
    import { querystring } from "svelte-spa-router";
    import { pushHistoryState, addQuery, removeQuery } from "../utils/url";
    import ViewLinkModal from "../lib/modals/ViewLinkModal.svelte";
    import type { GroupDto } from "../models/GroupDto";
    import { authenticated } from "../stores/auth";
    import CreateGroupSubscriptionModal from "../lib/modals/CreateGroupSubscriptionModal.svelte";
    import { FlatToast, ToastContainer, toasts } from "svelte-toasts";
    import { Icon, PlusCircle } from "svelte-hero-icons";
    import { classNames } from "../utils/dom";

    let id;

    let group: GroupDto | undefined;
    let links: LinkDto[] = [];

    let isEditLinkModalOpen: boolean = false;
    let isCreateGroupSubscriptionModalOpen: boolean = false;

    let selectedItem: LinkDto | undefined;

    $: if ($params?.groupId) {
        id = $params.groupId;

        Promise.all([groupService.getGroupPublic(id), groupService.getLinksPublic(id)]).then((data) => {
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
</script>

{#if group}
    <CreateGroupSubscriptionModal
        isOpen={isCreateGroupSubscriptionModalOpen}
        onClose={() => (isCreateGroupSubscriptionModalOpen = false)}
        value={group.id}
        readonly
        onSuccess={async (subscription) => {
            toasts.success("Group subscribed");
            isCreateGroupSubscriptionModalOpen = false;
        }}
        onError={(error) => {
            console.log(error);
            if (error.status === 409) {
                toasts.warning("Already subscribed");
                isCreateGroupSubscriptionModalOpen = false;
            }
        }} />

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

    <!-- svelte-ignore a11y-label-has-associated-control -->
    <NavBar showMenuToggle={false}>
        <svelte:fragment slot="navbar-start">
            <div class="flex flex-col md:flex-row">
                <a href="/" class="btn btn-ghost" aria-label="Homepage">
                    <div class="inline-block text-2xl font-title text-primary flex">
                        <img src={logo} class="w-8 h-8 mr-3" alt="Logo" />
                        <span class="lowercase">rag</span><span class="uppercase text-base-content">bag</span>
                    </div>
                </a>

                <div class="md:divider md:divider-horizontal hidden " />

                <div class="ml-4 md:ml-0">
                    <div class={classNames("text-xl lg:text-2xl font-bold flex", group?.description ? "" : "mt-2")}>
                        <div class="w-[30px] lg:w-[40px]">{group?.icon}</div>
                        <span class="line-clamp-1">{group?.name}</span>
                    </div>
                    <i class="text-xs leading-4 opacity-80 hidden lg:block ml-1">{group?.description}</i>
                </div>
            </div>
        </svelte:fragment>

        <svelte:fragment slot="navbar-end">
            {#if $authenticated}
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
                            <button on:click={() => (isCreateGroupSubscriptionModalOpen = true)}>
                                <Icon src={PlusCircle} class="h-5 w-5" />
                                Subscribe
                            </button>
                        </li>
                    </ul>
                </div>
            {/if}
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

    <ToastContainer placement="bottom-right" let:data>
        <FlatToast {data} />
    </ToastContainer>
{/if}
