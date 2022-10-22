<script lang="ts">
    import { classNames } from "../../utils/dom";
    import { Icon, QuestionMarkCircle } from "svelte-hero-icons";
    import { version } from "../../stores/infos";

    export let classes = "";
</script>

<div class={classNames("dropdown", classes)}>
    <!-- svelte-ignore a11y-label-has-associated-control -->
    <label tabindex="0" class="btn btn-ghost btn-sm btn-circle">
        <Icon src={QuestionMarkCircle} class="h-5 w-5" />
    </label>

    <div tabindex="0" class="card compact dropdown-content shadow bg-base-100 rounded-box w-64">
        <div class="card-body">
            <span class="text-center font-mono line-clamp-1 text-xs">
                {$version}
            </span>
        </div>
    </div>
</div>
