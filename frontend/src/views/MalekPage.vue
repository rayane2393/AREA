<template>
  <div class="app">
    <!-- Sidebar -->
    <Sidebar class="sidebar" />
    <div class="main">
      <DarkModeButton class="darkbutton" :initialState="isDarkMode" @switched="isDarkMode = $event" />

      <div class="side-content">
        <span class="title">Areas</span>
        <router-link to="/createArea">
          <button class="btn">Add Areas</button>
        </router-link>


      </div>


      <div class="main-content">
        <div class="search-bar">
          <input type="text" placeholder=" 🔍 Search Areas..." v-model="searchText" />
        </div>
        <div class="sort-dropdown">
          <select>
            <option value="last_updated">Sort by last updated</option>
            <option value="last_created">Sort by last created</option>
            <option value="az">Sort by name (A-Z)</option>
            <option value="za">Sort by name (Z-A)</option>
          </select>
        </div>
        <div class="card-list" style="max-height: 820px; overflow-y: scroll;">
          <CardArea v-for="card in sortedCardList" :key="card.id" :id="card.id" :title="card.title"
            :created_at="card.created_at" :updated_at="card.updated_at" :status="card.status" />
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import CardArea from '../component/CardArea.vue'; // Assurez-vous d'importer le composant Card

export default {
  components: {
    Sidebar,
    DarkModeButton,
    CardArea,
  },
  data() {
    return {
      sortOption: 'az',
      searchText: '',
      cardList: [
        {
          id: 1,
          title: "Lemlist notification to Slack (Kaizzen)",
          updated_at: "2023-10-27T14:30:00", // Exemple de timestamp ISO
          created_at: "2023-09-26T14:30:00",
          status: "Active",
        },
        {
          id: 2,
          title: "Airtable",
          updated_at: "2023-10-10T12:15:00", // Exemple de timestamp ISO
          created_at: "2023-10-24T14:30:00",
          status: "Inactive",
        },
        {
          id: 3,
          title: "Engage LinkedIn Post Commenters",
          updated_at: "2023-09-10T12:15:00", // Exemple de timestamp ISO
          created_at: "2023-08-08T14:30:00",
          status: "Inactive",
        },
        {
          id: 4,
          title: "Indeed Job intent to Airtable & Lemlist",
          updated_at: "2023-10-16T12:15:00", // Exemple de timestamp ISO
          created_at: "2023-10-08T14:30:00",
          status: "Inactive",
        },
        {
          id: 5,
          title: "SalesNav search OR Indeed People to Clay",
          updated_at: "2023-08-08T14:30:00", // Exemple de timestamp ISO
          created_at: "2023-08-08T14:30:00",
          status: "Inactive",
        },
        {
          id: 6,
          title: "Hyper-personalized message with OpenAI",
          updated_at: "2023-09-27T14:30:00", // Exemple de timestamp ISO
          created_at: "2023-09-27T14:30:00",
          status: "Inactive",
        },
        {
          id: 7,
          title: "Scraper Welcome To The Jungle (Template)",
          updated_at: "2023-10-30T14:30:00", // Exemple de timestamp ISO
          created_at: "2023-08-30T14:30:00",
          status: "Active",
        },
        {
          id: 8,
          title: "Engage LinkedIn Post Commenters with OpenAI",
          updated_at: "2023-10-13T14:30:00", // Exemple de timestamp ISO
          created_at: "2023-10-13T14:30:00",
          status: "Inactive",
        },
        {
          id: 9,
          title: "Social Warming LeadGPT",
          updated_at: "2023-08-08T14:30:00", // Exemple de timestamp ISO
          created_at: "2023-08-08T14:30:00",
          status: "Inactive",
        },
        {
          id: 10,
          title: "Test workflow Lemlist to G Sheet",
          updated_at: "2023-09-14T14:30:00", // Exemple de timestamp ISO
          created_at: "2023-09-14T14:30:00",
          status: "Inactive",
        },
        {
          id: 11,
          title: "Target companies that raised funds",
          updated_at: "2023-10-11T14:30:00", // Exemple de timestamp ISO
          created_at: "2023-10-11T14:30:00",
          status: "Inactive",
        },
        {
          id: 12,
          title: "My workflow 2",
          updated_at: "2023-10-11T14:30:00", // Exemple de timestamp ISO
          created_at: "2023-10-11T14:30:00",
          status: "Active",
        },
      ],
    };
  },
  methods: {
    sortCards() {
      if (this.sortOption === 'az') {
        this.cardList.sort((a, b) => a.title.localeCompare(b.title));
      } else if (this.sortOption === 'za') {
        this.cardList.sort((a, b) => b.title.localeCompare(a.title));
      } else if (this.sortOption === 'last_updated') {
        this.cardList.sort((a, b) => new Date(b.updated_at) - new Date(a.updated_at));
      } else if (this.sortOption === 'last_created') {
        this.cardList.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
      }
    },
  },
  watch: {

    sortOption() {

      this.sortCards(); // Appelez la fonction pour trier la liste
      // Mise à jour de la liste triée lorsque l'option de tri change
      this.sortedCardList;
    },
  },
  computed: {
    sortedCardList() {
      // Ajoutez ici la logique pour trier votre liste de cartes
      // en fonction de la valeur de sortOption
      if (this.sortOption === 'az') {
        return this.cardList.sort((a, b) => a.title.localeCompare(b.title));
      } else if (this.sortOption === 'za') {
        return this.cardList.sort((a, b) => b.title.localeCompare(a.title));
      } else if (this.sortOption === 'last_updated') {
        return this.cardList.sort((a, b) => new Date(b.updated_at) - new Date(a.updated_at));
      } else if (this.sortOption === 'last_created') {
        return this.cardList.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
      }
      // Par défaut, retournez la liste de cartes non triée
      return this.cardList;
    },
    filteredCardList() {
      return this.cardList.filter(card => card.title.toLowerCase().includes(this.searchText.toLowerCase()));
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
    root.style.setProperty('--card-color', 'var(--card-color-light)');
    root.style.setProperty('--text-side-bar', 'var(--text-side-bar-light)');
    root.style.setProperty('--search-bar', 'var(--search-bar-light)');
    root.style.setProperty('--hover-color', 'var(--hover-color-light)');
    root.style.setProperty('--hover-text-color', 'var(--hover-text-color-light)');
  }
});
</script>






