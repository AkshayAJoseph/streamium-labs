<script>
    import Sidebar from "$lib/components/Sidebar.svelte";
    import { onMount } from "svelte";

    const typeCategoryMap = {
        "Live Streaming": [
            { category: "Virtual Events", price: 200 },
            { category: "Hybrid Events", price: 300 },
        ],
        "Media Production": [
            { category: "Videography", price: 150 },
            { category: "Photography", price: 100 },
            { category: "Editing", price: 120 },
        ],
        "Digital Marketing": [
            { category: "Social Media Management", price: 250 },
            { category: "SEO", price: 180 },
            { category: "Ad Campaigns", price: 220 },
        ],
        "Event Management": [
            { category: "Event Planning", price: 400 },
            { category: "Execution", price: 350 },
        ],
    };

    let type = ""; // Selected Type
    let categories = []; // Categories based on Type
    let category = ""; // Selected Category
    let price = 0; // Fixed price based on Category

    function updateCategories() {
        categories = typeCategoryMap[type] || [];
        category = categories.length ? categories[0].category : "";
        updatePrice(); // Auto-select first category
    }

    function updatePrice() {
        let selected = categories.find((c) => c.category === category);
        price = selected ? selected.price : 0;
    }

    onMount(() => {
        updateCategories();
    });
</script>

<main>
    <div class="box">
        <Sidebar />
        <section>
            <h1>Create an Event</h1>
            <article class="card">
                <div class="form">
                    <div class="form__row">
                        <label>Type</label>
                        <select
                            id="type"
                            bind:value={type}
                            on:change={updateCategories}
                        >
                            <option value="" disabled selected
                                >Select Type</option
                            >
                            {#each Object.keys(typeCategoryMap) as t}
                                <option value={t}>{t}</option>
                            {/each}
                        </select>
                    </div>
                    <div class="form__row">
                        <label for="category">Category</label>
                        <select
                            id="category"
                            bind:value={category}
                            on:change={updatePrice}
                            disabled={!categories.length}
                        >
                            {#each categories as c}
                                <option value={c.category}>{c.category}</option>
                            {/each}
                        </select>
                    </div>
                    <div class="form__row">
                        <label for="price">Price</label>
                        <input
                            type="text"
                            id="price"
                            bind:value={price}
                            readonly
                        />
                    </div>
                    <div class="form__row">
                        <input type="submit" value="Submit" />
                    </div>
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

    .form label {
        font-size: 0.5rem;
        text-transform: uppercase;
    }

    .form select {
        margin-top: 0.5rem;
        margin-bottom: 2rem;
        display: block;
        width: 100%;
        padding: 1rem 2rem;
        background: #111;
        -webkit-appearance: none;
        color: white;
        border: 0;
        border-radius: 10rem;
        font-family: "Inter";
        font-size: 1rem;
    }
    .form__row {
        margin-bottom: 1rem;
    }
    .form input {
        display: block;
        padding: 1rem 2rem;
        border-radius: 10rem;
        width: 100%;
        background: #111;
        border: 0;
        color: white;
        margin-top: 1rem;
        font-family: "Inter";
        font-weight: 600;
        font-size: 1rem;
    }
    .form input[type="submit"] {
        margin-top: 5rem;
        background: white;
        color: black;
    }
    .card {
        width: 50%;
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
