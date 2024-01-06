<template>
  <div>
    <PopUp :show="showPopup" :message="Message" @close="closePopup" :type="typeMessage" />
    <div class="container" :class="{ active: isRegistering }">
      <div class="form-container sign-up" :class="{ active: isRegistering }" v-if="isRegistering">
        <h1>Create Account</h1>
        <div class="social-icons">
          <button @click="loginWithGoogle" class="icon">
            <font-awesome-icon :icon="['fab', 'google']" />
          </button>
          <button @click="loginWithGithub" class="icon box">
            <font-awesome-icon :icon="['fab', 'github']" />
          </button>
          <button @click="loginWithDiscord" class="icon box">
            <font-awesome-icon :icon="['fab', 'discord']" />
          </button>
        </div>
        <form @submit.prevent="register">
          <span>or use your email for registration</span>
          <input type="username" id="register-username" v-model="register_username" required placeholder="Username" />
          <input type="email" id="register-email" v-model="register_email" required placeholder="Email" />
          <input type="password" id="register-password" v-model="register_password" placeholder="Password" required>

          <button class="btn">Sign Up</button>
        </form>
      </div>
      <div class="form-container sign-in" :class="{ active: !isRegistering }" v-if="!isRegistering">
        <h1>Sign In</h1>
        <div class="social-icons">
          <button @click="loginWithGoogle" class="icon">
            <font-awesome-icon :icon="['fab', 'google']" />
          </button>
          <button @click="loginWithGithub" class="icon box">
            <font-awesome-icon :icon="['fab', 'github']" />
          </button>
          <button @click="loginWithDiscord" class="icon box">
            <font-awesome-icon :icon="['fab', 'discord']" />
          </button>
        </div>
        <form @submit.prevent="login">
          <span>or use your email password</span>
          <input type="email" id="login-email" v-model="login_email" required placeholder="Email" />
          <input type="password" id="login-password" v-model="login_password" placeholder="Password" required>

          <button class="btn">Sign In</button>
          <router-link to="/forgot-password" class="centered-link">Forgotten password</router-link>
        </form>
      </div>

      <div class="toggle-container">
        <div class="toggle" :class="{ slide: isRegistering }">
          <div class="toggle-panel toggle-left">
            <img class="rounded-image image-logo" src="../assets/Logo_copy.png" alt="Description de l'image">
            <h1>Welcome Back!</h1>
            <p>Enter your personal details to use all site features</p>
            <button class="hidden btn" id="login" @click="toggleRegistration">Sign In</button>
          </div>
          <div class="toggle-panel toggle-right">
            <img class="rounded-image image-logo" src="../assets/Logo_copy.png" alt="Description de l'image">
            <h1>Hello, Friend!</h1>
            <p>Register with your personal details to use all site features</p>
            <button class="hidden btn" id="register" @click="toggleRegistration">Sign Up</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faGoogle, faGithub, faDiscord } from "@fortawesome/free-brands-svg-icons";
import PopUp from '../component/PopUp.vue';
import { API_KEY, API_URL, HOST_URL, GOOGLE_CLIENT_ID, GITHUB_CLIENT_ID } from '../config.js';
import { globalVariable } from '../global.js';

library.add(faGoogle, faGithub, faDiscord);

