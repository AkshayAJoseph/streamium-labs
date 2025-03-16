<script>
  import { onMount } from "svelte";
  import { getJob } from "../../../script";

  let jobs = [];

  // Fetch jobs on page load
  async function dogetJob() {
    try {
      jobs = await getJob();
    } catch (error) {
      console.error("Error fetching jobs:", error);
    }
  }

  // Run function when the component is mounted
  onMount(dogetJob);
</script>

<h1>Job Listings</h1>

{#if jobs.length > 0}
  <ul>
    {#each jobs as job}
      <li>
        <h2>{job.title}</h2>
        <p><strong>Role:</strong> {job.role}</p>
        <p><strong>Description:</strong> {job.description}</p>
        <p><strong>Requirements:</strong> {job.requirments}</p>
        <p><strong>Posted On:</strong> {job.postedOn}</p>
      </li>
    {/each}
  </ul>
{:else}
  <p>Loading jobs...</p>
{/if}
