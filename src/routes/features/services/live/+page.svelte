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
      type: "Live Streaming",
      category: "Virtual Events",
      status: "live",
      price: 200,
    },
    {
      id: 5,
      type: "Live Streaming",
      category: "Hybrid Events",
      status: "scheduled",
      price: 300,
    },
  ];

  // Filter only Live Streaming category
  let liveItems = data.filter((item) => item.type === "Live Streaming");
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
