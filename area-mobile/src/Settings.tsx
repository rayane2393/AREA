import React, { useContext, useEffect } from "react";
import { View, Text, Image, Alert, StyleSheet, Vibration } from "react-native";
import { TouchableOpacity } from "react-native-gesture-handler";
import { AuthContext } from "../context/AuthContext";
import { Ionicons } from "@expo/vector-icons";
import { WebView } from 'react-native-webview';
import { postGithubLogged, postDiscordLogged } from "./api/AuthCalls";
import { userInfo } from "./api/ApiCalls";

const Settings = ({ navigation }) => {
    const { logout } = useContext(AuthContext);
    const [authGithub, setAuthGithub] = React.useState(false);
    const [authGoogle, setAuthGoogle] = React.useState(false);
    const [authDiscord , setAuthDiscord] = React.useState(false);
    const [isGithub, setIsGithub] = React.useState(false);
    const [isGoogle, setIsGoogle] = React.useState(false);
    const [isDiscord, setIsDiscord] = React.useState(false);
    const [userName, setUserName] = React.useState("");

    const showHelp = ():void => {
        Alert.alert(
            "Help",
            "Fuck you.",
            [ { text: "OK" } ],
            { cancelable: true }
        );
    }

    const confirmLogout = ():void => {
        Alert.alert(
            "Logout",
            "Are you sure you want to log out ?",
            [
                { text: "Cancel", style: "cancel",},
                { text: "Logout", onPress: () => logout() },
            ],
            { cancelable: true }
        );
    }

    async function getUserInfo() {
        const info:any = await userInfo();
        if (info.user.github_id) {
            setIsGithub(true);
        }
        if (info.user.discord_id) {
            setIsDiscord(true);
        }
        if (info.user.username) {
            setUserName(info.user.username);
        }
    }

    function handleWebViewWarning(message:string) {
        console.log('WebView Warning:', message);
        const code:string = message.split('code=')[1];
        console.log(code);
        setAuthGithub(false);
        postGithubLogged(code).then((success) => {
            if (success) {
                setIsGithub(true);
            } else {
                alert("GitHub Authentication failed, please use a different sign in method for now.");
            }
        });
    }

    useEffect(() => {
        const originalConsoleWarn = console.warn;
        console.warn = function (message:string) {
            if (message.includes("Can't open url:")) {
                handleWebViewWarning(message);
                return;
            }
            originalConsoleWarn.apply(console, arguments);
        };
        getUserInfo();
    }, []);

    const githubAuth = () => {
        const clientID:string = "41cb0cc52810a5bc428a";
        const redirectUrl:string = "area://github";
        const githubScope:string[] = [
            'repo',
            'repo:status',
            'repo_deployment',
            'public_repo',
            'admin:repo_hook',
            'write:repo_hook',
            'admin:org',
            'gist',
            'notifications',
            'user',
            'delete_repo',
            'write:discussion',
            'write:packages',
            'read:packages',
            'delete:packages',
            'admin:gpg_key',
            'admin:org_hook',
            'admin:repo',
            'admin:enterprise',
            'read:user',
            'read:discussion',
            'read:enterprise',
            'read:org',
            'workflow',
        ]

        return (
            <WebView
                source={{
                    uri: `https://github.com/login/oauth/authorize?client_id=${clientID}&redirect_uri=${redirectUrl}&scope=${githubScope.join(',')}`
                }}
                onMessage={(event) => {
                    const data = event.nativeEvent.data;
                    console.log(`Github Data : ${data}`);
                    setAuthGithub(false);
                }}
                onError={(event) => {
                    console.log(`Github Event : ${event}`);
                    setAuthGithub(false);
                }}
            />
        );
    };

    const discordAuth = () => {
        const clientID:string = "1168952746733473902";

        return (
            <WebView
                source={{
                    uri: `https://discord.com/api/oauth2/authorize?client_id=${clientID}&redirect_uri=https%3A%2F%2Fgoogle.com&response_type=code&scope=identify%20email`,
                }}
                onMessage={(event) => {
                    const data:string = event.nativeEvent.data;
                    console.log(`Discord Data : ${data}`);
                    setAuthDiscord(false);
                }}
                onError={(event) => {
                    console.log(`Discord Event : ${event}`);
                    setAuthDiscord(false);
                }}
                onNavigationStateChange={(navState) => {
                    const url:string = navState.url;
                    if (url.includes("google.com") && !url.includes("discord.com")) {
                        const code:string = url.split("code=")[1];
                        console.log(`Discord code : ${code}`);
                        setAuthDiscord(false);
                        postDiscordLogged(code).then((success) => {
                            if (success) {
                                setIsDiscord(true);
                            } else {
                                alert("Discord Authentication failed, please use a different sign in method for now.");
                            }
                        });
                    }
                }}
            />
        );
    };

    const googleAuth = () => {
        // Google auth
    }

    return (
        authGithub ? githubAuth() :
        authDiscord ? discordAuth() :
        <View style={{ flex: 1, justifyContent: "center", alignItems: "center",backgroundColor: "#FFFFFF",}}>
            <Image style={styles.logo} source={require("../assets/Logo_white.png")} />
            <Ionicons
                name="return-down-back-outline"
                size={40} color="#007AFF"
                onPress={() => {
                    navigation.navigate("Home");
                    Vibration.vibrate(50)}}
                style={{
                    position: "absolute",
                    left: '6%',
                    top: '5%'
                }}
            />
            <Text style={styles.title}>Settings</Text>
            <Text style={styles.smallerTitle}>{userName ? `Welcome, ${userName}` : ""}</Text>
            <View style={styles.settingContainer}>
                {/* <TouchableOpacity onPress={showHelp} style={styles.button}>
                    <Text style={styles.buttonText}>Help</Text>
                </TouchableOpacity> */}

                <TouchableOpacity onPress={() => {setAuthGithub(true)}} style={!isGithub ? styles.button : styles.disabledButton} disabled={isGithub}>
                    <Text style={styles.buttonText}>{isGithub ? "Authenticated with Github" : "Authenticate with Github"}</Text>
                </TouchableOpacity>
                {/* <TouchableOpacity onPress={googleAuth} style={!isGoogle ? styles.button : styles.disabledButton} disabled={isGoogle}>
                    <Text style={styles.buttonText}>{isGoogle ? "Authenticated with Google" : "Authenticate with Google"}</Text>
                </TouchableOpacity> */}
                <TouchableOpacity onPress={() => {setAuthDiscord(true)}} style={!isDiscord ? styles.button : styles.disabledButton} disabled={isDiscord}>
                    <Text style={styles.buttonText}>{isDiscord ? "Authenticated with Discord" : "Authenticate with Discord"}</Text>
                </TouchableOpacity>

                <TouchableOpacity onPress={confirmLogout} style={styles.logout}>
                    <Text style={styles.logoutText}>Logout</Text>
                </TouchableOpacity>
            </View>
        </View>
    );

}

