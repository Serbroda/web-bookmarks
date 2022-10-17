<script context="module" lang="ts">
    export interface ConfirmModalPros {
        title: string;
        message: string;
        buttons?: ModalButton[];
        onClose?: () => void;
        classes?: string;
    }
</script>

<script lang="ts">
    import Modal from "./Modal.svelte";
    import type { ModalButton } from "./Modal.svelte";

    export let props: ConfirmModalPros;

    let modal: Modal;

    export const show = (p?: ConfirmModalPros) => {
        if (p) {
            props = p;
        }
        modal.show();
    };

    export const hide = () => {
        modal.hide();
    };
</script>

<Modal
    bind:this={modal}
    title={props?.title}
    buttons={props?.buttons}
    onClose={props?.onClose}
    classes={props?.classes}
    size="sm"
    zIndex={30}>
    {props?.message}
</Modal>