export default {
  components: {
    FontAwesomeIcon,
    PopUp,
  },
  data() {
    return {
      isRegistering: false,
      register_username: '',
      register_password: '',
      register_email: '',
      login_email: '',
      login_password: '',
      Message: '',
      showPopup: false,
      typeMessage: '',
    };
  },
  methods: {
    showMessage(message, type) {
      this.Message = message;
      this.showPopup = true;
      this.typeMessage = type;
      setTimeout(() => {
        this.closePopup();
      }, 3000);
    },
    closePopup() {
      this.showPopup = false;
      this.Message = '';
      this.typeMessage = '';
    },
    login() {
      const url = API_URL + '/login';
      const headers = {
        'Content-Type': 'application/json',
        'x-api-key': API_KEY,
      };
      // console.log("apikey login", API_KEY)
      const body = JSON.stringify({
        username: this.login_email.toLowerCase(),
        password: this.login_password
      });
      fetch(url, {
        method: 'POST',
        headers: headers,
        body: body
      })
        .then(response => response.json())
        .then(data => {
          if (data.message === 'Login successful') {
            globalVariable.value = data.token;
            localStorage.setItem('user', 'authenticated');
            localStorage.setItem('token', data.token);
            this.$router.push({
              name: 'HomePage',
            });
          }
          else {
            this.showMessage(data.message + "!", "error");
          }
        })
        .catch(error => {
          console.error('Erreur lors de la connexion', error);
          this.showMessage('Une erreur s\'est produite. Veuillez réessayer.', "error");
        });
    },
    register() {
      const url = API_URL + '/register';
      const data = {
        username: this.register_username,
        password: this.register_password,
        email: this.register_email
      };
      console.log('body', data)
      console.log("apikey register", API_KEY)
      fetch(url, {
        method: 'POST',
        headers: {
          'x-api-key': API_KEY,
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
      })
        .then(response => response.json())
        .then(data => {
          console.log('Succès:', data);
          if (data.message === 'Accout registered succesfully') {
            this.isRegistering = false;
            this.showMessage('Votre compte a bien été créé. Vous pouvez vous connecter.', "info");
          }
          else {
            this.showMessage('Une erreur s\'est produite. Veuillez réessayer.', "error");
          }
        })
        .catch((error) => {
          console.error('Erreur:', error);
          this.showMessage('Catch Error. Veuillez réessayer.', "error");

        });
    },
    toggleRegistration() {
      this.isRegistering = !this.isRegistering;
    },
    loginWithGoogle() {
      const clientId = GOOGLE_CLIENT_ID;
      const redirectUri = HOST_URL + '/callbackgoogle';
      const scope = 'https://www.googleapis.com/auth/userinfo.profile';
      const authUrl = `https://accounts.google.com/o/oauth2/auth?client_id=${clientId}&redirect_uri=${redirectUri}&scope=${scope}&response_type=code`;
      window.location.href = authUrl;
    },
    loginWithGithub() {
      const clientID = GITHUB_CLIENT_ID;
      const redirectURI = HOST_URL + '/callbackgithub';
      window.location.href = `https://github.com/login/oauth/authorize?client_id=${clientID}&redirect_uri=${redirectURI}&scope=user,public_repo`;
    },
    loginWithDiscord() {
      this.showMessage('Fonctionnalité en cours de développement.', "error");
      // 1168952746733473902
      // 92048d5151b31727bddb86720a85afde3298cbe30235645bc768942c54560f36
      // 9RgUgV_gFqXSuWtB88yAmTpa3cmpDFU-
      /*
      const clientID = '1168952746733473902';
      const redirectURI = 'http://localhost:8080/login'; // Correspond à la route DiscordCallback
      const scope = 'identify'; // Ajoutez les scopes nécessaires

      const authURL = `https://discord.com/api/oauth2/authorize?client_id=${clientID}&redirect_uri=${redirectURI}&response_type=code&scope=${scope}`;

      window.location.href = authURL;
      */
    }

  },
};
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@300;400;500;600;700&display=swap');



* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  font-family: 'Montserrat', sans-serif;
}

body {
  background-color: #c9d6ff;
  background: linear-gradient(to right, #e2e2e2, #c9d6ff);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  height: 100vh;
}

.container {
  margin-top: 10%;
  margin-left: 22%;
  background-color: #fff;
  border-radius: 30px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.35);
  position: relative;
  overflow: hidden;
  width: 57%;
  max-width: 100%;
  min-height: 580px;
}

.container p {
  font-size: 14px;
  line-height: 20px;
  letter-spacing: 0.3px;
  margin: 20px 0;
}