const styles = StyleSheet.create({
    container: {
        backgroundColor: "#FFFFFF",
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
    },
    logo: {
        width: 230,
        height: 100,
        bottom: '35%',
    },
    title: {
        color: "#000000",
        fontWeight: "bold",
        fontSize: 40,
    },
    smallerTitle: {
        color: "#000000",
        fontWeight: "bold",
        marginBottom: '40%',
        fontSize: 25,
    },
    settingContainer: {
        // backgroundColor: "#DDDDDD",
        width: '100%',
        alignItems: "center",
    },
    button: {
        backgroundColor: "#78CBF9",
        width: 200,
        alignItems: "center",
        paddingVertical: '1%',
        borderRadius: 5,
        marginBottom: '2%',
    },
    disabledButton: {
        backgroundColor: "#55FF55",
        width: 200,
        alignItems: "center",
        paddingVertical: '1%',
        borderRadius: 5,
        marginBottom: '2%',
    },
    buttonText: {
        color: "#FFFFFF",
        fontWeight: "bold",
        textAlign: "center",
        fontSize: 20,
    },
    logout: {
        backgroundColor: "#FF3333",
        width: 200,
        alignItems: "center",
        paddingVertical: '1%',
        borderRadius: 5,
    },
    logoutText: {
        color: "#FFFFFF",
        fontWeight: "bold",
        fontSize: 20,
    },
});

export default Settings;