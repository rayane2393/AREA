import { Image, StyleSheet, Text, TextInput, TouchableOpacity, View } from 'react-native';
import React, { useState } from 'react';

const ForgotPassword = () => {
  const [email, setEmail] = useState('');

  const defaultEmail = (text: string):void => {
    setEmail(text);
  }

  const checkEmail = (email: string):boolean => {
    let reg:RegExp = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w\w+)+$/;
    if (reg.test(email) === false) {
      alert("Please enter a valid email address");
      return false;
    }
    else {
      alert("Email sent to " + email + "");
    }
  }

  const fetchToken = ():void => {
    checkEmail(email);
  }

  return (
    <View style={styles.container}>
      <Image style={styles.logo} source={require("../assets/animation_area.gif")} />
      <View style={styles.formContainer}>
        <Text style={styles.enterText}>{"Enter your email to reset your password"}</Text>
        <View style={styles.inputContainer}>
          <TextInput
            style={styles.input}
            placeholder="Email"
            onChangeText={(newText) => defaultEmail(newText)}
            defaultValue={email}
          />
        </View>
        <TouchableOpacity onPress={fetchToken} style={[styles.button, email ? styles.buttonFilled : null]}>
          <Text style={[styles.buttonText, email ? styles.buttonTextFilled : null]}>{"Send Reset Link"}</Text>
        </TouchableOpacity>
      </View>
    </View>
  );
}

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
    marginBottom: '5%',
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
  },
  buttonFilled: {
    backgroundColor: "#007AFF",
  },
  buttonTextFilled: {
    color: "#FFFFFF",
  },
});

export default ForgotPassword;