import * as React from "react";

import AppNav from "./navigation/AppNav";
import { AuthProvider } from "./context/AuthContext";

import { LogBox } from "react-native";

// Ignore in app notification
LogBox.ignoreAllLogs();

import { Logs } from 'expo'

Logs.enableExpoCliLogging()

export default function App() {
  return (
    <AuthProvider>
      <AppNav />
    </AuthProvider>
  );
}