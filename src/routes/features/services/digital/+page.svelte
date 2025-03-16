
<script>
    //done
  import { getService } from "../../../../script";
  let data;

  const dogetService = async () => {
    data = await getService();
    data = data.filter((item) => item.type === "Digital");
    console.log(data);
  };

  dogetService();
</script>

<h1>Digital Services</h1>

<ul>
  {#each digitalItems as item}
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
