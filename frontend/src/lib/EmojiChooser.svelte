<script lang="ts">
    import { random } from "../utils/number";
    import { classNames } from "../utils/dom.js";
    import { emojies } from "../consts/emojies.js";

    export let classes: string = "";
    export let onChange: (emoji: string) => void;
    export let selected: string = "❓";

    const rnd: number = random(1, 99999999);
    const id = `emojichooser__${rnd}`;

    let search = "";
</script>

<div {id} class={classNames("dropdown", classes)}>
    <button tabindex="0" class="btn btn-ghost border border-gray-300 text-xl">
        <span style="min-width: 20px">{selected}</span>
    </button>

    <div tabindex="0" class="dropdown-content menu shadow bg-base-100 rounded-box h-64 w-52 py-2 pl-2 pr-1">
        <div class="overflow-x-auto">
            <div class="p-1 sticky top-0">
                <input
                    type="search"
                    bind:value={search}
                    class="input input-bordered input-xs w-full"
                    placeholder="Search" />
            </div>
            <div class="flex flex-wrap text-xl">
                {#each emojies.filter((e) => !search || e.name.toLowerCase().includes(search.toLowerCase())) as emoji}
                    <span
                        class="p-1 cursor-pointer hover:bg-base-200"
                        style="min-width: 20px"
                        title={emoji.name}
                        on:click={() => {
                            selected = emoji.value;
                            if (onChange) {
                                onChange(emoji.value);
                            }
                        }}>{emoji.value}</span>
                {/each}
            </div>
        </div>
    </div>
</div>
