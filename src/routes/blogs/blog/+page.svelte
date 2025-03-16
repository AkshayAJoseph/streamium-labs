<script>
    import { onMount } from "svelte";
    import { getBlog } from "../../../script";

    let blogs = [];
    let loading = true;
    let data;

    async function fetchBlogs() {
        blogs = await getBlog();
        loading = false;
    }

    onMount(() => {
        const searchValue = window.location.search;
        const urlParams = new URLSearchParams(searchValue);
        data = urlParams.get("value");
        console.log(data);
        fetchBlogs();
    });

    const b = blogs[data];
    console.log(b);
</script>

<main>hey</main>

<style>
</style>
