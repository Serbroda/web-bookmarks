<script lang="ts">
    import { authService } from "../services/Services.js";
    import MainContainer from "../lib/MainContainer.svelte";
    import { replace } from "svelte-spa-router";
    import ConfirmModal from "../lib/modals/ConfirmModal.svelte";
    let username = "";
    let password = "";

    const login = async () => {
        const response = await authService.login(username, password);
        console.log(response);
        username = "";
        password = "";
        window.location = "/";
    };

    const logout = async () => {
        authService.logout();
        username = "";
        password = "";
        window.location = "/";
    };
</script>

<MainContainer>
    <h2 class="my-6 text-4xl font-bold">
        <span class="text-primary">Home</span>
    </h2>

    <button
        class="btn"
        on:click={() => {
            authService.logout();
            username = "";
            password = "";
            replace("/login");
        }}>
        <svg
            class="animate-spin -ml-1 mr-3 h-5 w-5 text-white"
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
</MainContainer>
