<script>
  import { fetchYesterdayMatches, getLive } from "../../../script";
  import { onMount } from "svelte";
  import { fetchLiveMatches } from "../../../script";

  let lives = [];

  // Fetch jobs on page load
  async function dogetLive() {
    try {
      lives = await getLive();
    } catch (error) {
      console.error("Error fetching live:", error);
    }

    
    await fetchLiveMatches();
    await fetchYesterdayMatches();
  }

  // Run function when the component is mounted
  onMount(dogetLive);
</script>

<h1>LiveStream</h1>

{#if lives}
  <ul>
    {#each lives as live}
      <li>
        <h2>{live.title}</h2>
        <p><strong>Duration:</strong> {live.duration}</p>
        <p><strong>Description:</strong> {live.description}</p>
        <p><strong>Requirements:</strong> {live.requirments}</p>
      </li>
    {/each}
  </ul>
{:else}
  <p>Loading lives...</p>
{/if}
