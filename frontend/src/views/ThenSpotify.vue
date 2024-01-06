<template>
  <div class="background">
    <div class="home">
      <div class="page-button">
        <router-link to="/areas" class="logo-link">
          <img src="../assets/logo/Logo_white.png" alt="Logo de votre site" class="logo" />
        </router-link>
        <router-link to="/select-service-then" class="return-link">
          <img src="../assets/HomePage_Web/return.png" alt="Logo de votre site" class="return" />
        </router-link>
      </div>
      <div class="title-page">Select a Spotify Reaction.</div>
      <div class="service-container">
        <div class="service-rectangle" @click="UpdateService">
          <div class="service-image">
            <img src="../assets/logo_services/spotify.png" />
          </div>
          <div class="service-title">
            <div>Plays the playlist <br>x</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>



<script>
import { NameActionService } from '../global.js';

export default {
  methods: {
    UpdateService() {
      NameActionService.value = "Spotify";
      console.log("Selected then:", NameActionService.value);
      this.$router.push({
          name: 'createArea',
      });
    },
  },
};
</script>

<style scoped>

.page-button {
  flex-direction: row-reverse;
  display: inline-flex;
  margin-left: 1%;
}

.return-link {
  margin-top: 1%;
  display: inline-block;
  height: 100%;
  background-color: white;
  padding: 1%;
  border-radius: 10px;
  transition: filter 0.3s ease;
}

.return {
  transition: filter 0.3s ease;
  height: 50%;
  width: 60%;
  filter: brightness(0.7);
}

.return-link:hover {
  background-color: #bbbbbb;
}

.return-link:hover .return {
  filter: brightness(1);
}

.logo-link {
  display: inline-block;
  height: 100%;
}

.logo {
  transition: filter 0.3s ease;
}

.logo-link:hover .logo {
  filter: brightness(0.8);
}

.background {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: #fFC144;
}

.title-page {
  text-align: center;
  font-weight: bold;
  color: #fFC144;
  font-size: 400%;
  margin-top: 2%;
  background-color: white;
  padding: 3%;
  margin-left: 28%;
  margin-right: 28%;
  border-radius: 10px;
}
.logo {
  margin-right: 88%;
  margin-top: 1%;
  border-radius: 10px;
  height: 10%;
  width: 10%;
}

.service-container {
  margin-top: 9%;
  display: flex;
  overflow-x: auto;
  width: 100%;
  margin-left: 0.3%;
  gap: 1%;
  white-space: nowrap;
  margin-left: 1%;
}

.service-rectangle {
  flex: 0 0 calc(15.5% - 1%);
  height: 36vh;
  border-radius: 10px;
  display: flex;
  flex-direction: column;
  text-align: center;
  transition: background-color 0.3s, color 0.3s;
  background-color: #FFFFFF;
}
.service-image img {
  height: 60%;
  width: 20%;
  margin-top: 10%;
}

.service-title {
  margin-top: 6%;
  margin-left: 10%;
  padding-bottom: 50%;
  flex: 1;
  width: 80%;
  border-top: 2px solid #ccc;
  display: flex;
  align-items: center;
  text-align: center;
  justify-content: center;
  font-weight: bold;
  font-size: 140%;
  color: #9d9d9d;
}

.service-rectangle:hover {
  background-color: #CCCCCC;
}

.service-rectangle:hover .service-title {
  border-top: 1px solid white;
  color: white;
}

</style>
