<script lang="ts">
    import { authService } from "../services/Services.js";
    import { onMount } from "svelte";
    import MainContainer from "../lib/MainContainer.svelte";
    import { replace } from "svelte-spa-router";
    import NavBar from "../lib/NavBar.svelte";
    import { groupService } from "../services/Services";
    import type { GroupDto } from "src/models/GroupDto.js";
    import Card from "../lib/Card.svelte";

    let groups: GroupDto[] = [];

    onMount(async () => {
        groups = await groupService.getGroups("created_at desc", "6");
    });
</script>

<MainContainer>
    <NavBar>
        <svelte:fragment slot="navbar-start">
            <h2 class="text-xl lg:text-2xl font-bold flex ml-0 lg:ml-3">
                <span class="text-primary">Home</span>
            </h2>
        </svelte:fragment>
    </NavBar>

    <div id="content" class="p-4">
        {#if groups}
            <h2 class="text-xl font-bold opacity-60">Recently added Groups</h2>

            <div class="flex flex-wrap mt-4">
                {#each groups || [] as group}
                    <Card
                        onClick={async () => {
                            await replace(`/groups/${group.id}`);
                        }}>
                        <div class="leading-none">
                            <div class="font-bold flex">
                                <div class="w-[30px]">{group?.icon}</div>
                                <span class="line-clamp-1">{group?.name}</span>
                            </div>
                            <div class="text-sm flex leading-none mt-2">
                                <div>{group.description}</div>
                            </div>
                        </div>
                    </Card>
                {/each}
            </div>
        {/if}
    </div>
</MainContainer>
