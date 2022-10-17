<script lang="ts">
    import { authService } from "../services/Services.js";
    import { replace } from "svelte-spa-router";
    import { toasts } from "svelte-toasts";

    let error: string | undefined = undefined;

    async function handleSubmit(e) {
        try {
            const formData = new FormData(e.target);

            const response = await authService.register(formData);

            if(authService.isResponseOk(response)) {
                await replace("/");
            } else if (response.status === 409) {
                throw new Error(`The given email address already exists`)
            } else {
                throw new Error(response.statusText)
            }
        } catch (err) {
            error = err
        }
    }
</script>

<div class="flex min-h-full h-screen flex-col justify-center py-12 sm:px-6 lg:px-8 bg-base-300">
    <div class="sm:mx-auto sm:w-full sm:max-w-md">
        <img
            class="mx-auto h-12 w-auto"
            src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600"
            alt="Your Company" />
        <h2 class="mt-6 text-center text-3xl font-bold tracking-tight">Sign up</h2>
    </div>

    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
        <div class="bg-base-100 py-8 px-4 shadow sm:rounded-lg sm:px-10">
            <form class="space-y-6" on:submit|preventDefault={handleSubmit}>
                <div class="space-y-6">
                    <div class="form-control w-full">
                        <label class="label" for="username">
                            <span class="label-text">Email address</span>
                        </label>
                        <input
                            type="email"
                            id="username"
                            name="username"
                            placeholder="Email address"
                            class="input input-bordered w-full" />
                    </div>

                    <div class="form-control w-full">
                        <label class="label" for="password">
                            <span class="label-text">Password</span>
                        </label>
                        <input
                            type="password"
                            id="password"
                            name="password"
                            placeholder="Password"
                            class="input input-bordered w-full" />
                    </div>

                    <div class="form-control w-full">
                        <input
                            type="password"
                            id="passwordConfirm"
                            name="passwordConfirm"
                            placeholder="Password (repeat)"
                            class="input input-bordered w-full" />
                    </div>

                    <div class="text-center">
                        <p>Already have an account? <a class="link link-primary" href="/#/login">Sign in</a></p>
                    </div>

                    {#if error}
                        <p class="text-error">{error}</p>
                    {/if}

                    <div>
                        <button type="submit" class="btn btn-primary w-full">Sign up</button>
                    </div>
                </div>
            </form>
        </div>
    </div>
</div>
