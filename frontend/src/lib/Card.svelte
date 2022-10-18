<script lang="ts">
    export let onClick: () => void = undefined;
</script>

<div
    class="card invisble-hover-container hover:border-primary border-2 border-transparent shadow cursor-pointer bg-base-200 grow w-full md:grow-0 md:w-80 m-2 bg-base-200"
    on:click={onClick}>
    <div class="card-body px-3 py-4">
        <slot />
    </div>
</div>
