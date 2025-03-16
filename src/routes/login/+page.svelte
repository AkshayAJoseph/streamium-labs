<script>
    import Nav from "$lib/components/Nav.svelte";
    import { login } from "../../script";

    let email = $state("");
    let password = $state("");
    const onclick = async () => {
        const data = { email, password };
        await login(data);
    };
</script>

<main>
    <div class="box">
        <Nav />
        <div class="login">
            <div class="login__banner">
                <img src="/events.jpg" alt="" />
            </div>
            <div class="login__form">
                <h5>Login with ease</h5>
                <div class="login__row">
                    <label>NAME</label>
                    <input
                        type="text"
                        placeholder="Hari Narayan"
                        bind:value={email}
                    />
                </div>
                <div class="login__row">
                    <label>PASSWORD</label>
                    <input
                        type="password"
                        placeholder="********"
                        bind:value={password}
                    />
                </div>
                <div class="login__row">
                    <input type="submit" value="Submit" />
                    <p>
                        Don't have an account? <a href="/register"
                            >Register Here</a
                        >
                    </p>
                </div>
            </div>
        </div>
    </div>
</main>

<style>
    .box {
        min-height: 100vh;
    }

    .login {
        margin-top: 5rem;
        display: flex;
        justify-content: space-evenly;
        align-items: center;
        height: 100vh;
    }

    .login h5 {
        font-size: 2rem;
        margin-bottom: 2rem;
    }

    .login > * {
        height: 80%;
    }

    .login__banner {
        width: 60%;
    }

    .login__form {
        width: 30%;
        background-color: #060606;
        margin-right: 5rem;
        padding: 2rem;
        border-radius: 1rem;
        border: 3px solid transparent;
        transition: border 0.3s ease;
        position: relative;
    }

    .login__form:hover {
        border: 3px solid var(--color-labs);
    }

    .login__form label {
        display: block;
        margin-bottom: 0.5rem;
        font-size: 0.5rem;
    }

    .login__row {
        margin-bottom: 2rem;
    }

    .login__row input {
        outline: 0;
        width: 100%;
        padding: 1rem;
        border-radius: 0.5rem;
        background: #111;
        border: 0;
        font-family: "Inter";
        font-size: 1rem;
        color: white;
    }

    .login__row input[type="submit"] {
        border-radius: 10rem;
        background: white;
        color: black;
        font-weight: 600;
    }

    .login__row:last-child {
        position: absolute;
        width: calc(100% - 4rem);
        bottom: 0;
    }

    .login__row:last-child p {
        margin-top: 1rem;
        text-align: center;
    }
    .login__row:last-child a {
        display: inline-block;
        color: white;
    }
    .login__banner img {
        width: 100%;
        height: 100%;
        object-fit: cover;
        animation: rotateImage 20s linear infinite; /* Added animation */
    }

    @keyframes rotateImage {
        from {
            transform: rotate(0deg);
        }
        to {
            transform: rotate(360deg);
        }
    }

    @media only screen and (max-width: 640px) {
        .login {
            flex-direction: column;
        }
        .login > * {
            width: 100%;
        }
        .login__banner {
            width: 50%;
        }
        .login__form {
            margin-right: 0;
            margin: 2rem;
            width: calc(100% - 2rem);
            padding-bottom: 20vh;
        }
    }
</style>
