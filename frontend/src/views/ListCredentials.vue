<template>
  <div class="app">
    <Sidebar class="sidebar" />
    <div class="main">
      <DarkModeButton class="darkbutton" :initialState="isDarkMode" @switched="isDarkMode = $event" />








      <div class="side-content">
        <span class="title">Credentials</span>
        <button class="btn" @click="openCredentialsModal">Add Credentials</button>

        <div v-if="showAdditionalModal" class="modal">
          <button class="close-button" @click="closeModal">×</button>
          <div class="modal-content">
            <h2 class="modal-title">Add New Credentials</h2>
            <h4>Select an app or service to connect to</h4>
            <select v-model="selectedService" @change="openAdditionalModal" class="custom-dropdown">
              <option value="">Select an option</option>
              <option v-for="service in services" :key="service.id" :value="service">{{ service.name }}</option>
            </select>
            <component :is="currentModalComponent"></component>
          </div>

          <!--<div v-if="showAdditionalModal" class="modal">
              <div class="modal-content">
                <h2>{{ selectedService.name }} Details</h2>
                <h3>salut malke</h3>
                <button @click="closeAdditionalModal">Close</button>
              </div>
            </div>-->
        </div>
      </div>

      <div class="main-content">
        <div class="card-list" style="max-height: 880px; overflow-y: scroll;">
          <CardCredentials v-for="card in cardList" :key="card.id" :id="card.id" :title="card.title"
            :created_at="card.created_at" :updated_at="card.updated_at" :name="card.name" :types="card.types" />

        </div>


        <div class="top-bar">
          <div class="search-bar">
            <input type="text" placeholder="🔍 Search credentials...">
          </div>
          <div class="sort-dropdown">
            <select class="custom-dropdown">
              <option value="az">Sort by name (A-Z)</option>
              <option value="za">Sort by name (Z-A)</option>
              <option value="last_updated">Sort by last Updated</option>
              <option value="last_created">Sort by last Created</option>
            </select>
          </div>
        </div>
        




      </div>
    </div>
  </div>
</template>
  
<script>

import CardCredentials from '../component/CardCredentials.vue'; // Assurez-vous d'importer le composant Card

import ModalGitHubApi from './modalCredentials/ModalGithubApi.vue';

import ModalGitHubOAuth from './modalCredentials/ModalGithubOAuth2.vue';
import ModalSpotifyOAuth from './modalCredentials/ModalSpotifyOAuth2.vue';
export default {
  components: {
    Sidebar,
    DarkModeButton,
    ModalGitHubApi,
    ModalGitHubOAuth,
    CardCredentials,
    ModalSpotifyOAuth,
  },
  data() {
    return {
      cardList: [
        {
          id: 1,
          title: "Account Google perso",
          updated_at: "2023-09-03T14:30:00",
          created_at: "2023-09-03T14:30:00",
          types: "google",
          name: "Google Drive OAuth2 Api",
        },
        {
          id: 2,
          title: "Account Github perso",
          updated_at: "2023-09-03T14:30:00",
          created_at: "2023-09-03T14:30:00",
          types: "github",
          name: "Github OAuth2 Api",
        },
        {
          id: 3,
          title: "Account Spotify perso",
          updated_at: "2023-09-03T14:30:00",
          created_at: "2023-09-03T14:30:00",
          types: "spotify",
          name: "Spotify OAuth2 Api",
        },
        {
          id: 4,
          title: "Account Github pro",
          updated_at: "2023-09-03T14:30:00",
          created_at: "2023-09-03T14:30:00",
          types: "github",
          name: "Github Api",
        },
        {
          id: 5,
          title: "Account Google perso",
          updated_at: "2023-09-03T14:30:00",
          created_at: "2023-09-03T14:30:00",
          types: "google",
          name: "Google Drive OAuth2 Api",
        },
        {
          id: 6,
          title: "Account Github perso",
          updated_at: "2023-09-03T14:30:00",
          created_at: "2023-09-03T14:30:00",
          types: "github",
          name: "Github OAuth2 Api",
        },
        {
          id: 7,
          title: "Account Spotify perso",
          updated_at: "2023-09-03T14:30:00",
          created_at: "2023-09-03T14:30:00",
          types: "spotify",
          name: "Spotify OAuth2 Api",
        },
        {
          id: 8,
          title: "Account Github pro",
          updated_at: "2023-09-03T14:30:00",
          created_at: "2023-09-03T14:30:00",
          types: "github",
          name: "Github Api",
        },
        {
          id: 9,
          title: "Account Google perso",
          updated_at: "2023-09-03T14:30:00",
          created_at: "2023-09-03T14:30:00",
          types: "google",
          name: "Google Drive OAuth2 Api",
        },
        {
          id: 10,
          title: "Account Github perso",
          updated_at: "2023-09-03T14:30:00",
          created_at: "2023-09-03T14:30:00",
          types: "github",
          name: "Github OAuth2 Api",
        },
        {
          id: 11,
          title: "Account Spotify perso",
          updated_at: "2023-09-03T14:30:00",
          created_at: "2023-09-03T14:30:00",
          types: "spotify",
          name: "Spotify OAuth2 Api",
        },
        {
          id: 12,
          title: "Account Github pro",
          updated_at: "2023-09-03T14:30:00",
          created_at: "2023-09-03T14:30:00",
          types: "github",
          name: "Github Api",
        },
        {
          id: 13,
          title: "Account Google perso",
          updated_at: "2023-09-03T14:30:00",
          created_at: "2023-09-03T14:30:00",
          types: "google",
          name: "Google Drive OAuth2 Api",
        },
        {
          id: 14,
          title: "Account Github perso",
          updated_at: "2023-09-03T14:30:00",
          created_at: "2023-09-03T14:30:00",
          types: "github",
          name: "Github OAuth2 Api",
        },
        {
          id: 15,
          title: "Account Spotify perso",
          updated_at: "2023-09-03T14:30:00",
          created_at: "2023-09-03T14:30:00",
          types: "spotify",
          name: "Spotify OAuth2 Api",
        },
        {
          id: 16,
          title: "Account Github pro",
          updated_at: "2023-09-03T14:30:00",
          created_at: "2023-09-03T14:30:00",
          types: "github",
          name: "Github Api",
        },
      ],

      searchQueryService: '', // Pour la recherche
      showCredentialsModal: false,
      showAdditionalModal: false,
      selectedService: '',
      services: [
        { name: 'Github Api', id: 1 },
        { name: 'Github OAuth2', id: 2 },
        { name: 'Spotify OAuth2', id: 3 },
        // Ajoutez d'autres services ici
      ],
    };
  },
  computed: {
    currentModalComponent() {
      // Associez chaque option avec le composant modal correspondant
      const modalComponents = {
        'Github Api': ModalGitHubApi,
        'Github OAuth2': ModalGitHubOAuth,
        'Spotify OAuth2': ModalSpotifyOAuth,
        // Associez d'autres options avec leurs composants modaux respectifs
      };
      return modalComponents[this.selectedService.name] || null;
    },
  },
  methods: {
    closeModal() {
      this.showAdditionalModal = false;
    },
    openCredentialsModal() {
      this.showAdditionalModal = true;
    },
  },
};
</script>
  
  
  
  
<script setup>
import { ref, watch } from 'vue';
import Sidebar from '../component/MySideBar.vue';
import DarkModeButton from '../component/DarkButton.vue';

