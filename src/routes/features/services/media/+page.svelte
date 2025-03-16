<script>
  //done
  import { getService } from "../../../../script";
  let data;

  const dogetService = async () => {
    data = await getService();
    data = data.filter((item) => item.type === "Media");
    console.log(data);
  };

  dogetService();
</script>

<h1>Past Media projects</h1>

<ul>
  {#if data}
    {#each data as item}
      <li>
        <strong>Type:</strong>
        {item.type} |
        <strong>Status:</strong>
        {item.status} |
        <strong>Price:</strong> ${item.price}
      </li>
    {/each}
  {:else}
    <p>loading...</p>
  {/if}
</ul>
<a href="/features/services/add">
  <button>add</button>
</a>
