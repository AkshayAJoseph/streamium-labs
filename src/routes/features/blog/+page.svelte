<script>
    import { onMount } from "svelte";
    import { getBlog } from "../../../script";

    let blogs = [];
    let loading = true;

    async function fetchBlogs() {
        blogs = await getBlog();
        loading = false;
    }

    onMount(fetchBlogs);
</script>

<h1>Blog Page</h1>

{#if loading}
    <p>Loading...</p>
{:else if blogs.length > 0}
    <ul>
        {#each blogs as blog}
            <li>
                <h2>{blog.title}</h2>
                <img src={blog.banner} alt={blog.title} width="400" />
                <p>
                    <strong>Posted By:</strong>
                    {blog.postedByName} | <strong>Date:</strong>
                    {blog.date}
                </p>
                <p>{blog.content}</p>
            </li>
        {/each}
    </ul>
{:else}
    <p>No blogs available.</p>
{/if}

<a href="/features/labs/blog/add">
    <button>Add</button>
</a>
