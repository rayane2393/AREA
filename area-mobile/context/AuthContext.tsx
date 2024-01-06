import React, { createContext, useState } from "react";
import { AreaInterface } from "../src/interfaces/AreaInterface";

export const AuthContext = createContext(null);

export const AuthProvider = ({children}) => {
    const [isLoading, setIsLoading] = useState(false);
    const [userToken, setUserToken] = useState('');
    const [selectedArea, setSelectedArea] = useState({} as AreaInterface);

    const login = (token:string) => {
        token.length > 0 ? console.log('CONTEXT : Logged In') : console.log("CONTEXT : Login Nope")
        setUserToken(token);
        setIsLoading(false);
    };

    const logout = () => {
        setUserToken("");
        setIsLoading(false);
    };

    const setLoading = (loadBool:boolean) => {
        setIsLoading(loadBool);
    }

    const setArea = (area:AreaInterface) => {
        setSelectedArea(area);
    }

    return (
        <AuthContext.Provider value={{
            login, logout, isLoading, setLoading, userToken,

            selectedArea, setArea
            }}>
                {children}
        </AuthContext.Provider>
    );
}
