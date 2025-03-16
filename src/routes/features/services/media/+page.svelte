<script>
  // Dummy data
  let data = [
    { id: 1, type: "Movie", category: "media", status: "available", price: 10 },
    {
      id: 2,
      type: "Book",
      category: "education",
      status: "out of stock",
      price: 15,
    },
    { id: 3, type: "Music", category: "media", status: "available", price: 8 },
    {
      id: 4,
      type: "Game",
      category: "entertainment",
      status: "available",
      price: 20,
    },
  ];

  // Filter only media category
  let mediaItems = data.filter((item) => item.category === "media");
</script>

<h1>Past Media projects</h1>

<ul>
  {#each mediaItems as item}
    <li>
      <strong>Type:</strong>
      {item.type} |
      <strong>Status:</strong>
      {item.status} |
      <strong>Price:</strong> ${item.price}
    </li>
  {/each}
</ul>
<a href="/features/services/add">
  <button>add</button>
</a>
