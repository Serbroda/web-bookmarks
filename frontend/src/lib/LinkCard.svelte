<script lang="ts">
    import moment from "moment";
    import { toasts } from "svelte-toasts";
    import { hostname } from "../utils/url";
    import { faviconUrl } from "../utils/url.js";
    import type { LinkDto } from "../models/LinkDto";

    // https://stackoverflow.com/questions/10282939/how-to-get-favicons-url-from-a-generic-webpage-in-javascript

    export let item: LinkDto;
    export let onClick: () => void;

    const onCardClick = (e: Event) => {
        if (onClick) {
            onClick();
        }
    };

    const onIconClick = (e: Event) => {
        e.stopPropagation();
        window.open(item.url, "_blank").focus();
    };
</script>

{#if item}
    {@const host = hostname(item.url)}
    <div
        class="card invisble-hover-container hover:border-primary border-2 border-transparent shadow cursor-pointer bg-base-200 grow w-full md:grow-0 md:w-80 m-2 bg-base-200"
        on:click={onCardClick}>
        <div class="card-body px-3 py-4">
            <div class="leading-none">
                <span class="font-semibold line-clamp-2">{item.name}</span>
                <div class="text-sm flex leading-none mt-2">
                    <img class="favicon favicon-sm pr-1" src={faviconUrl(host)} alt="favicon" />
                    <div>{host}</div>
                </div>
                <i class="text-xs text-gray-500">{moment(item.updatedAt).format("DD.MM.YYYYY, HH:mm")}</i>
            </div>
        </div>

        <div class="invisble-hover-item absolute bottom-1 right-2">
            {#if navigator.clipboard !== undefined}
                <button
                    on:click={(e) => {
                        e.stopPropagation();

                        navigator.clipboard.writeText(item.url).then(
                            function () {
                                toasts.success("Link copied to clipboard");
                            },
                            function (err) {
                                console.error("Async: Could not copy text: ", err);
                            },
                        );
                    }}
                    class="hover:text-primary">
                    <svg
                        class="w-5 h-5"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                        xmlns="http://www.w3.org/2000/svg">
                        <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
                    </svg>
                </button>
            {/if}

            <button on:click={onIconClick} class="hover:text-primary">
                <svg
                    class="w-5 h-5"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                    xmlns="http://www.w3.org/2000/svg">
                    <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
                </svg>
            </button>
        </div>
    </div>
{/if}
