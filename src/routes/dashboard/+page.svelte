<script>
    import Sidebar from "$lib/components/Sidebar.svelte";
</script>

<main>
    <div class="box">
        <Sidebar />
        <section>
            <h1>Dashboard</h1>
            <div class="logo">
                <p>THE</p>
                <h5>STREAMIUM</h5>
                <p class="labs end">LABS</p>
            </div>
            <article class="card">
                <h5>Welcome to your dashboard!</h5>
                <p>Manage your services here</p>
            </article>
        </section>
    </div>
</main>

<style>
    .box {
        display: flex;
    }
    section {
        padding: 2rem;
        width: 100%;
    }

    section h1 {
        font-size: 4rem;
    }
    @media only screen and (max-width: 640px) {
        .box {
            flex-direction: column;
        }
        .box > * {
            width: 100%;
        }
        section > h1 {
            padding-bottom: 5rem;
        }
        .logo {
            margin: auto;
            margin-bottom: 5rem;
        }
    }
</style>
