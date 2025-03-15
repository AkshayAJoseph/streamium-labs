import { goto } from "$app/navigation";

const baseUrl = "https://api.laddu.cc/api/v1";

async function setToken(token) {
    localStorage.setItem("token", token);
  }

  async function login(data) {
    try {
      console.log(data);
      const response = await fetch(`${baseUrl}/login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });
      const res = await response.json();
      if (!response.ok) {
        alert(res.message);
        return;
      }
      await setToken(res.token);
      goto("/features", { replaceState: true });
    } catch (error) {
      console.log(error);
    }
  }

  async function signup(data) {
    try {
      console.log(data);
      const response = await fetch(`${baseUrl}/register`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });
      const res = await response.json();
      if (!response.ok) {
        alert(res.message);
        return;
      }
      await setToken(res.token);
      goto("/features", { replaceState: true });
    } catch (error) {   
      console.log(error);
    }
  }

  export { login , signup};