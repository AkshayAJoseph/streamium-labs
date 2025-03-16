<script>
  // Dummy data
  let data = [
    {
      id: 1,
      type: "Media",
      category: "Videography",
      status: "available",
      price: 150,
    },
    { id: 2, type: "Digital", category: "SEO", status: "active", price: 180 },
    {
      id: 3,
      type: "Event",
      category: "Event Planning",
      status: "ongoing",
      price: 400,
    },
    {
      id: 4,
      type: "Event",
      category: "Execution",
      status: "scheduled",
      price: 350,
    },
    {
      id: 5,
      type: "Digital",
      category: "Ad Campaigns",
      status: "available",
      price: 220,
    },
  ];

  // Filter only Event category
  let eventItems = data.filter((item) => item.type === "Event");
</script>

<h1>Event Services</h1>

<ul>
  {#each eventItems as item}
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