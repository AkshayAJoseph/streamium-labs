<script>
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";
    import { fade } from "svelte/transition";

    let showLoader = false;
    let showContent = true;

    onMount(() => {
        setTimeout(() => {
            showLoader = true;
            setTimeout(() => {
                const loader = document.querySelector(".loader");
                loader.style.setProperty("--loader-width", "100%");
                setTimeout(() => {
                    showContent = false;
                    setTimeout(() => {
                        goto("/esports");
                    }, 300);
                }, 5000);
            }, 10);
        }, 10);
    });
</script>

{#if showLoader && showContent}
    <main transition:fade={{ duration: 300 }}>
        <div class="box">
            <section>
                <div class="logo">
                    <p>THE</p>
                    <h5>STREAMIUM</h5>
                    <p class="end esports">ESPORTS</p>
                    <div class="loader">hey</div>
                </div>
            </section>
        </div>
    </main>
{/if}

<style>
    section {
        height: 100vh;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }

    .loader {
        margin-top: 2.5rem;
        width: 100%;
        font-size: 1rem;
        color: transparent;
        background-color: #111;
        position: relative;
        border-radius: 1rem;
        overflow: hidden;
    }

    .loader::before {
        border-radius: 1rem;
        content: "";
        position: absolute;
        display: block;
        height: 100%;
        width: var(--loader-width, 20%);
        background-color: var(--color-esports, #00ffff);
        transition: width 5s ease-out;
    }
</style>
