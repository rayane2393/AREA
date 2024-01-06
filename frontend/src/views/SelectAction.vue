<template>
  <div class="background">
    <div v-if="showDialog" class="dialog-overlay">
      <div class="dialog">
        <h3>{{ capitalizeFirstLetter(formatServiceName(currentAction.name)) }}</h3>
        <p>{{ capitalizeFirstLetter(formatServiceName(currentAction.description)) }}</p>
        <form @submit.prevent="confirmArguments(selectedService)">
          <input type="text" v-model="currentArguments" :placeholder="'Enter ' + formatServiceName(argumentName[0])">
          <button type="submit">Confirm</button>
          <button @click="closeDialog">Cancel</button>
        </form>
      </div>
    </div>

    <div class="home">
      <div class="page-button">
        <router-link to="/areas" class="logo-link">
          <img src="../assets/logo/Logo_white.png" alt="Logo" class="logo" />
        </router-link>
        <router-link to="/select-service-if" class="return-link">
          <img src="../assets/HomePage_Web/return.png" alt="Return" class="return" />
        </router-link>
      </div>
      <div class="title-page">Select an action for {{ selectedServiceAction }}</div>
      <div class="service-container">
        <div v-for="action in actions" :key="action.name" class="service-rectangle" @click="selectAction(action, selectedService, action.arguments)">
          <div class="service-image">
            <img :src="getImage(selectedServiceAction)" />
          </div>
          <div class="service-title">{{ formatServiceName(action.name) }}</div>
          <div class="service-description">{{ capitalizeFirstLetter(action.description) }}</div>
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
      selectedService: localStorage.getItem('selectedServiceActionTMP'),
      selectedServiceAction: '',
      actions: [],
      showDialog: false,
      currentArguments: '',
      currentAction: null,
      argumentName: [],
    };
  },
  mounted() {
    this.selectedServiceAction = localStorage.getItem('selectedServiceActionTMP');
    if (!this.selectedServiceAction) {
      this.$router.push('/select-service-if');
    } else {
      this.fetchActions();
    }
  },
  methods: {
    getImage(serviceName) {
      try {
        return require(`../assets/logo_services/${serviceName}.png`);
      } catch {
        return require(`../assets/logo_services/none.png`);
      }
    },
    async fetchActions() {
      try {
        const response = await axios.get('https://back-area.lekmax.fr/about.json');
        const service = response.data.server.services.find(s => s.name === this.selectedServiceAction);
        if (service && service.actions) {
          this.actions = service.actions;
        }
      } catch (error) {
        console.error('Error fetching actions:', error);
      }
    },
    selectAction(action, serviceName, arg) {
      if (arg && arg.length > 0) {
        this.argumentName = arg;
        this.currentAction = action;
        this.showDialog = true;
        document.body.style.overflow = 'hidden';
      } else {
        localStorage.setItem('selectedAction', action.name);
        localStorage.setItem('selectedServiceAction', serviceName);
        this.$router.push('/createArea');
      }
    },
    confirmArguments(serviceName) {
      localStorage.setItem('selectedActionArguments', this.argumentName[0]);
      localStorage.setItem(this.argumentName[0], this.currentArguments);
      localStorage.setItem('selectedAction', this.currentAction.name);
      localStorage.setItem('selectedServiceAction', serviceName);
      this.closeDialog();
      this.$router.push('/createArea');
    },
    closeDialog() {
      this.showDialog = false;
      document.body.style.overflow = '';
    },
    formatServiceName(name) {
      return name.replace(/_/g, ' ');
    },
    capitalizeFirstLetter(string) {
      return string.charAt(0).toUpperCase() + string.slice(1);
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
  z-index: 1000;
}

.dialog {
  background-color: white;
  padding: 30px;
  border-radius: 5px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.2);
}

.dialog form {
  display: flex;
  flex-direction: column;
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
  background-color: #FF66C4;
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
  color: #FF66C4;
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
  font-weight: bold;
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
