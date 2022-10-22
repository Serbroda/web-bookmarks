<script context="module" lang="ts">
    export type ZIndex = 0 | 10 | 20 | 30 | 40 | 50;
    export type ModalSize = "sm" | "normal";

    export interface ModalButton {
        content: any;
        onClick: () => void;
        classes?: string;
        position?: "left" | "right";
        condition?: () => boolean;
    }
</script>

<script lang="ts">
    import { afterUpdate, onMount } from "svelte";
    import { classNames } from "../../utils/dom.js";
    import { generate } from "shortid";

    export let isOpen: boolean = false;

    export let title: string = "";
    export let buttons: ModalButton[] = [];
    export let classes: string = "";
    export let size: ModalSize = "normal";
    export let zIndex: ZIndex = 10;

    export let onClose: () => void = undefined;
    export let onShow: () => void = undefined;

    export const show = () => {
        isOpen = true;
    };

    export const hide = () => {
        isOpen = false;
    };

    const containerId = `__modal-${generate()}`;

    onMount(async () => {
        const a = document.getElementById(containerId);

        const modalContainer = `__modal-container-${zIndex}`;
        if (modalContainer) {
            const container = document.getElementById(modalContainer);
            if (container) {
                container.appendChild(a);
            } else {
                document.body.appendChild(a);
            }
        } else {
            document.body.appendChild(a);
        }
    });

    let isOpening: boolean = false;

    $: if (isOpen) {
        isOpening = true;
    }

    afterUpdate(() => {
        if (isOpening) {
            if (onShow) {
                onShow();
            }
            isOpening = false;
        }
    });
</script>

<div id={containerId} class={classes}>
    <div class="modal items-start pt-[10%] lg:pt-[4%]" class:modal-open={isOpen}>
        <div class={classNames("modal-box overflow-visible", size === "normal" ? "w-11/12 max-w-5xl" : "")}>
            <button
                class="btn btn-ghost btn-sm btn-circle absolute right-2 top-2"
                on:click={() => {
                    if (onClose) {
                        onClose();
                    }
                    isOpen = false;
                }}
                >✕
            </button>

            {#if title}
                <h3 class="text-lg font-bold">{title}</h3>
                <br />
            {/if}

            <slot />

            {#if buttons && buttons.length > 0}
                <div class="flex mt-4">
                    <div class="space-x-2">
                        {#each buttons.filter((b) => b.position === "left" && (b.condition === undefined || b.condition())) as btn}
                            <button class={classNames("btn", btn.classes)} on:click={() => btn.onClick()}>
                                {@html btn.content}
                            </button>
                        {/each}
                    </div>

                    <div class="flex-1" />

                    <div class="space-x-2">
                        {#each buttons.filter((b) => b.position === undefined || (b.position === "right" && (b.condition === undefined || b.condition()))) as btn}
                            <button class={classNames("btn", btn.classes)} on:click={() => btn.onClick()}>
                                {@html btn.content}
                            </button>
                        {/each}
                    </div>
                </div>
            {:else}
                <slot name="actions" />
            {/if}
        </div>
    </div>
</div>

<style>
    .modal .modal-box {
    }
</style>
