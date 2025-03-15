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
                <h5>WHO<span>DOOOOGE </span> DI.</h5>
                <h1>MARKETING</h1>
            </div>
            <img src="/girl-removed.png" alt="car" />
            <div class="hero__box"></div>
        </article>
    </section>
</main>

<style>
    section {
        background: url("/girl.jpg");
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
            rgba(0, 0, 0, 0.2),
            rgba(0, 0, 0, 0.5)
        );
        animation: breathe 5s infinite linear;
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
        text-align: center;
        padding-top: 10rem;
        font-size: 12rem;
    }

    .hero__text h1,
    .hero__text h5 {
        line-height: 1;
        font-family: "Inter";
    }

    .hero__text h5 {
        display: flex;
        justify-content: space-between;
    }

    .hero__text h5 span {
        font-style: italic;
        color: var(--color-labs);
    }
</style>
