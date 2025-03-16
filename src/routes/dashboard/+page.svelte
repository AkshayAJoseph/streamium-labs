<script>
    import Sidebar from "$lib/components/Sidebar.svelte";
</script>

<main>
    <div class="box">
        <Sidebar />
        <section>
            <h1>Dashboard</h1>
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
</style>
