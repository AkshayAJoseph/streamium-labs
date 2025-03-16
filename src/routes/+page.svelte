<script></script>

<main>
    <div class="box">
        <section>
            <a href="/preloader-labs">
                <article>
                    <div class="logo">
                        <p>THE</p>
                        <h5>STREAMIUM</h5>
                        <p class="end labs">LABS</p>
                    </div>
                </article>
            </a>
            <a href="/preloader-esports">
                <article>
                    <div class="logo">
                        <p>THE</p>
                        <h5>STREAMIUM</h5>
                        <p class="end esports">ESPORTS</p>
                    </div>
                </article>
            </a>
        </section>
    </div>
</main>

<style>
    section {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100vh;
    }

    article {
        padding: 15rem;
        background: #000;
    }

    a:last-child article {
        border-left: 2px solid #111;
    }

    .logo {
        line-height: 1;
        transition: transform ease-in 0.3s;
    }

    .logo:hover {
        transform: rotate(-1deg);
    }

    @media only screen and (max-width: 640px) {
        section {
            flex-direction: column;
        }
        article {
            padding: 10rem;
        }
        a:last-child article {
            border-left: 2px solid transparent;
            border-top: 2px solid #111;
        }
    }
</style>
