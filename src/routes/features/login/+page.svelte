<script>
  import { login } from "../../../script";

  let email = $state("");
  let password = $state("");
  const onclick = async () => {
    const data = { email, password };
    await login(data);
  };
</script>

<h1>Login</h1>
<label>email</label>
<input type="text" placeholder="hari@laddu.cc" bind:value={email} />
<label>password</label>
<input type="password" bind:value={password} />
<button {onclick}>login</button>
