import React, { useContext, useEffect, useState } from "react";
import { useFocusEffect } from '@react-navigation/native';
import { View, Image, ScrollView, StyleSheet, Vibration, Text} from "react-native";
import { TouchableOpacity } from "react-native-gesture-handler";
import { AuthContext } from "../context/AuthContext";
import { Ionicons } from "@expo/vector-icons";

import { getAreas } from "./api/AreaCalls";
import { userInfo } from "./api/ApiCalls";
import Area from "./Area";
import { AreaInterface } from "./interfaces/AreaInterface";

const Home = ({ navigation }) => {
    const [areas, setAreas] = useState([]);
    const { setArea } = useContext(AuthContext);

    function cleanServices() {
      setArea({} as AreaInterface);
    }

    function parseAreas(areaJson:any, userData:any):AreaInterface[] {
      const areaArray:AreaInterface[] = [];

      if (areaJson.areas) {
        areaJson.areas.forEach((area: any) => {
          const newArea:AreaInterface = {
            id: area.id,
            name: area.name,
            color: area.color,
            action: area.action,
            reaction: area.reaction,
            toggled: area.toggled,
            action_name: area.action_name,
            reaction_name: area.reaction_name,
            summoner_name: area.summoner_name,
            city: area.city,
            artist_id: area.artist_id,
          };
          areaArray.push(newArea);
          try {
            if (!userData.user.github_id) {
              if (newArea.action === "github" || newArea.reaction === "github") {
                alert("You need to link your Github account to use the Github service.");
              }
            }
            if (!userData.user.discord_id) {
              if (newArea.action === "discord" || newArea.reaction === "discord") {
                alert("You need to link your Discord account to use the Discord service.");
              }
            }
          } catch (error) {
            // console.log("User might not connected with either github or discord");
          }
        });
      }
      return areaArray;
    }

  async function renderAreas() {
    const areaJson:any = await getAreas();
    const userData:any = userInfo();
    const areaArray:AreaInterface[] = parseAreas(areaJson, userData);

    if (areaJson.areas) {
      setAreas(areaArray.map((area: any, key:number) => (
        <Area navigation={navigation}
              key={key}
              area={area}
        />
      )));
    } else {
      setAreas([<Text key={0} style={styles.noAreaText}>No areas found. Add an area by clicking the + button below.</Text>]);
    }
  }

  useEffect(() => {
    cleanServices();
    renderAreas();
  }, []);

  useFocusEffect(
    React.useCallback(() => {
      renderAreas();
    }, [])
  );

  return (
    <View style={{ flex: 1, justifyContent: "center", alignItems: "center",backgroundColor: "#FFFFFF",}}>
      <Image style={styles.logo} source={require("../assets/Logo_white.png")} />
      <Ionicons name="settings-outline" size={30} color="#007AFF" onPress={() => {navigation.navigate("Settings"); Vibration.vibrate(50)}} style={{position: "absolute", right: '6%', top: '5%'}}/>
      <ScrollView style={styles.areascontainer}>
        <View style={styles.areaView}>
          {areas}
        </View>
        <TouchableOpacity style={styles.addArea} onPress={() => {navigation.navigate("AreaCreator", {area: {} as AreaInterface}); Vibration.vibrate(50); cleanServices();}}>
          <Ionicons name="add-circle-outline" size={70} color="#007AFF" />
        </TouchableOpacity>
      </ScrollView>
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
        width: 420,
        height: 110,
        marginTop: '10%',
        alignItems: "center",
    },
    title: {
        fontSize: 30,
        marginTop: '2%',
    },
    areascontainer: {
        backgroundColor: "#FAFAFA",
        flex: 1,
        width: "100%",
        alignContent: "center",
    },
    areaView: {
      paddingVertical: "2%",
      justifyContent: "center",
      alignItems: "center",
      marginVertical: "2%",
    },
    addArea: {
      flexDirection: "row",
      justifyContent: "center",
      alignItems: "center",
      borderRadius: 25,
      marginBottom: "2%",
    },
    noAreaText: {
      fontSize: 28,
      fontWeight: "bold",
      marginTop: '2%',
      padding: "2%",
      textAlign: "center",
    },
});

export default Home;
