<script>
  import { onMount } from "svelte";

  // Mapping of Type -> Categories with Prices
  const typeCategoryMap = {
    "Live Streaming": [
      { category: "Virtual Events", price: 200 },
      { category: "Hybrid Events", price: 300 },
    ],
    "Media Production": [
      { category: "Videography", price: 150 },
      { category: "Photography", price: 100 },
      { category: "Editing", price: 120 },
    ],
    "Digital Marketing": [
      { category: "Social Media Management", price: 250 },
      { category: "SEO", price: 180 },
      { category: "Ad Campaigns", price: 220 },
    ],
    "Event Management": [
      { category: "Event Planning", price: 400 },
      { category: "Execution", price: 350 },
    ],
  };

  let type = ""; // Selected Type
  let categories = []; // Categories based on Type
  let category = ""; // Selected Category
  let price = 0; // Fixed price based on Category

  // Update categories when Type changes
  function updateCategories() {
    categories = typeCategoryMap[type] || [];
    category = categories.length ? categories[0].category : "";
    updatePrice(); // Auto-select first category
  }

  // Update price when Category changes
  function updatePrice() {
    let selected = categories.find((c) => c.category === category);
    price = selected ? selected.price : 0;
  }

  // Ensure categories are updated if type changes on load
  onMount(() => {
    updateCategories();
  });
</script>

<h1>Add a New Item</h1>

<form>
  <!-- Type Dropdown -->
  <label for="type">Type:</label>
  <select id="type" bind:value={type} on:change={updateCategories}>
    <option value="" disabled selected>Select Type</option>
    {#each Object.keys(typeCategoryMap) as t}
      <option value={t}>{t}</option>
    {/each}
  </select>

  <br />

  <!-- Category Dropdown -->
  <label for="category">Category:</label>
  <select
    id="category"
    bind:value={category}
    on:change={updatePrice}
    disabled={!categories.length}
  >
    {#each categories as c}
      <option value={c.category}>{c.category}</option>
    {/each}
  </select>

  <br />

  <!-- Fixed Price (Non-editable) -->
  <label for="price">Price:</label>
  <input type="text" id="price" bind:value={price} readonly />

  <br />

  <button type="submit">Submit</button>
</form>
