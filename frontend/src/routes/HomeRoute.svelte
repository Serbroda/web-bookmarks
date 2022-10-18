<script lang="ts">
    import { authService } from "../services/Services.js";
    import MainContainer from "../lib/MainContainer.svelte";
    import { replace } from "svelte-spa-router";
    import NavBar from "../lib/NavBar.svelte";
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
        <button
            class="btn"
            on:click={() => {
                authService.logout();
                replace("/login");
            }}>
            <svg
                class="-ml-1 mr-3 h-5 w-5 text-white"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
                <path
                    class="opacity-75"
                    fill="currentColor"
                    d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z" />
            </svg>
            Logout
        </button>
    </div>
</MainContainer>
