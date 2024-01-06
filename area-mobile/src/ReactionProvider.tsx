import React, { useEffect, useState, useContext } from "react";
import { View, StyleSheet, Text, ScrollView, Vibration, TouchableOpacity} from "react-native";
import { getReactionServices } from "./api/AreaCalls";
import { AuthContext } from "../context/AuthContext";
import { cleanName } from "./utils";

const ReactionProvider = ({ navigation }) => {
    const [actionProviders, setReactionProviders] = useState([]);
    const { selectedArea, setArea } = useContext(AuthContext);

    async function renderServices() {
        const providers:any = await getReactionServices();
        setReactionProviders(providers);
    }

    useEffect(() => {
        renderServices();
    }, []);

    return (
        <View style={styles.container}>
            <Text style={styles.title}>Choose a Service for Reactions</Text>
            <ScrollView style={styles.reaContainer}>
                <View style={styles.serviceView}>
                    {actionProviders.map((provider, index) => (
                        <TouchableOpacity
                            key={index}
                            style={styles.button}
                            onPress={() => {
                                navigation.navigate("AddReaction", { name: provider.name, icon: provider.name });
                                Vibration.vibrate(50);
                                setArea({
                                    ...selectedArea,
                                    reaction: provider.name,
                                });
                            }}
                        >
                            <Text style={styles.buttonText}>{cleanName(provider.name)}</Text>
                        </TouchableOpacity>
                    ))}
                </View>
            </ScrollView>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        backgroundColor: "#FFC144",
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
    reaContainer: {
        backgroundColor: "#FFC144",
        flex: 1,
        width: "100%",
        alignContent: "center",
    },
    serviceView: {
        paddingVertical: "2%",
        justifyContent: "center",
        alignItems: "center",
        marginVertical: "2%",
        flexDirection: 'row',
        flexWrap: 'wrap',
    },
    title: {
        marginTop: '12%',
        fontSize: 35,
        textAlign: "center",
        color: "#FFFFFF",
        fontWeight: "bold",
        marginBottom: '3%',
    },
    button: {
        backgroundColor: "#FFFFFF",
        borderRadius: 25,
        width: 100,
        height: 100,
        justifyContent: "center",
        alignItems: "center",
        marginBottom: '6%',
        marginHorizontal: '2%',
    },
    buttonText: {
        color: "#000000",
        fontSize: 20,
        fontWeight: "bold",
    },
});

export default ReactionProvider;