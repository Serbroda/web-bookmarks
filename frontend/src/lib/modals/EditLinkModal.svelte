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

    export let link: LinkDto;

    export let onShow: () => void;
    export let onClose: () => void;
    export let onSave: (link: LinkDto) => void;
    export let onDelete: () => void = undefined;

    let selectedGroup: GroupDto;

    const handleSave = async () => {
        try {
            isBusy = true;

            const dto: CreateLinkDto = {
                url: link.url,
                name: link.name,
                description: link.description,
                groupId: selectedGroup?.id,
            };

            let response: LinkDto;
            response = await groupService.updateLink(link.id, dto);

            if (onSave) {
                onSave(response);
            }

            isOpen = false;
        } catch (err) {
            console.error(err);
        } finally {
            isBusy = false;
            selectedGroup = undefined;
        }
    };

    const _onShow = () => {
        selectedGroup = $groups.find((g) => g.id === link.groupId);
        if (onShow) {
            onShow();
        }
    };
</script>

<Modal {isOpen} {onClose} onShow={_onShow}>
    <div class="flex flex-col space-y-3 mt-2">
        {#if link}
            <input
                bind:value={link.name}
                type="text"
                placeholder="Link name (optional)"
                class="input hover:input-bordered font-semibold text-xl" />
            <a class="link flex invisble-hover-container ml-4 hover:text-primary" target="_blank" href={link.url}>
                <img class="favicon pr-1 pt-1" src={faviconUrl(hostname(link.url))} alt="favicon" />
                <span class="break-all">{link.url}</span>
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
                    <option value={group}>{group.icon} {group.name}</option>
                {/each}
            </select>

            <textarea
                class="textarea textarea-bordered w-full mt-2 leading-5"
                rows="3"
                placeholder="Description (optional)"
                bind:value={link.description} />
        {/if}
    </div>

    <svelte:fragment slot="actions">
        <div class="flex mt-4">
            <div class="space-x-2">
                <button on:click={onDelete} class="btn btn-error" disabled={isBusy}>
                    <Icon src={Trash} class="h-5 w-5" /> Delete
                </button>
            </div>
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
                <button on:click={onClose} class="btn gap-3">
                    <Icon src={X} class="h-5 w-5" /> Close
                </button>
            </div>
        </div>
    </svelte:fragment>
</Modal>
