<script>
  import { fetchYesterdayMatches, getLive } from "../../../script";
  import { onMount } from "svelte";
  import RazorPay from "$lib/components/RazorPay.svelte";
  let matches;

  // Fetch jobs on page load
  async function dogetLive() {
    matches = await fetchYesterdayMatches();
    console.log(matches);
  }

  // Run function when the component is mounted
  onMount(dogetLive);
</script>

<div class="container">
  <h1>Esports Match Results</h1>

  {#each matches as match}
    <img src={match.team1Logo} alt={match.team1} width="100" height="100" />
    <span>{match.team1}</span>
    {match.team1Score} - {match.team2Score}
    <span>{match.team2}</span>
    <img src={match.team2Logo} width="100" height="100" alt={match.team2} />
    <RazorPay />
    <br />
    <br />
  {/each}
</div>
