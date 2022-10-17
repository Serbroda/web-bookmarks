<script lang="ts">
    import EmojiChooser from "../EmojiChooser.svelte";
    import Modal from "./Modal.svelte";

    export let isOpen: boolean = false;
    export let name: string = "";
    export let description: string = "";
    export let selectedEmoji: string = "❓";
    export let onClose: () => void;
    export let onSave: (emoji: string, name: string, description?: string) => void;
    export let mode: "Add" | "Edit" = "Add";

    let nameElement: HTMLInputElement;

    const onShow = () => {
        if (nameElement) {
            nameElement.focus();
        }
    };
</script>

<Modal
    {isOpen}
    {onShow}
    title={mode + " Group"}
    buttons={[
        {
            content: "Save",
            onClick: () => {
                if (onSave) {
                    onSave(selectedEmoji, name, description);
                }
            },
            classes: "btn-primary",
        },
        {
            content: "Cancel",
            onClick: () => {
                if (onClose) {
                    onClose();
                }
            },
        },
    ]}
    {onClose}>
    <div class="flex">
        <EmojiChooser
            selected={selectedEmoji}
            classes="mr-2"
            onChange={(emoji) => {
                selectedEmoji = emoji;
            }} />
        <input
            bind:this={nameElement}
            bind:value={name}
            type="text"
            placeholder="Group name"
            class="input input-bordered grow" />
    </div>
    <textarea
        class="textarea textarea-bordered w-full mt-2"
        rows="2"
        placeholder="Description (optional)"
        bind:value={description} />
</Modal>
