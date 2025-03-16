<script>
  import { signup } from "../../../script";

  let email = $state("");
  let password = $state("");
  let name = $state("");
  const onclick = async () => {
    const data = { email, password, name };
    await signup(data);
  };
</script>

<h1>Register</h1>
<label>name</label>
<input type="text" placeholder="hari" bind:value={name} />
<label>email</label>
<input type="text" placeholder="hari@laddu.cc" bind:value={email} />
<label>password</label>
<input type="password" bind:value={password} />
<button {onclick}>register</button>
