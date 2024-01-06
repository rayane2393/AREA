import React, { useEffect, useState, useContext } from "react";
import { View, StyleSheet, Text, ScrollView, TouchableOpacity} from "react-native";
import { useNavigation } from '@react-navigation/native';
import { Ionicons } from "@expo/vector-icons";
import { getActions, getActionArgument } from "./api/AreaCalls";
import { AuthContext } from "../context/AuthContext";
import { actionButtonStyles } from "./style";
import { nameToIonicon, cleanName } from "./utils";

const ActionButton = ({ icon, description, name, service }) => {
    const navigation:any = useNavigation();
    const [adesc, setAdesc] = useState(description);
    const maxDescLength:number = 72;
    const { selectedArea, setArea } = useContext(AuthContext);

    icon = nameToIonicon(service);

    useEffect(() => {
        if (description.length > maxDescLength) {
            setAdesc(description.substr(0, maxDescLength) + "...");
        }
    }, []);

    useEffect(() => {
        if (description.length > maxDescLength) {
            setAdesc(description.substr(0, maxDescLength) + "...");
        }
    }, [description]);

    async function selectAction() {
        const args:any = await getActionArgument(service, name);
        if (args.length > 0) {
            // navigate to details
            setArea({
                ...selectedArea,
                action_name: name,
            });
            navigation.navigate("AddDetails", { args: args });
            return;
        }
        setArea({
            ...selectedArea,
            action_name: name,
        });
        navigation.pop(2);
    }

    return (
        <TouchableOpacity style={actionButtonStyles.container} onPress={selectAction}>
            <Ionicons name={icon} size={35} color="#007AFF" style={actionButtonStyles.icon}/>
            <Text style={actionButtonStyles.desc}>{adesc}</Text>
        </TouchableOpacity>
    );
}

const AddAction = ({ route }) => {
    const { name, icon } = route.params;
    const [actionServices, setActionServices] = useState([]);

    async function renderServices() {
        const resp:any = await getActions(name);

        setActionServices(
            resp.actions.map(
                (action, index) => (
                    <ActionButton
                        key={index}
                        icon={icon}
                        description={cleanName(action.description)}
                        name={action.name}
                        service={name}
                    />
                )
            )
        );
    }

    useEffect(() => {
        renderServices();
    }, []);

    return (
        <View style={styles.container}>
            <Text style={styles.title}>Add an Action</Text>
            <ScrollView style={styles.aContainer}>
                <View style={styles.serviceView}>
                    {actionServices}
                </View>
            </ScrollView>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        backgroundColor: "#FF66C4",
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
    },
    logo: {
        width: 420,
        height: 200,
        resizeMode: "contain",
        marginBottom: '10%',
    },
    aContainer: {
        backgroundColor: "#FF66C4",
        flex: 1,
        width: "100%",
        alignContent: "center",
    },
    serviceView: {
        paddingVertical: "2%",
        justifyContent: "center",
        alignItems: "center",
        marginVertical: "2%",
    },
    title: {
        marginTop: '12%',
        fontSize: 40,
        color: "#FFFFFF",
        fontWeight: "bold",
        marginBottom: '3%',
    },
    button: {
        backgroundColor: "#007AFF",
        borderRadius: 25,
        width: '90%',
        height: 50,
        justifyContent: "center",
        alignItems: "center",
        marginBottom: '10%',
    },
    buttonText: {
        color: "#FFFFFF",
        fontSize: 20,
        fontWeight: "bold",
    },
});

export default AddAction;