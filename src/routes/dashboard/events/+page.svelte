<script>
    import Sidebar from "$lib/components/Sidebar.svelte";
    import { getService } from "../../../script";
    let data;

    const dogetService = async () => {
        data = await getService();
        data = data.filter((item) => item.type === "Event Management");
        console.log(data);
    };

    dogetService();
</script>

<main>
    <div class="box">
        <Sidebar selected="events" />
        <section>
            <h1>Events</h1>
            <a href="/dashboard/create" class="btn--create">Create New</a>
            <article class="card">
                <h5>Your Event History</h5>
                <div class="card__table">
                    <div class="card__row">
                        <div class="card__column">
                            <h5>Title</h5>
                        </div>
                        <div class="card__column">
                            <h5>Type</h5>
                        </div>
                        <div class="card__column">
                            <h5>Status</h5>
                        </div>
                        <div class="card__column">
                            <h5>Price</h5>
                        </div>
                    </div>
                    {#if data}
                        {#each data as item}
                            <div class="card__row">
                                <div class="card__column">
                                    <p>{item.title}</p>
                                </div>
                                <div class="card__column">
                                    <p>{item.type}</p>
                                </div>
                                <div class="card__column">
                                    <p>{item.status}</p>
                                </div>
                                <div class="card__column">
                                    <p>{item.price}</p>
                                </div>
                            </div>
                        {/each}
                    {:else}
                        <p>loading...</p>
                    {/if}
                </div>
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

    .card__table {
        display: flex;
        flex-direction: column;
    }

    .card__row {
        display: flex;
        justify-content: space-between;
    }

    .card__column h5 {
        font-size: 1rem;
    }

    .btn--create {
        padding: 1rem 2rem;
        background: white;
        width: fit-content;
        color: #000;
        font-weight: 500;
        margin: 2rem 0;
        border-radius: 10rem;
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
    }
</style>
