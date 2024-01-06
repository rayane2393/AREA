<script setup>
import { globalVariable } from '../global.js';
const callback = (response) => {
  console.log()
  const url = 'https://back-area.lekmax.fr/auth/login/google';
      const headers = {
        'Content-Type': 'application/json',
        'x-api-key': 'ef4f85d0-6128-11ee-8c99-0242ac120002'
      };
      const body = response;
      fetch(url, {
        method: 'POST',
        headers: headers,
        body: body
      })
      .then(response => response.json())
      .then(data => {
          // Vérifiez la réponse pour vous assurer que l'utilisateur est connecté avec succès
          // Cela dépend de la structure de votre réponse. 
          // Par exemple, si votre API renvoie { success: true } en cas de succès :
          if (data.message === 'Login successful') {
            this.$router.push({
              name: 'HomePage'
            });
          }
          else {
            console.error('Erreur lors de la connexion(malek):', data.message);
            // Vous pouvez également afficher un message d'erreur à l'utilisateur ici
          }
        })
        .catch(error => {
          console.error('Erreur lors de la connexion:(malek)', error);
        });
}
</script>
<template>
  <div class="area">
    <div class="left">
      <h1 class="title">Welcome to</h1>
      <img class="rounded-image image-logo" src="../assets/Logo_white.png" alt="Description de l'image">
    </div>
    <div class="right">
      <div class="image-container">
        <img class="centered-image" src="../assets/profile.png">
      </div>
      <h1 class="login-title"><strong>Connexion</strong></h1>
      <form class="login-form" @submit.prevent="login">
        <div class="form-group">
          <input type="username" id="username" v-model="username" required placeholder="Usernames" />
        </div>
        <div class="form-group">
          <input type="password" id="password" v-model="password" placeholder="Password" required>
        </div>
        <div class="button-container">
          <button class="login-button">Sign In</button>
        </div>
        <div class="links-container">
          <br>
          <router-link to="/forgot-password" class="centered-link">Forgotten password</router-link>
          <br>
          <br>
          <router-link to="/register" href="#" class="centered-link">Create an account</router-link>
        </div>
      </form>
        <GoogleLogin :callback="callback"/>

      <!--<button @click="loginWithGoogle" class="oauth-button google">
        <font-awesome-icon :icon="['fab', 'google']" /><span class="text-OAuth"> Login with Google</span>
      </button>-->
      <button @click="loginWithGithub" class="oauth-button github">
        <font-awesome-icon :icon="['fab', 'github']" /><span class="text-OAuth"> Login with GitHub</span>
      </button>
    </div>
  </div>
</template>

<script>
import * as Msal from 'msal'; 
export default {

  data() {
    return {
      username: '',
      password: ''
    };
  },
  methods: {
    login() {
      const url = ' https://back-area.lekmax.fr/login';
      const headers = {
        'Content-Type': 'application/json',
        'x-api-key': 'ef4f85d0-6128-11ee-8c99-0242ac120002'
      };
      const body = JSON.stringify({
        username: this.username,
        password: this.password
      });
      fetch(url, {
        method: 'POST',
        headers: headers,
        body: body
      })
        .then(response => response.json())
        .then(data => { 
          console.log('Réponse du serveur:()', globalVariable.value);
          // Vérifiez la réponse pour vous assurer que l'utilisateur est connecté avec succès
          // Cela dépend de la structure de votre réponse. 
          // Par exemple, si votre API renvoie { success: true } en cas de succès :
          if (data.message === 'Login successful') {
            globalVariable.value = data.token;
            this.$router.push({
              name: 'HomePage',
            });
          }
          else {
            alert("Mot de passe ou/et e-mail incorrect")
            console.error('Erreur lors de la connexion(malek):', data.message);
            // Vous pouvez également afficher un message d'erreur à l'utilisateur ici
          }
        })
        .catch(error => {
          console.error('Erreur lors de la connexion:(malek)', error);
        });
    },
    loginWithMicrosoft() {

      const msalConfig = {
        auth: {
          clientId: 'YOUR_MICROSOFT_CLIENT_ID',
          redirectUri: 'https://back-area.lekmax.fr/auth/callback'
        }
      };

      const msalInstance = new Msal.UserAgentApplication(msalConfig);

      msalInstance.loginPopup().then(response => {
        const token = response.idToken.rawIdToken;
        console.log('Token:', token)
        // Envoyez ce token à votre backend pour le vérifier et établir une session pour l'utilisateur
        // TODO: Ajoutez votre logique pour envoyer le token au backend
      }).catch(error => {
        console.error("Erreur lors de la connexion avec Microsoft:", error);
      });
    },
    loginWithGithub() {
      const clientId = '6c5a0ae6296fe2c356eb';
      const redirectUri = 'https://back-area.lekmax.fr/auth/callback';
      const scope = 'user';  // Demande l'accès aux informations de base de l'utilisateur
      window.location.href = `https://github.com/login/oauth/authorize?client_id=${clientId}&redirect_uri=${redirectUri}&scope=${scope}`;
    }
  }
};
</script>

<style scoped>
.area {
  display: flex;
  height: 100vh;
}

.title {
  color: white;
  font-size: 6rem;
}

.image-logo {
  max-width: 70%;
  max-height: 70%;
  height: auto;
  width: auto;
}

.left {
  flex: 1;
  background-color: #CB6CE6;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  margin-top: 2%;
  margin-bottom: 2%;
  margin-bottom: 2%;
  margin: 2%;
}

.rounded-image {
  border-radius: 40px;
}

.right {
  flex: 1;
  background-color: white;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.form-group {
  margin-bottom: 20px;
}

.login-form {
  width: 60%;
  max-width: 400px;
  padding: 0px;
}

.centered-image {
  max-width: 100%;
  /* Assurez-vous que l'image ne dépasse pas de la largeur du conteneur */
  max-height: 190px;
  /* Ajustez la hauteur maximale de l'image selon vos besoins */
}

.login-title {
  color: #808080;
  font-weight: bold;
  font-size: 4rem;
}

label {
  display: block;
  font-weight: bold;
}

input[type="password"] {
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 7px;
}

input[type="username"] {
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 7px;
}

.button-container {
  text-align: center;
  margin-top: 10px;
}

.login-button {
  background-color: #c1c2c3;
  color: white;
  border: none;
  border-radius: 5px;
  padding: 10px 20px;
  cursor: pointer;
  font-size: 1.1rem;
}

.login-button:hover {
  background-color: #858686;
}

.links-container {
  margin-top: 10px;
  text-align: center;
  /* Centrer le texte dans la section des liens */
}

.centered-link {
  text-decoration: none;
  color: #007BFF;
  margin: 0 10px;
  font-size: 1.2rem;
}

.links-container .separator {
  font-weight: bold;
}


.text-OAuth {
  margin-left: 10px;
}

.oauth-button {
  padding: 10px 30px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin-top: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
}

.oauth-button .fa-icon {
  margin-right: 5px;
}

.google {
  background-color: #db4437;
  color: #fff;
}

.microsoft {
  background-color: #00a1f1;
  color: #fff;
}

.github {
  background-color: #333;
  color: #fff;
}
</style>