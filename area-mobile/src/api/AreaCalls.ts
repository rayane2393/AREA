import 'react-native-gesture-handler';
import { getValueFor } from "../../context/SecureStore";
import { AreaInterface } from "../interfaces/AreaInterface";

const API_KEY:string = process.env.EXPO_PUBLIC_API_KEY;
const url:string = process.env.EXPO_PUBLIC_DEV_IP;

async function isAlive() {
    try {
        const response:any = await fetch(`${url}/about.json`, {
            method: 'GET',
        });
        const json:any = await response.json();
        return json;
    } catch (error) {
        console.error(error);
        alert("Could not connect to the server. Please try again later");
    }
}

async function getAreas() {
    try {
        const token:string = await getValueFor("token");
        const response:any = await fetch(`${url}/get-area`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'x-api-key': API_KEY,
                'bearer': token,
            },
        });
        const json:JSON = await response.json();
        // return response.status === 500 ? null : json;
        // console.log(json);
        return json;
    } catch (error) {
        console.error(error);
    }
}

async function editArea(id:number, area:AreaInterface) {
    const token:string = await getValueFor("token");
    console.log(`Edit area ${id} to ${area.name} ${area.color} ${area.action} ${area.reaction} ${area.toggled} ${area.action_name} ${area.reaction_name} ${area.city}`);
    try {
        const response = await fetch(`${url}/modify-area`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'x-api-key': API_KEY,
                'bearer': token,
            },
            body: JSON.stringify({
                id: id,
                name: area.name,
                color: area.color,
                action: area.action,
                reaction: area.reaction,
                toggled: area.toggled,
                action_name: area.action_name,
                reaction_name: area.reaction_name,
                ...(area.summoner_name && { summoner_name: area.summoner_name }),
                ...(area.city && { city: area.city }),
                ...(area.artist_id && { artist_id: area.artist_id }),
            })
        });
        const json:JSON = await response.json();
        return json;
    } catch (error) {
        console.error(error);
    }
}

async function toggleArea(id:number, toggled:boolean) {
    const token:string = await getValueFor("token");
    const areas:any = await getAreas();
    // Log the Area with the id passed in the function
    // Find the area with the correct id
    // console.log(areas);
    const area:any = areas.areas.find((area:any) => area.id === id);
    console.log(`Toggle area ${id} to ${toggled}`);

    try {
        const response:any = await fetch(`${url}/modify-area`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'x-api-key': API_KEY,
                'bearer': token,
            },
            body: JSON.stringify({
                id: id,
                name: area.name,
                color: area.color,
                action: area.action,
                reaction: area.reaction,
                toggled: toggled,
                action_name: area.action_name,
                reaction_name: area.reaction_name,
                ...(area.summoner_name && { summoner_name: area.summoner_name }),
                ...(area.city && { city: area.city }),
                ...(area.artist_id && { artist_id: area.artist_id }),
            })
        });
        const json:JSON = await response.json();
        return json;
    } catch (error) {
        console.error(error);
    }
}

async function postArea(area:AreaInterface) {
    const token:string = await getValueFor("token");
    console.log(`Creating Area Name:${area.name}
    Color:${area.color}
    Action Service:${area.action}
    Reaction Service:${area.reaction}
    Action:${area.action_name}
    Reaction:${area.reaction_name}
    Toggled:${area.toggled}
    Summoner Name:${area.summoner_name}
    City:${area.city}
    Artist ID:${area.artist_id}
    `);
    try {
        const response = await fetch(`${url}/create-area`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'x-api-key': API_KEY,
                'bearer': token,
            },
            body: JSON.stringify({
                name: area.name,
                color: area.color,
                action: area.action,
                reaction: area.reaction,
                toggled: area.toggled,
                action_name: area.action_name,
                reaction_name: area.reaction_name,
                ...(area.summoner_name && { summoner_name: area.summoner_name }),
                ...(area.city && { city: area.city }),
                ...(area.artist_id && { artist_id: area.artist_id }),
            })
        });
        const json:any = await response.json();
        alert(json.message);
        return json;
    } catch (error) {
        console.error(error);
    }
}

async function deleteArea(id:number, action:string) {
    const token:string = await getValueFor("token");
    try {
        const response = await fetch(`${url}/delete-area`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'x-api-key': API_KEY,
                'bearer': token,
            },
            body: JSON.stringify({
                area_id: id,
                action: action
            })
        });
        const json:any = await response.json();
        console.log(json);
        alert(json.message);
        return json;
    } catch (error) {
        console.error(error);
    }
}

async function getActionServices() {
    try {
        const response:any = await fetch(`${url}/about.json`, {
            method: 'GET',
        });
        const json:any = await response.json();

        const servicesWithActions = json.server.services.filter((service:any) => service.actions !== null);
        return servicesWithActions
    } catch (error) {
        console.error(error);
    }
}

async function getReactionServices() {
    try {
        const response:any = await fetch(`${url}/about.json`, {
            method: 'GET',
        });
        const json:any = await response.json();
        const servicesWithReactions = json.server.services.filter((service:any) => service.reactions !== null);
        return servicesWithReactions
    } catch (error) {
        console.error(error);
    }
}

async function getActions(provider:string) {
    try {
        const response:any = await fetch(`${url}/about.json`, {
            method: 'GET',
        });
        const json:any = await response.json();
        const actions = json.server.services
            .filter((service:any) => service.name === provider)
            .flatMap((service:any) => service.actions);
        return { actions };
    } catch (error) {
        console.error(error);
    }
}

async function getActionArgument(provider:string, actionName:string) {
    try {
        const response:any = await fetch(`${url}/about.json`, {
            method: 'GET',
        });
        const json: any = await response.json();
        const actions:any = json.server.services
            .filter((service:any) => service.name === provider)
            .flatMap((service:any) => service.actions);
        const actionArguments:any = actions
            .filter((action:any) => action.name === actionName)
            .flatMap((action:any) => action.arguments);
        console.log(`Action Arguments : ${actionArguments}`);
        return actionArguments[0] ? actionArguments : [];
    } catch (error) {
        console.error(error);
    }
}

async function getReactions(provider:string) {
    try {
        const response:any = await fetch(`${url}/about.json`, {
            method: 'GET',
        });
        const json:any = await response.json();
        const reactions = json.server.services
            .filter((service:any) => service.name === provider)
            .flatMap((service:any) => service.reactions);
        return { reactions };
    } catch (error) {
        console.error(error);
    }
}

async function getReactionArgument(provider:string, reactionName:string) {
    console.log(`getReactionArgument(${provider}, ${reactionName})`);
    try {
        const response:any = await fetch(`${url}/about.json`, {
            method: 'GET',
        });
        const json: any = await response.json();
        const reactions:any = json.server.services
            .filter((service:any) => service.name === provider)
            .flatMap((service:any) => service.reactions);
        const reactionArguments:any = reactions
            .filter((reaction:any) => reaction.name === reactionName)
            .flatMap((reaction:any) => reaction.arguments);
        console.log(`Reaction Arguments : ${reactionArguments}`);
        return reactionArguments[0] ? reactionArguments : [];
    } catch (error) {
        console.error(error);
    }
}

export {
    isAlive,
    getAreas,
    postArea,
    toggleArea,
    editArea,
    deleteArea,

    getActionServices,
    getReactionServices,

    getActions,
    getReactions,

    getActionArgument,
    getReactionArgument,
};