<style scoped>
.sort-dropdown {
  position: absolute;
  top: 0px;
  left: 300px;
  display: inline-block;
  margin-right: 20px; /* Ajoutez une marge à droite pour l'espacement */
  font-family: sans-serif;
  width: 150px;
  height: 24px;
  font-size: 14px;
  text-align: center;
  color: #333; /* Couleur du texte */
  background-color: #fff; /* Couleur de fond */
  border: 1px solid #ccc; /* Bordure */
  border-radius: 4px; /* Coins arrondis */
  padding: 5px 10px; /* Espacement intérieur */
}

/* Style pour les options du menu déroulant */
.sort-dropdown select {
  appearance: none; /* Supprime l'apparence par défaut du menu déroulant (pour Chrome) */
  border: none; /* Supprime la bordure du menu déroulant */
  background-color: transparent; /* Fond transparent pour les options */
  cursor: pointer; /* Curseur indiquant la sélection */
}

/* Style au survol */
.sort-dropdown select:hover {
  background-color: #f5f5f5;
}

/* Style pour les options sélectionnées */
.sort-dropdown select option:checked {
  background-color: #007bff; /* Couleur de fond de l'option sélectionnée */
  color: #fff; /* Couleur du texte de l'option sélectionnée */
}
</style>




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

  --border-color-dark: #736B5E;
  --border-color-light: #555555;

  --card-color-dark: #252829;
  --card-color-light: --side-bar-light;

  --text-side-bar-dark: #B8B2A8;
  --text-side-bar-light: #7D7D87;

  --search-bar-dark: #141516;
  --search-bar-light: --side-bar-light;

  --sidebar-width: 175px;
  --green: --green-light;
  --orange: --orange-light;
  --background: --background-light;
  --side-bar: var(--side-bar-light) --text-color: --text-color-light;
  --border-color: --border-color-light;
  --card-color: --card-color-light;
  --text-side-bar: --text-side-bar-light;
  --search-bar: --search-bar-light;
}

.main-content {
  margin: 40px;
  margin-left: 200px;
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
  top: 50%;
  left: 0%;
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
