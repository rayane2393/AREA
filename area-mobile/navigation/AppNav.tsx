import { View, ActivityIndicator } from "react-native";
import React, { useContext } from "react";

import { NavigationContainer } from "@react-navigation/native";
import { createStackNavigator } from "@react-navigation/stack";

import { AuthContext } from "../context/AuthContext";

import Login from "../src/Login";
import Home from "../src/Home";
import Authentication from "../src/Authtentication";
import ForgotPassword from "../src/ForgotPassword";
import AreaCreator from "../src/AreaCreator";
import Settings from "../src/Settings";
import AddAction from "../src/AddAction";
import AddReaction from "../src/AddReaction";
import ActionProvider from "../src/ActionProvider";
import ReactionProvider from "../src/ReactionProvider";
import AddDetails from "../src/AddDetails";

const Stack:any = createStackNavigator();

const AppNav = () => {
  const { isLoading, userToken } = useContext(AuthContext);

  if (isLoading) {
    return (
      <View style={{ flex: 1, justifyContent: "center", alignItems: "center" }}>
        <ActivityIndicator size={"large"} />
      </View>
    );
  }

  return (
    <NavigationContainer>
      <Stack.Navigator
        screenOptions={{
          headerShown: false,
        }}
      >
        {userToken.length > 0 ? (
          <>
            <Stack.Screen name="Home" component={Home} />
            <Stack.Screen name="AreaCreator" component={AreaCreator} />
            <Stack.Screen name="Settings" component={Settings} />
            <Stack.Screen name="ActionProvider" component={ActionProvider} />
            <Stack.Screen name="ReactionProvider" component={ReactionProvider} />
            <Stack.Screen name="AddAction" component={AddAction} />
            <Stack.Screen name="AddReaction" component={AddReaction} />
            <Stack.Screen name="AddDetails" component={AddDetails} />
          </>
        ) : (
          <>
            <Stack.Screen name="Login" component={Login} />
            <Stack.Screen name="Authentication" component={Authentication} />
            <Stack.Screen name="ForgotPassword" component={ForgotPassword} />
          </>
        )}
      </Stack.Navigator>
    </NavigationContainer>
  );
};

export default AppNav;
