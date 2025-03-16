<script>
  // done
  let data;

  const dogetService = async () => {
    data = await getService();
    data = data.filter((item) => item.type === "Live");
    console.log(data);
  };

  dogetService();
</script>

<h1>Live Streaming Services</h1>

<ul>
  {#each liveItems as item}
    <li>
      <strong>Category:</strong>
      {item.category} |
      <strong>Status:</strong>
      {item.status} |
      <strong>Price:</strong> ${item.price}
    </li>
  {/each}
</ul>

<!-- Add Button -->
<a href="/features/services/add">
  <button>add</button>
</a>
