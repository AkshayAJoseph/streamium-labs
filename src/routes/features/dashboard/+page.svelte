<script>
  import { goto } from "$app/navigation";
  import { checkUser } from "../../../script";

  const doCheck = async () => {
    const data = await checkUser();
    if (!data) {
      goto("/features", { replaceState: true });
      return;
    }
    console.log(data);
  };

  doCheck();
</script>

<h1>Dashboard</h1>
