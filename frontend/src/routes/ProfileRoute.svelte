<script lang="ts">
    import { authService } from "../services/Services";
    import MainContainer from "../lib/MainContainer.svelte";
    import NavBar from "../lib/NavBar.svelte";
    import { replace } from "svelte-spa-router";

    let error;

    async function handleSubmit(e) {
        try {
            error = undefined;

            const formData = new FormData(e.target);
            const response = await authService.changePassword(formData);

            if (authService.isResponseOk(response)) {
                authService.logout();
                await replace("/login");
            } else if (response.status === 403) {
                throw new Error(`Forbidden`);
            } else {
                throw new Error(response.statusText);
            }
        } catch (err) {
            error = err;
        }
    }
</script>

<MainContainer>
    <NavBar>
        <svelte:fragment slot="navbar-start">
            <h2 class="text-xl lg:text-2xl font-bold flex ml-0 lg:ml-3">
                <span class="text-primary">Profile</span>
            </h2>
        </svelte:fragment>
    </NavBar>

    <div id="content" class="card p-4" on:submit|preventDefault={handleSubmit}>
        <div class="card bg-base-200 shadow-md border">
            <h3 class="font-bold mt-4 ml-4">Change password</h3>
            <form class="space-y-6 card-body">
                <div class="form-control w-full">
                    <label class="label" for="oldPassword">
                        <span class="label-text">Current password</span>
                    </label>
                    <input
                        type="password"
                        id="oldPassword"
                        name="oldPassword"
                        placeholder="Current password"
                        class="input input-bordered w-full"
                        required />
                </div>

                <div class="form-control w-full">
                    <label class="label" for="newPassword">
                        <span class="label-text">New password</span>
                    </label>
                    <input
                        type="password"
                        id="newPassword"
                        name="newPassword"
                        placeholder="New password"
                        class="input input-bordered w-full"
                        required />
                </div>

                <div class="form-control w-full">
                    <input
                        type="password"
                        id="newPasswordConfirm"
                        name="newPasswordConfirm"
                        placeholder="New password (repeat)"
                        class="input input-bordered w-full" />
                </div>

                {#if error}
                    <p class="text-error">{error}</p>
                {/if}

                <div>
                    <button type="submit" class="btn btn-primary w-24">Save</button>
                </div>
            </form>
        </div>
    </div>
</MainContainer>
