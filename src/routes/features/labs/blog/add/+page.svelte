<script>
  import { goto } from "$app/navigation";
  import { addBlog, checkUser } from "../../../../../script";

  let data;

  // Form fields
  let title = "";
  let content = "";
  let banner = "";
  let date = new Date().toISOString().split("T")[0]; // Default to today's date
  let postedName;
  let postedId;

  // Function to handle form submission
  async function handleSubmit() {
    // Create a new blog post object
    const newBlog = {
      title,
      content,
      banner,
      date,
      postedName,
      postedId,
    };
    console.log(newBlog);
    await addBlog(newBlog);
    // Redirect to blog page after submission
  }
  const doCheck = async () => {
    data = await checkUser();
    if (!data) {
      goto("/features", { replaceState: true });
      return;
    }
    console.log(data);
    postedName = data.name;
    postedId = data.id;
  };

  doCheck();
</script>

<h1>Add a New Blog Post</h1>

<form on:submit|preventDefault={handleSubmit}>
  <label>
    Title:
    <input type="text" bind:value={title} required />
  </label>

  <label>
    Content:
    <textarea bind:value={content} required></textarea>
  </label>

  <label>
    Banner URL:
    <input type="url" bind:value={banner} required />
  </label>

  <button type="submit">Submit</button>
</form>

<a href="/blog">
  <button type="button">Cancel</button>
</a>
