<script lang="ts">
    import { hostname } from "../../utils/url";
    import Modal from "./Modal.svelte";
    import { faviconUrl } from "../../utils/url.js";
    import { Icon, Trash, X } from "svelte-hero-icons";
    import { groups } from "../../stores/groups.js";
    import { groupService } from "../../services/Services";
    import type { CreateLinkDto, LinkDto } from "src/models/LinkDto";
    import type { GroupDto } from "src/models/GroupDto";

    export let isOpen: boolean = false;
    export let isBusy: boolean = false;

    export let url: string = "";
    export let name: string = "";
    export let description: string = "";
    export let selectedGroup: GroupDto = undefined;
    export let linkId: string = "";

    export let onShow: () => void;
    export let onClose: () => void;
    export let onSave: (link: LinkDto) => void;
    export let onDelete: () => void = undefined;
    export let mode: "Add" | "Edit" = "Add";

    let urlElement: HTMLInputElement;

    const handleSave = async () => {
        try {
            isBusy = true;

            const dto: CreateLinkDto = {
                url,
                name,
                description,
                groupId: selectedGroup?.id,
            };

            let response: LinkDto;
            if (mode === "Edit" && linkId) {
                response = await groupService.updateLink(linkId, dto);
            } else {
                response = await groupService.createLink(selectedGroup.id, dto);
            }

            if (onSave) {
                onSave(response);
            }

            url = "";
            urlElement.value = "";

            isOpen = false;
        } catch (err) {
            console.error(err);
        } finally {
            isBusy = false;
        }
    };

    const _onShow = () => {
        if (urlElement) {
            urlElement.focus();
        }
        if (onShow) {
            onShow();
        }
    };
</script>

<Modal {isOpen} {onClose} onShow={_onShow}>
    <div class="flex flex-col space-y-3 mt-2">
        <input
            bind:value={name}
            type="text"
            placeholder="Link name (optional)"
            class="input hover:input-bordered font-semibold text-xl" />
        {#if mode === "Add"}
            <input
                bind:this={urlElement}
                bind:value={url}
                on:keydown={async (e) => {
                    if (e.key === "Enter") {
                        await handleSave();
                    }
                }}
                type="text"
                placeholder="https://example.com"
                class="input input-bordered" />
        {:else}
            <a class="link flex invisble-hover-container ml-4 hover:text-primary" target="_blank" href={url}>
                <img class="favicon pr-1 pt-1" src={faviconUrl(hostname(url))} alt="favicon" />
                {url}
                <svg
                    class="w-4 h-4 ml-1 invisble-hover-item"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                    xmlns="http://www.w3.org/2000/svg">
                    <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
                </svg>
            </a>

            <div class="divider" />
            <select class="select select-bordered w-full max-w-xs" bind:value={selectedGroup}>
                {#each $groups as group}
                    <option value={group} selected={selectedGroup?.id === group.id}>{group.icon} {group.name}</option>
                {/each}
            </select>
        {/if}
        <textarea
            class="textarea textarea-bordered w-full mt-2 leading-5"
            rows="3"
            placeholder="Description (optional)"
            bind:value={description} />
    </div>

    <svelte:fragment slot="actions">
        <div class="flex mt-4">
            {#if mode === "Edit"}
                <div class="space-x-2">
                    <button on:click={onDelete && onDelete()} class="btn btn-error" disabled={isBusy}>
                        <Icon src={Trash} class="h-5 w-5" /> Delete
                    </button>
                </div>
            {/if}
            <div class="flex-1" />
            <div class="space-x-2">
                <button on:click={handleSave} class="btn btn-primary gap-3" class:loading={isBusy} disabled={isBusy}>
                    {#if !isBusy}
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="feather feather-save h-5 w-5"
                            ><path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z" /><polyline
                                points="17 21 17 13 7 13 7 21" /><polyline points="7 3 7 8 15 8" /></svg>
                    {/if}
                    Save
                </button>
                <button on:click={onClose && onClose()} class="btn gap-3">
                    <Icon src={X} class="h-5 w-5" /> Close
                </button>
            </div>
        </div>
    </svelte:fragment>
</Modal>
