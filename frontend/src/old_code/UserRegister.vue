<template>
  <div class="area">
    <div class="left">
      <img class="rounded-image image-logo" src="../assets/Logo_white.png" alt="Description de l'image">
    </div>
    <div class="right">
      <div class="image-container">
        <img class="centered-image" src="../assets/profile.png">
      </div>
      <h1 class="login-title"><strong>New account</strong></h1>
      <form class="login-form" @submit.prevent="register">
        <div class="form-group">
          <div class="name-container">
            <input type="text" class="First-name"  v-model="name" name="First-name" placeholder="First name" required>
            <input type="text" class="Last-name"  v-model="last_name" name="Last-name" placeholder="Last name" required>
          </div>
        </div>
        <div class="form-group">
          <input type="email" id="email"  v-model="email" name="email" placeholder="Email" required>
        </div>
        <div class="form-group">
          <input type="username" id="username"  v-model="username" name="username" placeholder="username" required>
        </div>
        <div class="form-group">
          <input type="password" id="password"  v-model="password" name="password" placeholder="Password" required>
        </div>
        <div class="button-container">
          <button type="submit" class="login-button">Create</button>
        </div>
      </form>
      <router-link to="/login" class="Sign">Sign In?</router-link>
      <button @click="registerWithGoogle" class="oauth-button google">
        <font-awesome-icon :icon="['fab', 'google']" /><span class="text-OAuth"> Register with Google</span>
      </button>
      <button @click="registerWithMicrosoft" class="oauth-button microsoft">
        <font-awesome-icon :icon="['fab', 'microsoft']" /><span class="text-OAuth"> Register with Microsoft</span>
      </button>
      <button @click="registerWithGithub" class="oauth-button github">
        <font-awesome-icon :icon="['fab', 'github']" /><span class="text-OAuth"> Register with GitHub</span>
      </button>
    </div>
  </div>
</template>

<!-- Le reste du code (script et style) reste inchangé. -->


<script>
export default {
  data() {
    return {
      last_name: '',
      name: '',
      password: '',
      email: '',
    };
  },
  methods: {
    register() {
      const url = ' https://back-area.lekmax.fr/register';
      const data = {
        username: this.username,
        password: this.password,
        email: this.email
      };
      console.log('body', data)

      fetch(url, {
        method: 'POST',
        headers: {
          'x-api-key': 'ef4f85d0-6128-11ee-8c99-0242ac120002',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
      })
        .then(response => response.json())
        .then(data => {
          console.log('Succès:', data);
          if (data.message === 'Accout registered succesfully') {
            this.$router.push('/login');  // Redirige vers la page /dashboard
            // Optionally, you can also store the token (data.token) if needed
          }
          else {
            console.error('Erreur lors de la connexion(malek):', data.message);
          }
          // Vous pouvez rediriger l'utilisateur vers une autre page ou afficher un message de succès ici.
        })
        .catch((error) => {
          console.error('Erreur:', error);
          // Affichez une pop-up ou une notification d'erreur ici.
        });
    },
    registerWithGoogle() {
      // TODO: Ajoutez votre logique pour la connexion avec Google
    },
    registerWithMicrosoft() {
      // TODO: Ajoutez votre logique pour la connexion avec Microsoft
    },
    registerWithGithub() {
      // TODO: Ajoutez votre logique pour la connexion avec GitHub
    }
  }
};
</script>
<style scoped>
.area {
  display: flex;
  height: 100vh;
}

.Sign {
  margin-top: 100px;
  text-decoration: none;
  color: #007BFF;
  font-size: 1.2rem;
}
.image-logo {
  max-width: 70%;
  max-height: 70%;
  height: auto;
  width: auto;
}
.name-container{
  display: flex;
}

.left {
  flex: 1;
  background-color: #FFC144;
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
  max-width: 100%; /* Assurez-vous que l'image ne dépasse pas de la largeur du conteneur */
  max-height: 190px; /* Ajustez la hauteur maximale de l'image selon vos besoins */
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

.First-name{
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 7px;
  display: flex;
}

.Last-name{
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 7px;
  display: flex;
}

input[type="email"] {
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