<template>
  <div class="background">
    <div class="home">
      <div class="page-button">
        <router-link to="/areas" class="logo-link">
          <img src="../assets/logo/Logo_white.png" alt="Logo" class="logo" />
        </router-link>
        <router-link to="/createArea" class="return-link">
          <img src="../assets/HomePage_Web/return.png" alt="Return" class="return" />
        </router-link>
      </div>
      <div class="title-page">Select a Service for your REAction</div>
      <div class="service-container">
        <div v-for="service in servicesWithReactions" :key="service.name" class="service-rectangle" @click="selectService(service.name)">
          <div class="service-image">
            <img :src="getImage(service.name)" />
          </div>
          <div class="service-title">{{ formatServiceName(service.name) }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      services: [],
    };
  },
  computed: {
    servicesWithReactions() {
      return this.services.filter((service) => service.reactions && service.reactions.length > 0);
    },
  },
  mounted() {
    this.getServices();
  },
  methods: {
    getImage(serviceName) {
      try {
        return require(`../assets/logo_services/${serviceName}.png`);
      } catch {
        return require(`../assets/logo_services/none.png`);
      }
    },
    selectService(serviceName) {
      localStorage.setItem('selectedServiceReactionTMP', serviceName);
      this.$router.push('/select-reaction');
    },
    formatServiceName(name) {
      return name.charAt(0).toUpperCase() + name.slice(1).replace(/_/g, ' ');
    },
    async getServices() {
      try {
        const response = await axios.get('https://back-area.lekmax.fr/about.json');
        this.services = response.data.server.services;
      } catch (error) {
        console.error('Error fetching services:', error);
      }
    },
  },
};
</script>

<style scoped>
.dialog-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}

.dialog {
  background-color: white;
  padding: 20px;
  border-radius: 5px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.2);
}

.dialog form {
  display: flex;
  flex-direction: column;
}

.dialog label {
  margin-top: 10px;
}

.dialog input[type="text"] {
  margin-bottom: 10px;
  padding: 8px;
  border-radius: 4px;
  border: 1px solid #ccc;
}

.dialog button {
  margin-top: 20px;
  padding: 10px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
.background {
  display: flex;
  flex-direction: column;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: #fFC144;
}

.home {
  margin: 5%;
  margin-top: 5%;
}

.page-button {
  display: flex;
  justify-content: space-between;
  margin-bottom: 2.5rem;
  flex-direction: column;
  align-items: flex-start;
}

.logo-link img {
  height: 50px;
}

.return-link img {
  height: 30px;
}

.logo {
  border-radius: 20px;
  transition: opacity 0.2s;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.5);
  margin-bottom: 10%;
}

.return-link {
  background-color: #ffffff;
  border-radius: 20px;
  transition: opacity 0.2s;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.5);
  transition: transform 0.3s ease;
}

.logo-link {
  transition: transform 0.3s ease;
}

.logo-link:hover, .return-link:hover {
  transform: translateX(-10px);
  opacity: 0.7;
}
.return {
  padding: 10px;
  transition: filter 0.3s ease;
  filter: brightness(0.7);
}
.title-page {
  font-weight: bold;
  font-size: 2.5em;
  letter-spacing: 1.5px;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.1);
  font-family: 'Arial';
  text-align: center;
  margin-top: 20px;
  margin-bottom: 4%;
  margin-left: 28%;
  margin-right: 28%;
  color: #fFC144;
  font-size: 300%;
  margin-top: 2%;
  padding: 3%;
  background-color: white;
  border-radius: 10px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.5);
}

.service-container {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 1.5rem;
  width: 100%;
}

.service-rectangle {
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #fff;
  border-radius: 10px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.5);
  transition: transform 0.2s, box-shadow 0.2s;
  padding: 4%;
  opacity: 0.8;
}

.service-rectangle:hover {
  transform: scale(1.03);
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.1);
  opacity: 1;
}

.service-image img {
  width: 100%;
  height: 150px;
  object-fit: cover;
  padding-bottom: 15%;
  border-bottom: 2px solid #ccc;
}

.service-title {
  text-transform: capitalize;
  font-size: 1.5em;
  margin-bottom: 10px;
  color: #333;
  margin: 1%;
  margin-top: 2%;
}

.service-description {
  font-size: 1em;
  line-height: 1.5;
  color: #666;
  max-width: 90%;
}

@media (max-width: 600px) {
  .background {
    padding: 1rem;
  }

  .title-page {
    font-size: 1.6rem;
  }

  .service-image img {
    height: 130px;
  }
}
</style>
