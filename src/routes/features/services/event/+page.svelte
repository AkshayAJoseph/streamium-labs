<script>
    //done
  import { getService } from "../../../../script";
  let data;

  const dogetService = async () => {
    data = await getService();
    data = data.filter((item) => item.type === "Event");
    console.log(data);
  };

  dogetService();
</script>

<h1>Event Services</h1>

<ul>
  {#if data}
    {#each data as item}
      <li>
        <strong>Category:</strong>
        {item.category} |
        <strong>Status:</strong>
        {item.status} |
        <strong>Price:</strong> ${item.price}
      </li>
    {/each}
  {:else}
    <p>loading...</p>
  {/if}
</ul>

<!-- Add Button -->
<a href="/features/services/add">
  <button>add</button>
</a>
