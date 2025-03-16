import { goto } from "$app/navigation";
import { Storage } from "@capacitor/storage";

const baseUrl = "https://api.laddu.cc/api/v1";

async function setToken(token) {
    await Storage.set({
      key: "token",
      value: token,
    });
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
      goto("/features/dashboard", { replaceState: true });
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
      goto("/features/dashboard", { replaceState: true });
    } catch (error) {   
      console.log(error);
    }
  }

  async function checkUser() {
    const { value } = await Storage.get({ key: "token" });
    console.log(value);
    if (!value) {
      goto("/login", { replaceState: true });
      return;
    }
    const response = await fetch(`${baseUrl}/verify`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${value}`,
      },
    });
    const res = await response.json();
    if (!response.ok) {
      alert(res.message);
      await logout();
      goto("/login", { replaceState: true });
      return;
    }
    const id = res.id;
    const response2 = await fetch(`${baseUrl}/users/${id}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });
    const res2 = await response2.json();
    if (!response2.ok) {
      alert(res2.message);
      return;
    }
    console.log('done')
    return res2.data;
  }

    async function getService() {
        const response = await fetch(`${baseUrl}/service`, {
            method: "GET",
            headers: {
              "Content-Type": "application/json",
            },
          });
          const res = await response.json();
          if (!response.ok) {
            alert(res.message);
            return;
          }
            console.log('done')
            return res.data;
    }


    async function getBlog() {
        const response = await fetch(`${baseUrl}/blog`, {
            method: "GET",
            headers: {
              "Content-Type": "application/json",
            },
          });
          const res = await response.json();
          if (!response.ok) {
            alert(res.message);
            return;
          }
            console.log('done')
            return res.data;
    }


async function addBlog(data){
    const response = await fetch(`${baseUrl}/blog`, {
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
            console.log(res.data)
}

async function getJob() {
const response = await fetch(`${baseUrl}/job`, {
    method: "GET",
    headers: {
        "Content-Type": "application/json",
    },
    });
    const res = await response.json();
    if (!response.ok) {
    alert(res.message);
    return;
    }
    console.log('done')
    return res.data;
}

  export { login , signup, checkUser, getService, getBlog, addBlog, getJob};