import React, { useState, useCallback, useEffect } from "react";
import {
  View,
  StyleSheet,
  TextInput,
  Text,
  Image,
  TouchableOpacity,
  Vibration,
} from "react-native";
import { postAuth } from './api/ApiCalls';

const Authentication = ({ navigation }) => {
  const [username, defaultUsername] = useState("");
  const [email, defaultEmail] = useState("");
  const [password, defaultPassword] = useState("");
  const [confirmPassword, defaultConfirmPassword] = useState("");
  const [isDifferent, setIsDifferent] = useState(false);
  const [isErrorUsrName, setIsErrorUsrName] = useState(false);
  const [isErrorPassword, setIsErrorPassword] = useState(false);

  const checkEmail = (email: string) => {
    let reg:RegExp = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w\w+)+$/;
    if (reg.test(email) === false) {
      alert("Please enter a valid email address");
      return false;
    }
    return true;
  }

  const createAccount = useCallback(async () => {
    if (!checkEmail(email)) {
      return;
    }
    try {
      Vibration.vibrate(80);
      if (username.length === 0) {
        alert("Please enter a username");
        console.log("Please enter a username");
        setIsErrorUsrName(true);
        return;
      }
      if (email.length === 0) {
        alert("Please enter an email");
        console.log("Please enter an email");
        return;
      }
      if (password.length === 0) {
        alert("Please enter a password");
        console.log("Please enter a password");
        setIsErrorPassword(true);
        return;
      }
      const resp:any = await postAuth(username, password, email);
      console.log(resp);
      if (resp.status === 201 || resp.status === 200) {
        alert("Account created");
        navigation.navigate("Login");
      }
      if (resp.status === 406) {
        alert("This username is already taken");
        console.log("This username is already taken");
        return;
      }
      if (resp.status === 409) {
        alert("This email is already taken");
        console.log("This email is already taken");
        return;
      }
    } catch (error) {
      console.error(error);
    }
  }, [email, password]);

  useEffect(() => {
    if (password !== confirmPassword) {
      setIsDifferent(true);
    } else {
      setIsDifferent(false);
    }
  });

  return (
    <View style={styles.container}>
      <Image style={styles.logo} source={require("../assets/animation_area.gif")}/>
        <View style={styles.formContainer}>
          <View style={styles.inputContainer}>
            <TextInput
              style={styles.input}
              placeholder="Username"
              onChangeText={(newText) => defaultUsername(newText)}
              defaultValue={username}
            />
          </View>
          {/* {isErrorUsrName && <Text style={styles.errorTextUsername}>Please enter a username</Text>} */}
          <View style={styles.inputContainer}>
            <TextInput
              style={styles.input}
              placeholder="Email"
              onChangeText={(newText) => defaultEmail(newText)}
              defaultValue={email}
            />
          </View>
          {/* <View style={styles.inputContainer}>
            <TextInput
              style={styles.input}
              placeholder="Phone Number (Optional)"
              onChangeText={(newText) => defaultNumber(newText)}
              defaultValue={number}
            />
          </View> */}
          <View style={styles.inputContainer}>
            <TextInput
              style={styles.input}
              secureTextEntry={true}
              placeholder="Password"
              onChangeText={(newText) => defaultPassword(newText)}
              defaultValue={password}
            />
          </View>
          {/* {isErrorPassword && <Text style={styles.errorTextPassword}>Please enter a password</Text>} */}
          <View style={styles.inputContainer}>
            <TextInput
              style={styles.input}
              secureTextEntry={true}
              placeholder="Confirm Password"
              onChangeText={(newText) => defaultConfirmPassword(newText)}
              defaultValue={confirmPassword}
              />
          </View>
        </View>
        <TouchableOpacity
          onPress={createAccount}
          style={[styles.button, email && username && password && confirmPassword ? styles.buttonFilled : null]}
          disabled={email && username && password && confirmPassword ? false : true}
        >
          <Text style={styles.buttonText}> {"Create Account"} </Text>
        </TouchableOpacity>
        <TouchableOpacity style={styles.authButton} onPress={() => { navigation.navigate("Login"); }}>
          <Text style={styles.authButtonText}>{"Already have an account ? Log In."}</Text>
        </TouchableOpacity>
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
  buttonFilled: {
    backgroundColor: "#007AFF",
    elevation: 4,
  },
  buttonText: {
    color: "#FFFFFF",
    fontSize: 18,
  },
  buttonTextFilled: {
    color: "#FFFFFF",
  },
  authButton: {
    top: '1%',
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
  errorTextUsername: {
    position: "absolute",
    color: "#FF3333",
    fontSize: 11,
    height: 40,
    fontWeight: "bold",
    marginRight: "5%",
    marginLeft: "7%",
    marginTop: '44%',
  },
  errorTextPassword: {
    position: "absolute",
    color: "#FF3333",
    fontSize: 11,
    height: 40,
    fontWeight: "bold",
    marginRight: "5%",
    marginLeft: "7%",
    marginTop: '94%',
  },
});

export default Authentication;