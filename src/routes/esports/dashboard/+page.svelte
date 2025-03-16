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
            <img src="/car-removed.png" alt="car" />
            <div class="stream">
                <a href="/dashboard/stream">Stream Now</a>
            </div>
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
    .stream {
        position: absolute;
        bottom: -95vh;
        left: 10vw;
        transform: translateX(-45%);
        padding: 1rem 2rem;
        border-radius: 10rem;
        text-align: center;
        background: white;
        color: black;
        backdrop-filter: blur(2px);
        -webkit-backdrop-filter: blur(2px);
        font-size: 1.2rem;
        font-weight: 500;
    }
</style>