.image-logo {
  max-width: 70%;
  max-height: 70%;
  height: auto;
  width: auto;
}

.container span {
  font-size: 12px;
}

.container a {
  color: #333;
  font-size: 13px;
  text-decoration: none;
  margin: 15px 0 10px;
}

.container .btn {
  background-color: #CB6CE6;
  color: #fff;
  font-size: 12px;
  padding: 10px 45px;
  border: 1px solid transparent;
  border-radius: 8px;
  font-weight: 600;
  letter-spacing: 0.5px;
  text-transform: uppercase;
  margin-top: 10px;
  cursor: pointer;
}

.container button.hidden {
  background-color: transparent;
  border-color: #fff;
}

.toggle-container.slide {
  animation: slideAnimation 0.6s forwards;
}

@keyframes slideAnimation {
  0% {
    transform: translateX(0);
  }

  100% {
    transform: translateX(-100%);
  }
}


.container  form {
  background-color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
}

.container input {
  background-color: #eee;
  border: none;
  margin: 8px 0;
  padding: 10px 15px;
  font-size: 13px;
  border-radius: 8px;
  width: 100%;
  outline: none;
}

.form-container {
  position: absolute;
  top: 20%;
  padding: 0 40px;
  height: 100%;
  transition: all 0.6s ease-in-out;
}

.sign-in {
  left: 0;
  width: 50%;
  opacity: 1;
  z-index: 1;
  transform: translateX(0);
  animation: move 0.6s;
}

.container.active .sign-in {
  transform: translateX(100%);
  opacity: 0;
  z-index: 5;
  animation: move 0.6s;
}

.container.active .sign-up {
  transform: translateX(100%);
  opacity: 1;
  z-index: 5;
  animation: move-t 0.6s;
}

.sign-up {
  left: 0;
  width: 50%;
  opacity: 0;
  z-index: 1;
  transform: translateX(0);
  animation: move 0.6s;
}

@keyframes move {
  0% {
    transform: translateX(0);
  }

  100% {
    transform: translateX(100%);
  }
}

@keyframes move-t {
  0% {
    transform: translateX(100%);
  }

  100% {
    transform: translateX(0);
  }
}


.social-icons {
  margin: 20px 0;
}

.icon {
  border: 1px solid #ccc;
  border-radius: 20%;
  display: inline-flex;
  justify-content: center;
  align-items: center;
  margin: 0 3px;
  width: 40px;
  height: 40px;
}

.social-icons a {
  border: 1px solid #ccc;
  border-radius: 20%;
  display: inline-flex;
  justify-content: center;
  align-items: center;
  margin: 0 3px;
  width: 40px;
  height: 40px;
}

.toggle-container {
  position: absolute;
  top: 0;
  left: 50%;
  width: 50%;
  height: 100%;
  overflow: hidden;
  transition: all 0.6s ease-in-out;
  border-radius: 150px 0 0 100px;
  z-index: 1000;
}

.container.active .toggle-container {
  transform: translateX(-100%);
  border-radius: 0 150px 100px 0;
}

.toggle {
  background-color: #CB6CE6;
  height: 100%;
  background: linear-gradient(to right, #5c6bc0, #CB6CE6);
  color: #fff;
  position: relative;
  left: -100%;
  height: 100%;
  width: 200%;
  transform: translateX(0);
  transition: all 0.6s ease-in-out;
}


.container.active .toggle {
  transform: translateX(50%);
}

.toggle-panel {
  position: absolute;
  width: 50%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  padding: 0 30px;
  text-align: center;
  top: 0;
  transform: translateX(0);
  transition: all 0.6s ease-in-out;
}

.toggle-left {
  transform: translateX(-200%);
}

.container.active .toggle-left {
  transform: translateX(0);
}

.toggle-right {
  right: 0;
  transform: translateX(0);
}

.container.active .toggle-right {
  transform: translateX(200%);
}

</style>
