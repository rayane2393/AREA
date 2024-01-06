import 'react-native-gesture-handler';
import { getValueFor } from "../../context/SecureStore";

const API_KEY:string = process.env.EXPO_PUBLIC_API_KEY;
const url:string = process.env.EXPO_PUBLIC_DEV_IP;

async function userInfo() {
    const token:string = await getValueFor("token");
    try {
        const response:any = await fetch(`${url}/user`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'x-api-key': API_KEY,
                'bearer': token,
            },
        });
        const json:any = await response.json();
        // console.log(json.user);
        console.log("--------------------");
        console.log(`USERNAME: ${json.user.username}`);
        console.log(`DISCORD ID: ${json.user.discord_id}`);
        console.log(`GITHUB ID: ${json.user.github_id}`);
        return json;
    } catch (error) {
        console.error(error);
    }
}

async function postLogin(login:string, password:string) {
    try {
        const response:any = await fetch(`${url}/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'x-api-key': API_KEY,
            },
            body: JSON.stringify({
                username: login,
                password: password
            })
        });
        const json:any = await response.json();
        console.log(json);
        return json.token;
    } catch (error) {
        console.error(error);
        return "";
    }
}

async function postAuth(username:string, password:string, email:string) {
    try {
        const response:any = await fetch(`${url}/register`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'x-api-key': API_KEY,
            },
            body: JSON.stringify({
                username: username,
                email: email,
                password: password
            })
        });
        const json:any = await response.json();
        json.status = response.status;
        return json;
    } catch (error) {
        console.error(error);
    }
}

export {
    userInfo,
    postLogin,
    postAuth,
};
