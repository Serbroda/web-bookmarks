<script lang="ts">
    import Modal from "./Modal.svelte";
    import { Icon, ClipboardCopy } from "svelte-hero-icons";
    import { copyToClipboard } from "../../utils/clipboard";
    import { toasts } from "svelte-toasts";

    export let isOpen: boolean = false;
    export let onClose: () => void;
    export let title: string = "Share";
    export let message: string = "Copy the URL below to share with anyone!";
    export let link: string;
</script>

<Modal {isOpen} {title} {onClose}>
    <p>{message}</p>
    <br />
    <div class="form-control">
        <div class="input-group">
            <input bind:value={link} type="text" class="input input-bordered w-full" readonly />
            <button
                class="btn btn-square"
                disabled={link === undefined || link.trim() === ""}
                on:click={() => {
                    if (link) {
                        copyToClipboard(link);
                        toasts.success("Link copied");
                    }
                }}>
                <Icon src={ClipboardCopy} class="h-6 w-6" />
            </button>
        </div>
    </div>
</Modal>
