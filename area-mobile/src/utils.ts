function nameToIonicon(serviceName: string): string {
    if (serviceName === 'mail' || serviceName === 'email') {
        return 'mail-outline';
    }
    switch (serviceName) {
        case "discord":
            return "chatbox-outline";
        case "facebook":
            return "logo-facebook";
        case "spotify":
            return "musical-notes-outline";
        case "github":
            return "logo-github";
        case "instagram":
            return "logo-instagram";
        case "lol":
            return "game-controller-outline";
        case "meteo":
            return "cloud-outline";
        case "steam":
            return "logo-steam";
        default:
            return "cube-outline";
    }
}

function cleanName(name:string):string {
    // replace the first letter by a capital letter
    return name.charAt(0).toUpperCase() + name.slice(1);
}

export { nameToIonicon, cleanName };