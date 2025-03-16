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
        <div class="blogs__container">
            <h1>Read our Blogs</h1>
            {#if loading}
                <p>Loading...</p>
            {:else if blogs.length > 0}
                {#each Array(Math.ceil(blogs.length / 3)) as _, rowIndex}
                    <div class="blogs">
                        {#each Array(3) as _, colIndex}
                            {@const blogIndex = rowIndex * 3 + colIndex}
                            {#if blogIndex < blogs.length}
                                {@const blog = blogs[blogIndex]}
                                <a href={`/blogs/blog?value=${blogIndex}`}>
                                    <div class="blog">
                                        <div class="blog__banner">
                                            <img
                                                src={blog.banner}
                                                alt="banner"
                                            />
                                        </div>
                                        <div class="blog__text">
                                            <div class="blog__info">
                                                <h2>{blog.title}</h2>
                                            </div>
                                            <div class="blog__details">
                                                <p>{blog.postedByName}</p>
                                                <p>{blog.date}</p>
                                            </div>
                                        </div>
                                    </div>
                                </a>
                            {/if}
                        {/each}
                    </div>
                {/each}
            {:else}
                <p>No blogs available.</p>
            {/if}
        </div>
    </div>
</main>

<style>
    .blogs__container {
        margin-top: 5vh;
    }

    .blogs__container h1 {
        text-align: center;
        font-size: 5rem;
        padding: 5rem;
    }
    .blogs {
        display: flex;
        padding: 2rem;
        justify-content: space-between;
    }
    .blogs > * {
        width: 30%;
    }
    .blog__text {
        display: flex;
        justify-content: space-between;
        padding: 2rem;
        background: #070707;
        border-radius: 1rem;
    }
    .blog__text > * {
        width: 45%;
    }
    .blog__banner img {
        width: 100%;
        height: 40vh;
        object-fit: cover;
        border-top-left-radius: 1rem;
        border-top-right-radius: 1rem;
    }
    @media only screen and (max-width: 640px) {
        .blogs__container h1 {
            font-size: 2rem;
        }
        .blogs {
            flex-direction: column;
        }
        .blogs > * {
            width: 100%;
            margin-bottom: 2rem;
        }
    }
</style>
