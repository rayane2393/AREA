<template>
  <div class="app">
    <!-- Sidebar -->
    <Sidebar class="sidebar" />
    <div class="main">
      <DarkModeButton class="darkbutton" :initialState="isDarkMode" @switched="isDarkMode = $event" />
      <div class="side-content">
        <h1>Save Credentials</h1>
        
        <form @submit.prevent="create_credentials">
          <span>or use your email for registration</span>
          <input type="email" id="register-email" v-model="email" required placeholder="Email" />
          <input type="password" id="register-password" v-model="password" placeholder="Password" required>
          <input type="password" id="register-username" v-model="client_id" required placeholder="Client Id" />
          <input type="password" id="register-username" v-model="client_secret" required placeholder="Client Secret" />

          <button class="btn">save</button>
        </form>
        <router-link to="/list-credentials">
          <span>Done</span>
        </router-link>
        <!-- test save and send to backend-->
      </div>

    </div>
  </div>
</template>
  
<script>
export default {
  data() {
    return {
      email: '',
      password: '',
      client_secret: '',
      client_id: '',
    };
  },
  components: {
    Sidebar,
    DarkModeButton,
  },
  created() {
    // TODO: get list of area
  },
  methods: {
    create_credentials() {
      console.log(this.email, this.password, this.client_id, this.client_secret);
      // TODO: create credentials, send them to the server
    },

    test_credentials() {
      console.log(  )
      // TODO: test credentials, send them to the server
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
  --text-color: --text-color-light;
  --border-color: --border-color-light;
  --card-color: --card-color-light;
  --text-side-bar: --text-side-bar-light;
  --search-bar: --search-bar-light;
}

.side-content {
  top: 50%;
  color: var
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
  