const isDarkMode = ref(true);
// Activer le mode sombre par défaut si l'utilisateur ne l'a pas configuré
if (localStorage.getItem("isDarkMode") === null) {
  isDarkMode.value = true;
}
const root = document.documentElement;

root.style.setProperty('--green', 'var(--green-dark)');
root.style.setProperty('--orange', 'var(--orange-dark)');
root.style.setProperty('--side-bar', 'var(--side-bar-dark)');
root.style.setProperty('--background', 'var(--background-dark)');
root.style.setProperty('--text', 'var(--text-dark)');
root.style.setProperty('--text-color', 'var(--text-color-dark)');
root.style.setProperty('--border-color', 'var(--border-color-dark)');
root.style.setProperty('--card-color', 'var(--card-color-dark)');
root.style.setProperty('--text-side-bar', 'var(--text-side-bar-dark)');
root.style.setProperty('--search-bar', 'var(--search-bar-dark)');
root.style.setProperty('--hover-color', 'var(--hover-color-dark)');
root.style.setProperty('--hover-text-color', 'var(--hover-text-color-dark)');



// Utilisez watch pour surveiller les changements de isDarkMode et mettre à jour le thème
watch(isDarkMode, (newDarkModeValue) => {
  const root = document.documentElement;

  if (isDarkMode.value) {
    root.style.setProperty('--green', 'var(--green-dark)');
    root.style.setProperty('--orange', 'var(--orange-dark)');
    root.style.setProperty('--side-bar', 'var(--side-bar-dark)');
    root.style.setProperty('--background', 'var(--background-dark)');
    root.style.setProperty('--text-color', 'var(--text-color-dark)');
    root.style.setProperty('--border-color', 'var(--border-color-dark)');
    root.style.setProperty('--text', 'var(--text-dark)');
    root.style.setProperty('--card-color', 'var(--card-color-dark)');
    root.style.setProperty('--text-side-bar', 'var(--text-side-bar-dark)');
    root.style.setProperty('--search-bar', 'var(--search-bar-dark)');
    root.style.setProperty('--hover-color', 'var(--hover-color-dark)');
    root.style.setProperty('--hover-text-color', 'var(--hover-text-color-dark)');
  } else {
    root.style.setProperty('--green', 'var(--green-light)');
    root.style.setProperty('--orange', 'var(--orange-light)');
    root.style.setProperty('--side-bar', 'var(--side-bar-light)');
    root.style.setProperty('--background', 'var(--background-light)');
    root.style.setProperty('--text-color', 'var(--text-color-light)');
    root.style.setProperty('--border-color', 'var(--border-color-light)');
    root.style.setProperty('--text', 'var(--text-light)');
    root.style.setProperty('--card-color', 'var(--card-color-light)');
    root.style.setProperty('--text-side-bar', 'var(--text-side-bar-light)');
    root.style.setProperty('--search-bar', 'var(--search-bar-light)');
    root.style.setProperty('--hover-color', 'var(--hover-color-light)');
    root.style.setProperty('--hover-text-color', 'var(--hover-text-color-light)');
  }
});
</script>
  
  
  
  
  
  
  
  
  
  
<style lang="scss">
:root {

  --hover-color-light: #DCDFE7;
  --hover-color-dark: #32394A;
  --hover-text-color-light: #555555;
  --hover-text-color-dark: #B8B2A8;

  --hover-color: --hover-color-light;
  --hover-text-color: --hover-text-color-light;

  --green-dark: #0FA552;
  --green-light: #13CE66;
  --orange-light: #FF6D5A;
  --orange-dark: #FF4D4D;
  --side-bar-dark: #252829;
  --side-bar-light: #FFFFFF;

  --background-dark: #1D1F20;

  --background-light: #FBFCFE;
  --text-color-light: #7D7D87;
  --text-color-dark: #E8E6E3;

  --text-light: #555555;
  --text-dark: #ebebeb;

  --border-color-dark: #736B5E;
  --border-color-light: #555555;

  --card-color-dark: #252829;
  --card-color-light: --side-bar-light;

  --text-side-bar-dark: #B8B2A8;
  --text-side-bar-light: #7D7D87;

  --search-bar-dark: #141516;
  --search-bar-light: --side-bar-light;
  --side-bar: var(--side-bar-light);

  --sidebar-width: 175px;
  --green: --green-light;
  --orange: --orange-light;
  --background: --background-light;
  --text-color: --text-color-light;
  --text: --text-light --border-color: --border-color-light;
  --card-color: --card-color-light;
  --text-side-bar: --text-side-bar-light;
  --search-bar: --search-bar-light;
}

