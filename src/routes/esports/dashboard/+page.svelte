<script>
    import Nav from "$lib/components/Nav.svelte";
</script>

<main>
    <div class="box"></div>
    <section>
        <div class="overlay"></div>
        <Nav />
        <article class="hero">
            <div class="hero__text">
                <h1>GET READY TO</h1>
                <h5><span>KILL. </span> WIN. RACE.</h5>
            </div>
            <img src="/car-removed-nobs.png" alt="car" />
        </article>
    </section>
</main>

<style>
    section {
        background: url("/car-nobs.jpeg");
        background-size: cover;
        height: 100vh;
        position: relative; /* Needed for absolute positioning of overlay */
    }

    .overlay {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background: linear-gradient(
            to bottom,
            rgba(0, 0, 0, 0.1),
            rgba(0, 0, 0, 0.5)
        );
        animation: breathe 2s infinite linear;
    }

    @keyframes breathe {
        0%,
        100% {
            opacity: 1;
        }
        50% {
            opacity: 0.2;
        }
        25%,
        75% {
            opacity: 0.5;
        }
    }

    /* Rest of your CSS remains the same */

    .hero {
        position: relative;
    }

    .hero img {
        position: absolute;
        top: 0;
        width: 100vw;
    }

    .hero__text {
        position: absolute;
        top: 20vh;
        left: 10vw;
        transform: rotate(-1deg);
        font-size: 10rem;
    }

    .hero__text h1,
    .hero__text h5 {
        line-height: 1;
        font-family: "Inter";
    }

    .hero__text h5 span {
        font-style: italic;
        color: var(--color-esports);
    }
</style>
