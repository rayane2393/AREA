const API_KEY:string = process.env.EXPO_PUBLIC_API_KEY;
const url:string = process.env.EXPO_PUBLIC_DEV_IP;
import { getValueFor } from "../../context/SecureStore";

async function postDiscordOAuth(code:string) {
  try {
    const response:any = await fetch(`${url}/login/discord`, {
      method: 'POST',
      headers: {
          'Content-Type': 'application/json',
          'x-api-key': API_KEY,
      },
      body: JSON.stringify({
        "code": code,
      })
    });
    const json:any = await response.json();
    console.log(json);
    return json.token ? json.token : "";
  } catch (error) {
    console.error(error);
    alert('There was an error fetching your data, please try again later or use another sign in method.');
  }
}

async function postDiscordLogged(code:string) {
  const token:string = await getValueFor("token");
  try {
    const response:any = await fetch(`${url}/login/discord`, {
      method: 'POST',
      headers: {
          'Content-Type': 'application/json',
          'x-api-key': API_KEY,
          'bearer': token,
      },
      body: JSON.stringify({
        "code": code,
      })
    });
    const json:any = await response.json();
    console.log(json);
    if (json.message === "User already logged with discord") {
      alert("This discord account is already linked to another account.");
      return false;
    }
    return json.message === "Login successful" ? true : false;
  } catch (error) {
    console.error(error);
    alert('There was an error fetching your data, please try again later or use another sign in method.');
  }
}

async function postGithubOAuth(code:string) {
  try {
    const response:any = await fetch(`${url}/login/github`, {
      method: 'POST',
      headers: {
          'Content-Type': 'application/json',
          'x-api-key': API_KEY,
      },
      body: JSON.stringify({
        "code": code,
      })
    });
    const json:any = await response.json();
    console.log(json);
    return json.token ? json.token : "";
  } catch (error) {
    console.error(error);
    alert('There was an error fetching your data, please try again later or use another sign in method.');
  }
}

async function postGithubLogged(code:string) {
  const token:string = await getValueFor("token");
  try {
    const response:any = await fetch(`${url}/login/github`, {
      method: 'POST',
      headers: {
          'Content-Type': 'application/json',
          'x-api-key': API_KEY,
          'bearer': token,
      },
      body: JSON.stringify({
        "code": code,
      })
    });
    const json:any = await response.json();
    console.log(json);
    if (json.message === "User already logged with github") {
      alert("This github account is already linked to another account.");
      return false;
    }
    return json.message === "Login successful" ? true : false;
  } catch (error) {
    console.error(error);
    alert('There was an error fetching your data, please try again later or use another sign in method.');
  }
}

export {postDiscordOAuth, postDiscordLogged,
        postGithubOAuth, postGithubLogged}
