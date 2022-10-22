<script lang="ts">
    import { classNames } from "../utils/dom";
    import { Icon, Menu } from "svelte-hero-icons";

    export let showMenuToggle: boolean = true;
    export let classes: string = "";
</script>

<div class={classNames("navbar bg-base-100 shadow sticky top-0 z-10 ", classes)}>
    <div class="navbar-start w-full">
        {#if showMenuToggle}
            <label for="main-menu" class="btn btn-ghost btn-circle lg:hidden">
                <Icon src={Menu} class="h-5 w-5" />
            </label>
        {/if}

        <slot name="navbar-start" />
    </div>
    <!--<div class="navbar-center">
        <slot name="navbar-center" />
    </div>-->
    <div class="navbar-end">
        <slot name="navbar-end" />
    </div>
</div>
