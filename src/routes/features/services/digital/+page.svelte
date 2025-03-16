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
      type: "Digital",
      category: "Social Media Management",
      status: "available",
      price: 250,
    },
    {
      id: 4,
      type: "Event",
      category: "Planning",
      status: "ongoing",
      price: 400,
    },
    {
      id: 5,
      type: "Digital",
      category: "Ad Campaigns",
      status: "available",
      price: 220,
    },
  ];

  // Filter only Digital category
  let digitalItems = data.filter((item) => item.type === "Digital");
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
