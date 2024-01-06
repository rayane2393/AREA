<template>
  <div class="dashboard">
    <h1>Welcome to the Dashboard</h1>

    <!-- Barre de recherche -->
    <input v-model="searchTerm" placeholder="Rechercher un workflow..." />

    <!-- Bouton pour ajouter un workflow -->
    <button class="btn btn-primary" @click="addWorkflow">Ajouter un workflow</button>

    <!-- Liste de workflows -->
    <div v-for="workflow in filteredWorkflows" :key="workflow.id" class="workflow-box">
      <div class="workflow-info">

        <div class="workflow-name">
          <h3>{{ workflow.name }}</h3>
        </div>
        <p>Date de création: {{ workflow.creationDate }}</p>
      </div>

      <!-- Toggle switch pour activer/désactiver le workflow -->
      <label class="switch">
        <input type="checkbox" v-model="workflow.active">
        <span class="slider round"></span>
      </label>

      <!-- Icône de paramètres -->
      <div class="settings">
        <div class="kebab-menu" @click="showMenu(workflow.id)">
          <div class="dot"></div>
          <div class="dot"></div>
          <div class="dot"></div>
        </div>

        <!-- Menu déroulant -->
        <div v-if="openedMenu === workflow.id" class="dropdown-menu">
          <a href="#" @click="deleteWorkflow(workflow.id)">Supprimer</a>
          <a href="#" @click="duplicateWorkflow(workflow.id)">Dupliquer</a>
          <a href="#" @click="openWorkflow(workflow.id)">Ouvrir</a>
          <a href="#" @click="shareWorkflow(workflow.id)">Partager</a>
        </div>
      </div>
    </div>
  </div>
</template>


<script>
export default {
  name: 'UserDashboard',
  data() {
    return {
      searchTerm: '',
      openedMenu: null,
      workflows: [
        { id: 1, name: 'Send mail auto', creationDate: '2023-01-01', active: true },
        { id: 2, name: 'Scrapping KMS', creationDate: '2023-01-01', active: true },
        { id: 3, name: 'Syncro Agenda', creationDate: '2023-02-01', active: true },
        { id: 4, name: 'Restart Server', creationDate: '2023-01-01', active: false },
        // Ajoutez plus de fausses données si nécessaire...
      ],
    };
  },
  computed: {
    filteredWorkflows() {
      return this.workflows.filter(wf => wf.name.includes(this.searchTerm));
    },
  },
  methods: {
    addWorkflow() {
      // Implémentez la logique pour ajouter un nouveau workflow
    },
    showMenu(id) {
      this.openedMenu = id;
    },
    deleteWorkflow(id) {
      console.log('delete workflow', id);
      // Implémentez la logique pour supprimer le workflow
    },
    duplicateWorkflow(id) {
      console.log('duplicate workflow', id);
      // Implémentez la logique pour dupliquer le workflow
    },
    openWorkflow(id) {
      console.log('open workflow', id);
      // Implémentez la logique pour ouvrir le workflow
    },
    shareWorkflow(id) {
      console.log('share workflow', id);
      // Implémentez la logique pour partager le workflow
    },
  },
};
</script>


<style scoped>
.dashboard {
  background-color: #2b2b2b;
  color: #e8e8e8;
  padding: 20px;
}

.workflow-name {
    font-family: "Open Sans", sans-serif;
    font-size: 14px;
    line-height: 18.9px;
    text-align: start;
    letter-spacing: normal;
    color: #555555;
    background-color: #ffffff;
}


/* ... (le reste de votre CSS est inchangé) ... */

.workflow-info {
  flex: 1;
  padding-right: 20px;
}

.workflow-box {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border: 1px solid #444;
  padding: 15px;
  margin-top: 10px;
  background-color: #333;
  border-radius: 6px;
  transition: background-color 0.3s;
}

.workflow-box:hover {
  background-color: #3b3b3b;
}


input {
  padding: 10px;
  border-radius: 6px;
  border: 1px solid #555;
  background-color: #222;
  color: #e8e8e8;
  margin-bottom: 10px;
}

.btn {
  padding: 10px 15px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.btn-primary {
  background-color: #007BFF;
  color: white;
}

.btn-primary:hover {
  background-color: #0056b3;
}

.kebab-menu {
    display: flex;
    flex-direction: column;
    align-items: center;
    cursor: pointer;
    gap: 4px;
}

.dot {
    width: 5px;
    height: 5px;
    background-color: black;
    border-radius: 50%;
}

.dropdown-menu {
    display: none;
    border: 1px solid #ccc;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    margin-top: 5px;
    background-color: #fff;
}

.dropdown-menu a {
    display: block;
    padding: 8px 16px;
    text-decoration: none;
    color: black;
    border-bottom: 1px solid #eee;
}

.dropdown-menu a:last-child {
    border-bottom: none;
}

.dropdown-menu a:hover {
    background-color: #f5f5f5;
}

/* CSS pour afficher le menu lorsque vous cliquez sur les trois points (si vous n'utilisez pas Vue.js) */
.kebab-menu:hover + .dropdown-menu {
    display: block;
}


.btn-secondary {
  background-color: #666;
  color: white;
  margin: 5px;
}

.btn-secondary:hover {
  background-color: #555;
}

.switch {
  position: relative;
  display: inline-block;
  width: 60px;
  height: 34px;
}

.switch input {
  display: none;
}

.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #666;
  transition: 0.4s;
}

.slider:before {
  position: absolute;
  content: "";
  height: 26px;
  width: 26px;
  left: 4px;
  bottom: 4px;
  background-color: #444;
  transition: 0.4s;
}

input:checked+.slider {
  background-color: #007BFF;
}

input:checked+.slider:before {
  transform: translateX(26px);
}

.slider.round {
  border-radius: 34px;
}

.slider.round:before {
  border-radius: 50%;
}

.dropdown-menu {
  background-color: #333;
  padding: 10px;
  border-radius: 6px;
  margin-top: 10px;
  box-shadow: 0px 0px 15px rgba(0, 0, 0, 0.3);
}

.settings:hover .dropdown-menu {
  display: block;
}
</style>
