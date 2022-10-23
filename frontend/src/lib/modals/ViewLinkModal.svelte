<script lang="ts">
    import { hostname } from "../../utils/url";
    import Modal from "./Modal.svelte";
    import { faviconUrl } from "../../utils/url.js";
    import { Icon, X } from "svelte-hero-icons";
    import type { LinkDto } from "src/models/LinkDto";

    export let isOpen: boolean = false;

    export let link: LinkDto;

    export let onShow: () => void;
    export let onClose: () => void;

    const _onShow = () => {
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
                readonly
                placeholder="Link name (optional)"
                class="input input-ghost hover:input-bordered font-semibold text-xl" />
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
            <textarea
                class="textarea textarea-bordered w-full mt-2 leading-5"
                rows="3"
                readonly
                placeholder="Description (optional)"
                bind:value={link.description} />
        {/if}
    </div>

    <svelte:fragment slot="actions">
        <div class="flex mt-4">
            <div class="flex-1" />
            <div class="space-x-2">
                <button on:click={onClose} class="btn gap-3">
                    <Icon src={X} class="h-5 w-5" /> Close
                </button>
            </div>
        </div>
    </svelte:fragment>
</Modal>
