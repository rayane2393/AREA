import React, { useState, useContext } from 'react';
import { useNavigation } from '@react-navigation/native';
import { View, Text, TextInput, StyleSheet, TouchableOpacity} from 'react-native';
import { AuthContext } from '../context/AuthContext';

const AddDetails = ({ route }) => {
    const { args } = route.params;
    const navigation:any = useNavigation();
    const [inputValues, setInputValues] = useState(Array(args.length).fill(''));
    const { selectedArea, setArea } = useContext(AuthContext);

    const cleanArg = (arg:string) => {
        arg = arg.replace("_", " ");
        arg = arg.charAt(0).toUpperCase() + arg.slice(1);
        return arg;
    }

    const handleInputChange = (text, index) => {
        setInputValues(inputValues[index] = text);
    };

    const handleDetailSave = () => {
        console.log("Details saved");
        console.log(inputValues);
        args.forEach((arg:string) => {
            // setInputValues(inputValues[index] = arg);
            if (arg === "summoner_name") {
                setArea({
                    ...selectedArea,
                    summoner_name: inputValues,
                });
            }
            if (arg === "city") {
                setArea({
                    ...selectedArea,
                    city: inputValues,
                });
            }
            if (arg === "artist_id") {
                setArea({
                    ...selectedArea,
                    artist_id: inputValues,
                });
            }
        });
        navigation.pop(3);
    };

    return (
        <View style={styles.container}>
            {args.map((arg:string, index:number) => (
                <View key={index} style={styles.inputContainer}>
                    <Text style={styles.detail}>{cleanArg(arg)}</Text>
                    <TextInput
                        style={styles.input}
                        onChangeText={(text:string) => handleInputChange(text, index)}
                    />
                </View>
            ))}
            <TouchableOpacity
                style={styles.saveButton}
                onPress={() => handleDetailSave()}
            >
                <Text style={styles.title}>Save</Text>
            </TouchableOpacity>
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        backgroundColor: "#FFFFFF",
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
    },
    label: {
        marginBottom: 10,
    },
    inputContainer: {
        marginBottom: '4%',
        alignItems: "center",
        width: '100%',
    },
    input: {
        backgroundColor: "#F2F2F2",
        borderRadius: 10,
        padding: 10,
        width: '50%',
    },
    detail: {
        color: "#000000",
        fontSize: 25,
        fontWeight: "bold",
        marginTop: '2%',
    },
    title: {
        color: "#FFFFFF",
        fontSize: 25,
        fontWeight: "bold",
        marginTop: '2%',
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

export default AddDetails;