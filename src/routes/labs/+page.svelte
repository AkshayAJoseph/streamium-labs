<script>
    import Nav from "$lib/components/Nav.svelte";
</script>

<main>
    <div class="box">
        <section>
            <div class="overlay"></div>
            <Nav />
            <article class="hero">
                <div class="hero__text">
                    <h5>WHO<span>DOOOOGE </span> DI.</h5>
                    <h1>MARKETING</h1>
                </div>
                <img src="/girl-removed.png" alt="car" />
            </article>
        </section>
        <div class="about" id="about">
            <div class="about__title">
                <h1>DE NOBIS</h1>
            </div>
            <div class="about__info">
                <div class="about__text">
                    <h5>Who are we?</h5>
                    <p>
                        We are Streamium Labs, a collective of digital
                        innovators and creators dedicated to crafting immersive
                        experiences in the modern media landscape. We specialize
                        in bridging the gap between audiences and content
                        through cutting-edge live streaming solutions,
                        high-quality media production, and meticulously
                        orchestrated event management. Our passion lies in
                        pushing the boundaries of digital engagement,
                        transforming ideas into captivating realities.
                        <br /><br />From dynamic esports broadcasts to impactful
                        corporate events and compelling media productions, we
                        are committed to delivering exceptional results that
                        resonate and inspire. We are the architects of digital
                        experiences, building the future of media, one stream at
                        a time.
                    </p>
                </div>
                <div class="about__banner">
                    <img src="/camera.jpg" alt="" />
                </div>
            </div>
        </div>
        <div class="services" id="services">
            <div class="services__title">
                <h1>LES SERVICES</h1>
            </div>
            <div class="services__row">
                <div class="services__column super">
                    <div class="services__banner">
                        <img src="/events.jpg" alt="" />
                    </div>
                    <div class="services__text">
                        <h5>Event Management</h5>
                        <p>
                            Our event management transforms gatherings into
                            immersive experiences, meticulously orchestrating
                            every detail to create moments that resonate.
                        </p>
                    </div>
                </div>
                <div class="services__column">
                    <div class="services__banner">
                        <img src="/media.jpg" alt="" />
                    </div>
                    <div class="services__text">
                        <h5>Media Production</h5>
                        <p>
                            Crafting compelling visual narratives through
                            high-quality video and audio production, from
                            concept to final delivery.
                        </p>
                    </div>
                </div>
                <div class="services__column">
                    <div class="services__banner">
                        <img src="/stream.jpg" alt="" />
                    </div>
                    <div class="services__text">
                        <h5>Live Streaming</h5>
                        <p>
                            We deliver dynamic, real-time connections, bridging
                            distances with engaging live content.
                        </p>
                    </div>
                </div>
            </div>
        </div>
    </div>
</main>

<style>
    .box {
        scroll-snap-type: y mandatory;
        height: 100vh;
        overflow-y: scroll;
        scroll-padding: 5rem;
        scroll-behavior: smooth;
    }
    .box > * {
        scroll-snap-align: start;
    }
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
        height: 100vh;
        object-fit: cover;
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

    .services__title h1,
    .about__title h1 {
        font-size: 5rem;
        text-align: center;
        margin: 5rem 0;
        font-weight: 300;
        letter-spacing: 0.5rem;
    }

    .about__info {
        display: flex;
    }

    .about__info > * {
        width: 50%;
    }

    .about__text {
        margin: 5rem;
    }

    .about__text h5 {
        font-size: 2.5rem;
        margin-bottom: 1rem;
    }

    .about__text p {
        font-size: 1.2rem;
        line-height: 1.8;
    }

    .about__banner img {
        width: 100%;
        height: 70vh;
        object-fit: cover;
        filter: brightness(0.7);
    }
    .services__row {
        padding: 5rem;
        display: flex;
        justify-content: space-between;
    }

    .services__row > * {
        width: 30%;
        height: 100%;
    }

    .services__column {
        border: 3px solid transparent;
        border-radius: 1rem;
        transition: border ease 0.3s;
    }

    .services__column:hover {
        border: 3px solid var(--color-labs);
    }

    .services__banner img {
        width: 100%;
        height: 100%;
        object-fit: cover;
        filter: saturate(0);
        transition: filter 0.3s ease;
        border-top-left-radius: 1rem;
        border-top-right-radius: 1rem;
    }

    .services__banner img:hover {
        filter: saturate(1);
    }

    .services__text {
        padding: 1rem;
        background-color: #060606;
        border-bottom-left-radius: 1rem;
        border-bottom-right-radius: 1rem;
    }

    .services__text h5 {
        font-size: 1.2rem;
        line-height: 1.5;
        margin-bottom: 1rem;
    }

    .services__text p {
        line-height: 1.8;
    }
    @media only screen and (max-width: 640px) {
        .box {
            scroll-snap-type: initial;
        }
        section {
            background-position: center;
        }
        .hero img {
            object-fit: cover;
        }
        .hero__text {
            padding-top: 5rem;
        }
        .hero__text h1 {
            font-size: 4rem;
        }
        .hero__text h5 {
            font-size: 2rem;
        }
        .about__info {
            flex-direction: column;
        }
        .about__info > * {
            width: calc(100% - 4rem);
            margin-left: 2rem;
        }
        .about__title h1,
        .services__title h1 {
            font-size: 3rem;
            margin-bottom: 0;
        }
        .about__banner img {
            height: 40vh;
        }
        .services__row {
            flex-direction: column;
            padding: 2rem;
        }
        .services__column {
            width: 100%;
            margin-bottom: 2rem;
        }
    }
</style>
