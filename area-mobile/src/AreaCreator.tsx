import React, { useState, useContext, useEffect } from 'react';
import { View, Text, TextInput, Image, TouchableOpacity, Vibration, StyleSheet, Alert, ActionSheetIOS} from 'react-native';
import { postArea, editArea } from './api/AreaCalls';
import { Ionicons } from "@expo/vector-icons";
import ColorPicker from 'react-native-wheel-color-picker';
import { AuthContext } from '../context/AuthContext';
import { AreaInterface } from './interfaces/AreaInterface';
import { cleanName } from './utils';

const AreaCreator = ({ navigation, route }) => {
    let { area } = route.params;
    const [ newColor, setNewColor ] = useState(area.color);
    const [ newName, setName ] = useState(area.name);
    const { selectedArea, setArea } = useContext(AuthContext);

    useEffect(() => {
        if (area)
            setArea(area);
    }, []);

    useEffect(() => {
        if (selectedArea)
            area = selectedArea
    });

    async function handleSave() {
        area.name = newName;
        area.color = newColor;
        area.toggled = true;
        const resp:any = await editArea(area.id, area);
        navigation.navigate("Home");
    };

    async function handleCreate() {
        if (!area.action_name || !area.reaction_name) {
            alert("You need to add an action and a reaction");
            return;
        }
        area.name = newName;
        area.color = newColor;
        area.toggled = true;
        const resp:any = await postArea(area);
        console.log(resp);
        navigation.navigate("Home");
        setArea({} as AreaInterface);
    }

    return (
        <View style={{ flex: 1, justifyContent: "center", alignItems: "center",backgroundColor: "#FFFFFF",}}>
            <Ionicons
                name="return-down-back-outline"
                size={40} color="#007AFF"
                onPress={() => {navigation.navigate("Home"); Vibration.vibrate(50)}}
                style={{position: "absolute", left: '6%', top: '5%'}}/>
            <Image style={styles.logo} source={require("../assets/Logo_white.png")} />
            <View style={styles.addButtons}>
                <TouchableOpacity style={styles.addA} onPress={() => {navigation.navigate("ActionProvider")}}>
                    <Text style={styles.title}>{selectedArea.action && selectedArea.action_name ? cleanName(selectedArea.action) : "Add an Action"}</Text>
                </TouchableOpacity>
                <TouchableOpacity style={styles.addREA} onPress={() => {navigation.navigate("ReactionProvider")}}>
                    <Text style={styles.title}>{selectedArea.reaction && selectedArea.reaction_name ? cleanName(selectedArea.reaction) : "Add a REAction"}</Text>
                </TouchableOpacity>
            </View>
            <View style={styles.addDesc}>
                <TextInput
                    style={styles.description}
                    placeholder="Add a name"
                    onChangeText={(newText) => setName(newText)}
                    defaultValue={area.name}
                    maxLength={80}
                    multiline={true}/>
            </View>
            <View style={styles.colorPickerView}>
                <ColorPicker
                    color={area.color}
                    row={true}
                    thumbSize={50}
                    sliderSize={40}
                    sliderHidden={false}
                    gapSize={10}
                    discrete={false}
                    discreteLength={5}
                    noSnap={false}
                    shadeWheelThumb={true}
                    shadeSliderThumb={true}
                    palette={ ["#FF0000", "#00FF00", "#0000FF", "#FFFFFF", "#000000"] }
                    onColorChange={(color) => {setNewColor(color);}}
                    onColorChangeComplete={(color) => {setNewColor(color);}}
                />
            </View>
            {(area.action || area.reaction) ?
                <TouchableOpacity style={newName ? styles.saveButton : styles.saveButtonDisabled} onPress={handleSave} disabled={newName ? false : true}>
                    <Text style={styles.title}>{"Save"}</Text>
                </TouchableOpacity>
            :
                <TouchableOpacity style={newName ? styles.saveButton : styles.saveButtonDisabled} onPress={handleCreate} disabled={newName ? false : true}>
                    <Text style={styles.title}>{"Create"}</Text>
                </TouchableOpacity>}
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
        width: 430,
        height: 120,
        marginTop: '5%',
        alignItems: "center",
    },
    title: {
        color: "#FFFFFF",
        fontSize: 25,
        fontWeight: "bold",
        marginTop: '2%',
    },
    description: {
        color: "#FFFFFF",
        fontSize: 25,
        fontWeight: "bold",
        marginTop: '2%',
        alignContent: "center",
        justifyContent: "center",
        textAlign: 'center',
        maxWidth: '90%',
    },
    addButtons: {
        height: '15%',
        flexDirection: 'row',
        padding: '3%',
    },
    addA: {
        backgroundColor: "#FF66C4",
        width: "50%",
        borderRadius: 10,
        flexDirection: 'row',
        justifyContent: 'center',
        alignItems: 'center',
        margin: '1%',
    },
    addREA: {
        backgroundColor: "#FFC144",
        width: "50%",
        borderRadius: 10,
        flexDirection: 'row',
        justifyContent: 'center',
        alignItems: 'center',
        margin: '1%',
    },
    addDesc: {
        backgroundColor: "#CB6CE6",
        height: '16%',
        width: "95%",
        borderRadius: 10,
        justifyContent: 'center',
        alignItems: 'center',
        textAlign: 'center',
    },
    colorPickerView: {
        backgroundColor: "#DDDDDD",
        borderRadius: 10,
        margin: '3%',
        height: '20%',
        padding: '3%',
        width: "95%",
        justifyContent: 'center',
        alignItems: 'center',
    },
    saveButton: {
        backgroundColor: "#5555FF",
        height: '9%',
        width: "35%",
        borderRadius: 10,
        justifyContent: 'center',
        alignItems: 'center',
        margin: '3%',
    },
    saveButtonDisabled: {
        backgroundColor: "#DDDDDD",
        height: '9%',
        width: "35%",
        borderRadius: 10,
        justifyContent: 'center',
        alignItems: 'center',
        margin: '3%',
    },
});

export default AreaCreator;