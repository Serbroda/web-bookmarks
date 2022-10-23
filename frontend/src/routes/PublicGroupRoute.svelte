<script lang="ts">
    import { params } from "svelte-spa-router";
    import { groupService } from "../services/Services";
    import LinkCard from "../lib/LinkCard.svelte";
    import type { LinkDto } from "../models/LinkDto";
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
{/if}
