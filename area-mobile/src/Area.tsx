import React, { ReactElement, useEffect, useState, useContext } from "react";
import { View, StyleSheet, Text, Switch, Vibration, Alert} from "react-native";
import { Ionicons } from "@expo/vector-icons";
import { toggleArea, deleteArea } from "./api/AreaCalls";
import { nameToIonicon } from "./utils";
import { AuthContext } from "../context/AuthContext";
import { LinearGradient } from 'expo-linear-gradient';

const isColorDark = (bgColor:string) => {
    let r:number = parseInt(bgColor.substring(1,3),16);
    let g:number = parseInt(bgColor.substring(3,5),16);
    let b:number = parseInt(bgColor.substring(5,7),16);
    return (((r*299)+(g*587)+(b*114))/1000 < 128);
}

const textStyleSheet = (bgColor:string, toggled:boolean) => {
    let color:string = isColorDark(bgColor) ? "#FFFFFF" : "#000000"
    color = toggled ? color : (color + "88");
    const styles:any = StyleSheet.create({
        description: {
            fontSize: 26,
            flex: 1,
            fontWeight: "bold",
            justifyContent: "center",
            alignItems: "center",
            textAlign: "center",
            color: color,
            top: '5%',
            padding: '5%',
        },
    });

    return styles.description;
}

const bgStyleSheet = (bgColor:string, toggled:boolean) => {
    // If the are isnt toggled, the color is paler
    const fbgColor:string = toggled ? bgColor : (bgColor + "44");

    const styles = StyleSheet.create({
        area: {
            backgroundColor: fbgColor,
            width: '90%',
            height: 240,
            borderRadius: 25,
            borderWidth: bgColor === "#FFFFFF" ? 1 : 0,
            marginBottom: '4%',
            elevation: toggled ? 6 : 0,
        },
    });

    return styles.area;
}

const getActionIcon = (actionService:string, color:string) => {
    let icon:ReactElement = null;
    let logoColor:string = isColorDark(color) ? "#FFFFFF" : "#000000"

    const styles:any = StyleSheet.create({
        icon: {
            position: "absolute",
            left: '6%',
            top: '80%',
        },
    });

    icon = <Ionicons name={nameToIonicon(actionService) as any} size={28} color={logoColor} style={styles.icon}/>;

    return icon;
}

const getReactionIcon = (reactionService:string, color:string) => {
    let icon:ReactElement = null;
    let logoColor:string = isColorDark(color) ? "#FFFFFF" : "#000000"

    const styles = StyleSheet.create({
        icon: {
            position: "absolute",
            left: '22%',
            top: '80%',
        },
    });

    icon = <Ionicons name={nameToIonicon(reactionService) as any} size={28} color={logoColor} style={styles.icon}/>;

    return icon;
}

const Area = ({ navigation, area }) => {
    const [isEnabled, setIsEnabled] = useState(area.toggled);
    const { setLoading } = useContext(AuthContext);
    const toggleSwitch = ():void => {
        Vibration.vibrate(50);
        setIsEnabled(previousState => !previousState);
        area.toggled = !area.toggled;
        console.log(area.id, !isEnabled);
        toggleArea(area.id, !isEnabled);
    };

    const handleDelete = async () => {
        Alert.alert(
            "Delete Area",
            "Are you sure you want to delete this area?",
            [
                {
                    text: "Cancel",
                    onPress: () => console.log("Cancel Pressed"),
                    style: "cancel"
                },
                {
                    text: "Delete",
                    onPress: () => {
                        setLoading(true);
                        deleteArea(area.id, area.action);
                        console.log(`Area ${area.id} deleted ${area.action}`);
                        setLoading(false);
                    }
                }
            ],
            { cancelable: true }
        );
    };

    useEffect(() => {
        setIsEnabled(area.toggled);
    } , [area.toggled]);

    return (
        <View style={bgStyleSheet(area.color, isEnabled)}>
            <LinearGradient
            colors={['rgba(0,0,0,0)', 'rgba(0,0,0,0.22)']}
            style={{ position: 'absolute', left: 0, right: 0, top: 0, height: '100%', borderRadius: 25 }}
            >
            <Text style={textStyleSheet(area.color, isEnabled)}>{area.name}</Text>
            <Switch
                style={{transform: [{scaleX: 1.6}, {scaleY: 1.6}], position: "absolute", right: '5%', top: '76%'}}
                trackColor={{false: '#767577', true: '#44FF44' }}
                thumbColor={'#f4f3f4'}
                ios_backgroundColor="#3e3e3e"
                onValueChange={toggleSwitch}
                value={isEnabled}
            />
            <Ionicons name="trash-bin-outline" size={24} color={isColorDark(area.color) ? "#FFFFFF" : "#000000"} onPress={handleDelete} style={{position: "absolute", right: '4%', top: '4%'}} />
            <Ionicons
                name="create-outline"
                size={28}
                color={isColorDark(area.color) ? "#FFFFFF" : "#000000"}
                onPress={() => {
                    Vibration.vibrate(50);
                    navigation.navigate("AreaCreator", {
                                        area: area,
                                    });
                                }}
                style={{position: "absolute", right: '23%', top: '79%',}}/>
            {getActionIcon(area.action, area.color)}
            <Ionicons name="arrow-forward" size={24} color={isColorDark(area.color) ? "#FFFFFF" : "#000000"} style={{position: "absolute", left: '14%', top: '81%',}}/>
            {getReactionIcon(area.reaction, area.color)}
            </LinearGradient>
        </View>
    );
}

export default Area;