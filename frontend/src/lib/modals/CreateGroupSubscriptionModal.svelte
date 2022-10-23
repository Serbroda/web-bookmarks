<script lang="ts">
    import Modal from "./Modal.svelte";
    import { Icon, X } from "svelte-hero-icons";
    import { groupService } from "../../services/Services";
    import type { GroupSubscriptionDto } from "../../models/GroupSubscriptionDto";

    export let isOpen: boolean = false;
    export let value: string = "";
    export let readonly: boolean = false;
    export let onClose: () => void;
    export let onError: (error: any) => void = undefined;
    export let onSuccess: (subscription: GroupSubscriptionDto) => void;

    const createGroupSubscription = async (linkOrId: string) => {
        try {
            let id = linkOrId;
            const regex: RegExp = /^.*public\/groups\/(.*)$/;
            if (regex.test(id)) {
                const match = id.match(regex);
                id = match[1];
            }

            const response = await groupService.createGroupSubscription(id);
            if (response !== undefined) {
                onSuccess(response);
            }
        } catch (error) {
            if (onError) {
                onError(error);
            }
        }
    };
</script>

<Modal {isOpen} title="Subscribe to an external Group" {onClose}>
    <div class="form-control">
        <input bind:value type="text" class="input input-bordered w-full" placeholder="Link or Group-Id" {readonly} />
    </div>

    <svelte:fragment slot="actions">
        <div class="flex mt-4">
            <div class="flex-1" />
            <div class="space-x-2">
                <button on:click={async (e) => await createGroupSubscription(value)} class="btn btn-primary gap-3"
                    ><svg
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
                            points="17 21 17 13 7 13 7 21" /><polyline points="7 3 7 8 15 8" /></svg> Save
                </button>
                <button on:click={onClose} class="btn gap-3">
                    <Icon src={X} class="h-5 w-5" /> Close
                </button>
            </div>
        </div>
    </svelte:fragment>
</Modal>
