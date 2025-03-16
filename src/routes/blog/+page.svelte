<script>
    import Nav from "$lib/components/Nav.svelte";
    import { onMount } from "svelte";
    import { getBlog } from "../../script";

    let blogs = [];
    let loading = true;

    async function fetchBlogs() {
        blogs = await getBlog();
        loading = false;
    }

    onMount(fetchBlogs);
</script>

<main>
    <Nav />
    <div class="box">
        <div class="blogs">
            {#if loading}
                <p>Loading...</p>
            {:else if blogs.length > 0}
                <ul>
                    {#each blogs as blog}
                        <div class="blog">
                            <div class="blog__banner">
                                <img src={blog.banner} alt="banner" />
                            </div>
                            <div class="blog__info">
                                <h2>{blog.title}</h2>
                                <p>{blog.content}</p>
                            </div>
                            <div class="blog__details">
                                <p>{blog.postedByName}</p>
                                <p>{blog.date}</p>
                            </div>
                        </div>
                    {/each}
                </ul>
            {:else}
                <p>No blogs available.</p>
            {/if}
        </div>
    </div>
</main>

<style>
    .blogs {
        width: calc(100% - 4rem);
        margin-left: 2rem;
        margin-top: 10rem;
    }
    .blog {
        margin-bottom: 2rem;
        background: #070707;
        border-radius: 1rem;
    }
    .blog__banner img {
        width: 100%;
        height: 40vh;
        object-fit: cover;
        border-top-left-radius: 1rem;
        border-top-right-radius: 1rem;
    }
    .blog__details,
    .blog__info {
        padding: 1rem;
    }
    .blog > * {
        margin-bottom: 2rem;
    }
</style>
