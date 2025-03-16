<script>
    export let selected;
</script>

<main>
    <nav>
        <div class="nav__logo">
            <div class="logo">
                <p>THE</p>
                <h5>STREAMIUM</h5>
                <p class="labs end">LABS</p>
            </div>
        </div>
        <div class="nav__links">
            <a href="/labs/#about">About</a>
            <a href="/labs/#services">Services</a>
            <a href="/labs/#contact">Contact</a>
            <a href="/blog">Blog</a>
            <a href="/login" class="solid">Login</a>
        </div>
    </nav>
</main>

<style>
    nav {
        padding: 2rem;
        display: flex;
        justify-content: space-between;
        width: 100%;
        z-index: 100;
        transition: top 0.3s ease; /* Transition only top */
        background: linear-gradient(
            to bottom,
            rgba(0, 0, 0, 0.5),
            rgba(0, 0, 0, 0)
        );
        position: fixed;
        top: 0;
    }

    .logo p {
        font-size: 0.5rem;
    }

    .logo h5 {
        font-size: 1rem;
    }

    .nav__links {
        display: flex;
        align-items: center;
        gap: 2rem;
    }

    .nav__links a {
        font-size: 0.9rem;
        font-weight: 500;
    }

    .nav__links a:hover {
        text-decoration: underline;
        text-underline-offset: 0.25rem;
    }

    .solid {
        background: white;
        color: black;
        height: min-content;
        padding: 0.5rem 1rem;
        border-radius: 1rem;
    }
    @media only screen and (max-width: 640px) {
        nav {
            justify-content: center;
            padding: 1rem;
        }
        .logo h5 {
            font-size: 1.5rem;
        }
        .logo p {
            font-size: 0.75rem;
        }
        .nav__links {
            display: none;
        }
    }
</style>
