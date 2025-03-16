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
      goto("/dashboard", { replaceState: true });
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

async function getLive() {
    const response = await fetch(`${baseUrl}/live`, {
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
        console.log(res)
        return res.data;
    }

async function logout() {
    await Storage.remove({ key
        : "token" });
    goto("/login", { replaceState: true
    });
}

async function fetchLiveMatches() {
    const API_KEY = "HcdhRC6J2dtLthJ6LDnM_-SyCMVRnsOi9InrUvnfKdi49xCzO1M"; // Replace with your Pandascore API key
    const url = `https://api.pandascore.co/matches/running?token=${API_KEY}`;

    try {
        const response = await fetch(url);
        const matches = await response.json();

        matches.forEach(match => {
            console.log(`Game: ${match.videogame.name}`);
            console.log(`Teams: ${match.opponents[0]?.opponent.name} vs ${match.opponents[1]?.opponent.name}`);
            console.log(`Score: ${match.results[0]?.score} - ${match.results[1]?.score}`);
            console.log("--------");
        });
    } catch (error) {
        console.error("Error fetching matches:", error);
    }
    
}

async function fetchYesterdayMatches() {
    const API_KEY = "HcdhRC6J2dtLthJ6LDnM_-SyCMVRnsOi9InrUvnfKdi49xCzO1M"; // Replace with your Pandascore API key
   const yesterday = new Date();
    yesterday.setDate(yesterday.getDate() - 1);
    const dateStr = yesterday.toISOString().split("T")[0];
    let data =[]
 

    const url = `https://api.pandascore.co/matches/past?range[begin_at]=${dateStr}T00:00:00Z,${dateStr}T23:59:59Z&token=${API_KEY}`;

        const response = await fetch(url);
        const matches = await response.json();

        matches.forEach(match => {
            const game = match.videogame.name;
        
            const team1 = match.opponents[0]?.opponent.name || "TBD";
            const team1Logo = match.opponents[0]?.opponent.image_url || null;

            const team2 = match.opponents[1]?.opponent.name || "TBD";
            const team2Logo = match.opponents[1]?.opponent.image_url || null;

            const team1Score = match.results[0]?.score || 0;
            const team2Score = match.results[1]?.score || 0

          data.push({game, team1, team1Logo, team2, team2Logo, team1Score, team2Score});
        });
        console.log(data)
        return data
}

async function addService(data) {
  const response = await fetch(`${baseUrl}/service`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
  console.log(response)
  const res = await response.json();
  if (!response.ok) {
    alert(res.message);
    return;
  }
  console.log('added success')

}

  export { login , signup, checkUser, getService, getBlog, addBlog, getJob, logout, getLive, fetchLiveMatches, fetchYesterdayMatches ,addService};