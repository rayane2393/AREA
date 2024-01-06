import React, { useState, useContext, useCallback, useEffect } from "react";
import {
  View,
  StyleSheet,
  TextInput,
  Image,
  Text,
  TouchableOpacity,
  Vibration,
} from "react-native";
import { AuthContext } from "../context/AuthContext";

import { save } from '../context/SecureStore';
import { postLogin } from './api/ApiCalls';
import { WebView } from 'react-native-webview';

import { postDiscordOAuth, postGithubOAuth } from "./api/AuthCalls";
import * as Google from 'expo-auth-session/providers/google';

// android  447998894360-8j4bqgucas754gp2rjtkoevt4rr64mnf.apps.googleusercontent.com

const Login = ({ navigation }) => {
  const { login, setLoading } = useContext(AuthContext);
  const [email, defaultEmail] = useState("");
  const [password, defaultPassword] = useState("");
  const [githubAuth, setGithubAuth] = useState(false);
  const [discordAuth, setDiscordAuth] = useState(false);
  const [accessToken, setAccessToken] = useState(null);
  const [request, response, promptAsync] = Google.useAuthRequest({
        androidClientId: '447998894360-8j4bqgucas754gp2rjtkoevt4rr64mnf.apps.googleusercontent.com',
  });

  useEffect(() => {
    if (response?.type === 'success') {
      setAccessToken(response.params.access_token);
    }
  }, [response]);

  const fetchToken:any = useCallback(async () => {
    setLoading(true);
    try {
      Vibration.vibrate(80);
      const token:string = await postLogin(email, password);
      if (token.length > 0) {
        save("token", token);
        setLoading(false);
        login("LoggedIn")
      }
    } catch (error) {
      setLoading(false);
      alert("Wrong email or password");
      console.error(error);
    }
  }, [email, password]);

  function handleWebViewWarning(message:string) {
    console.log('WebView Warning:', message);
    const code:string = message.split('code=')[1];
      console.log(code);
      setGithubAuth(false);
      postGithubOAuth(code).then((token) => {
        if (token.length > 0) {
          save("token", token);
          login("LoggedIn")
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
  }, []);

  const authorizeWithGithub = () => {
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
    ];

    return (<WebView
      source={{
        uri: `https://github.com/login/oauth/authorize?client_id=${clientID}&redirect_uri=${redirectUrl}&scope=${githubScope.join(',')}`
      }}
      onMessage={(event) => {
        const data = event.nativeEvent.data;
        console.log(`Github Data : ${data}`);
        setGithubAuth(false);
      }}
      onError={(event) => {
        console.log(`Github Event : ${event}`);
        setGithubAuth(false);
      }}
    />);
  };

  const authorizeWithDiscord = () => {
    const clientID:string = "1168952746733473902";

    return (
      <WebView
        source={{
          uri: `https://discord.com/api/oauth2/authorize?client_id=${clientID}&redirect_uri=https%3A%2F%2Fgoogle.com&response_type=code&scope=identify%20email`,
        }}
        onMessage={(event) => {
          const data:string = event.nativeEvent.data;
          console.log(`Discord Data : ${data}`);
          setDiscordAuth(false);
        }}
        onError={(event) => {
          console.log(`Discord Event : ${event}`);
          setDiscordAuth(false);
        }}
        onNavigationStateChange={(navState) => {
          const url:string = navState.url;
          if (url.includes("google.com") && !url.includes("discord.com")) {
            const code:string = url.split("code=")[1];
            console.log(`Discord code : ${code}`);
            setDiscordAuth(false);
            postDiscordOAuth(code).then((token) => {
              if (token.length > 0) {
                save("token", token);
                login("LoggedIn")
              } else {
                alert("Discord Authentication failed, please use a different sign in method for now.");
              }
            });
          }
        }}
      />
    );
  };

  return (
    githubAuth ? authorizeWithGithub() :
    discordAuth ? authorizeWithDiscord() :
    <View style={styles.container}>

      <Image style={styles.logo} source={require("../assets/animation_area.gif")} />

      <View style={styles.formContainer}>

        <View style={styles.inputContainer}>
          <TextInput
              style={styles.input}
              placeholder="Email"
              onChangeText={(newText) => defaultEmail(newText)}
              defaultValue={email}
            />
        </View>
        <View style={styles.inputContainer}>
            <TextInput
              style={styles.input}
              secureTextEntry={true}
              placeholder="Password"
              onChangeText={(newText) => defaultPassword(newText)}
              defaultValue={password}
              />
        </View>

        <TouchableOpacity onPress={fetchToken}
                          style={[styles.button, email && password ? styles.buttonFilled : null]}
                          disabled={email && password ? false : true}>
          <Text style={[styles.buttonText, email && password ? styles.buttonTextFilled : null]}>{"Log In"}</Text>
        </TouchableOpacity>

        <TouchableOpacity style={styles.authButton} onPress={() => { navigation.navigate("Authentication"); }}>
          <Text style={styles.authButtonText}>{"No Account ? Create one."}</Text>
        </TouchableOpacity>
        {/* <TouchableOpacity style={styles.authButton} onPress={() => { navigation.navigate("ForgotPassword"); }}>
          <Text style={styles.authButtonText}>{"Forgot Password ?"}</Text>
        </TouchableOpacity> */}

        <TouchableOpacity
          style={styles.discordBtn}
          onPress={() => {setDiscordAuth(true)}}>
          <Image style={{width:25, height:25,}} source={require("../assets/discord-logo.png")}/>
          <Text style={styles.discordText}>Sign in with Discord</Text>
        </TouchableOpacity>

        {/* <TouchableOpacity
          style={styles.discordBtn}
          onPress={() => {promptAsync();}}>
          <Image style={{width:25, height:25,}} source={require("../assets/google.png")}/>
          <Text style={styles.discordText}>Sign in with Google</Text>
        </TouchableOpacity> */}

        <TouchableOpacity
          style={styles.githubBtn}
          onPress={() => {setGithubAuth(true)}}>
          <Image style={{width:25, height:25,}} source={require("../assets/github-logo.png")}/>
          <Text style={styles.githubText}>Sign in with Github</Text>
        </TouchableOpacity>
      </View>

    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#FFFFFF",
    alignItems: "center",
    justifyContent: "center",
  },
  enterText: {
    fontSize: 18,
    fontWeight: "bold",
    marginBottom: '5%',
    textAlign: "center",
  },
  logo: {
    width: 250,
    height: 250,
    marginBottom: 50,
  },
  formContainer: {
    width: "80%",
  },
  inputContainer: {
    backgroundColor: "#F2F2F2",
    borderRadius: 10,
    marginBottom: '4%',
  },
  input: {
    padding: 10,
  },
  button: {
    backgroundColor: "#C4C4C4",
    borderRadius: 10,
    padding: 10,
    alignItems: "center",
  },
  buttonText: {
    color: "#FFFFFF",
    fontSize: 18,
  },
  buttonFilled: {
    backgroundColor: "#007AFF",
    elevation: 4,
  },
  buttonTextFilled: {
    color: "#FFFFFF",
  },
  authButton: {
    top: '5%',
    height: 25,
    width: "auto",
    alignItems: "center",
    justifyContent: "center",
  },
  authButtonText: {
    color: "#1155FF",
    fontSize: 12,
    fontWeight: "bold",
    textDecorationLine: "underline",
    height: 40,
    textAlign: "center",
    paddingTop: '5%',
    marginBottom: '5%',
    alignItems: "center",
  },
  discordBtn: {
    flexDirection: "row",
    flex: 1,
    top: '2%',
    alignItems: "center",
    borderRadius: 15,
    paddingVertical: "7%",
    marginHorizontal: "25%",
    marginTop: '3%',
    justifyContent: "center",
    backgroundColor: "#7289DA",
    elevation: 4,
  },
  discordText: {
    color: "#FFFFFF",
    fontSize: 12,
    fontWeight: "bold",
    height: 40,
    textAlign: "center",
    paddingVertical: '7%',
    marginLeft: '5%',
  },
  githubBtn: {
    flexDirection: "row",
    flex: 1,
    top: '2%',
    alignItems: "center",
    borderRadius: 15,
    paddingVertical: "7%",
    marginHorizontal: "25%",
    marginTop: '3%',
    justifyContent: "center",
    backgroundColor: "#000000",
    elevation: 4,
  },
  githubText: {
    color: "#FFFFFF",
    fontSize: 12,
    fontWeight: "bold",
    height: 40,
    textAlign: "center",
    paddingVertical: '7%',
    marginLeft: '5%',
  },
});

export default Login;