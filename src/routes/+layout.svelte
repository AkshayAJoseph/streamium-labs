<script>
    import { slide } from "svelte/transition";
    import { page } from "$app/stores";
    import { quintOut } from "svelte/easing";
</script>

<svelte:head>
    <link rel="stylesheet" href="/main.css" />
</svelte:head>

{#key $page.url.pathname}
    <div>
        <slot />
    </div>
{/key}
