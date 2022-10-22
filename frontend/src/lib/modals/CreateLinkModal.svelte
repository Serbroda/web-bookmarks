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

    export let onShow: () => void = undefined;
    export let onClose: () => void = undefined;
    export let onSave: (link: LinkDto) => void = undefined;

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

            const response = await groupService.createLink(selectedGroup.id, dto);

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

        <textarea
            class="textarea textarea-bordered w-full mt-2 leading-5"
            rows="3"
            placeholder="Description (optional)"
            bind:value={description} />
    </div>

    <svelte:fragment slot="actions">
        <div class="flex mt-4">
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