.main-content {
  position: absolute;
  top: 10%;
  left: 25%;
  width: 80%;
  // put in column
}

.top-bar {
  position: absolute;
  top: 10%;
  left: 350%;
  display: flex;
  justify-content: space-between;

  align-items: center;
  margin-right: 150%;
}

.search-bar {
  display: flex;
  align-items: center;
  border: 1px solid #ccc;
  border-radius: 4px;
  width: 200px;
  background-color: white;
}

.search-bar input {
  border: none;
  outline: none;
  flex: 1;
  padding: 8px;
  font-size: 16px;
}

.search-bar font-awesome-icon {
  color: var(--text);

  margin-right: 8px;
  font-size: 20px;
}


.close-button {
  position: absolute;
  background: none;
  border: none;
  top: 38%;
  right: 38%;
  font-size: 20px;
  cursor: pointer;
  color: var(--text);
}

.custom-dropdown {
  /* Ajoutez ici les styles personnalisés de votre choix */
  border: 1px solid #ccc;
  border-radius: 4px;
  padding: 8px;
  width: 200px;
  font-size: 16px;
  background-color: white;
  color: #333;
  /* Ajoutez d'autres styles selon vos besoins */
}

/* Styles pour les options du dropdown */
.custom-dropdown option {
  /* Styles pour les options */
  background-color: white;
  color: #333;
}


.btn {
  background-color: var(--orange);
  margin: 5px;
  border-radius: 4px;
  border: none;
  width: 160px;
  height: 40px;

  font-family: sans-serif;
  font-weight: 600;
  /* Semi-gras */
  font-size: 14px;
  line-height: 14px;
  color: #FFFFFF;

  //hover
  &:hover {
    background-color: var(--orange-dark);
    cursor: pointer;
  }
}









.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.6);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 2; 
}

.modal-content {
  background: #fff;
  padding: 20px;
  border-radius: 10px;
  width: 450px;
  height: 225px;
  color: var(--text);
  background-color: var(--side-bar);
  border: 2px solid var(--border-color);

  flex-direction: column;
  display: flex;
  align-items: center;
  z-index: 3;

}

.modal-title {
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 10px;
}








.title {
  font-family: sans-serif;
  font-size: 1.5rem;
  font-weight: 400;
  line-height: 35px;
  margin: 1rem;
  color: var(--text);
  font-size: 28px;
}

.side-content {
  position: absolute;
  top: 5%;
  left: 20%;
  // put in column
  display: flex;
  flex-direction: column;
  align-self: left;
}

.darkbutton {
  position: absolute;
  top: 10px;
  right: 10px;
  margin: 1rem;
  z-index: 3;
}

.sidebar {
  z-index: 2; // Sidebar au-dessus de la div Main
}

.main {
  z-index: 1;
  background: var(--background);
  color: var(--text-color);
  width: calc(100% - var(--sidebar-width));
  border-left: 1px solid var(--border-color);
}

.app {
  margin: -8px;
  background: var(--background);
  display: flex;
}
</style>
  