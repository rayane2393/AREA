import { StyleSheet } from "react-native";

const buttons = StyleSheet.create({
    button: {
        backgroundColor: "#FFFFFF",
        borderRadius: 25,
        width: 100,
        height: 100,
        justifyContent: "center",
        alignItems: "center",
        marginBottom: '10%',
    },
    buttonText: {
        color: "#000000",
        fontSize: 20,
        fontWeight: "bold",
    },
});

const actionButtonStyles = StyleSheet.create({
    container: {
        backgroundColor: "#FFFFFF",
        width: '90%',
        height: 140,
        borderRadius: 25,
        marginBottom: '4%',
        elevation: 5,
        justifyContent: "center",
        alignContent: "center",
        alignItems: "center",
    },
    icon: {
        position: "absolute",
        left: 'auto',
        right: 'auto',
        top: '6%',
    },
    desc: {
        fontSize: 22,
        flex: 1,
        fontWeight: "bold",
        justifyContent: "center",
        alignItems: "center",
        textAlign: "center",
        color: "#000000",
        padding: '5%',
        top: '20%',
    },
});

export { buttons, actionButtonStyles };
