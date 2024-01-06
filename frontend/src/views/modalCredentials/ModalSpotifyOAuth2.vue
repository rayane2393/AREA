<template>
  <div class="modal">
    <div class="modal-content">
      
    <img style="width: 30px; height: 30px;" src="../../assets/spotify.svg" alt="Logo Spotify" />
    <h2 class="modal-title">Spotify OAuth2 API</h2>
      <div class="form-group">
        <label for="cred_name">Credentials name</label>
        <input type="text" id="cred_name" v-model="cred_name" required />
      </div>
      <div class="form-group">
        <label for="client_id">Client ID:<span style="color: red; font-size:  20px;">*</span></label>
        <input type="text" id="client_id" v-model="clientId" required />
      </div>
      <div class="form-group">
        <label for="client_secret">Client Secret: <span style="color: red; font-size:  20px;">*</span></label>
        <input type="text" id="client_secret" v-model="clientSecret" required />
      </div>

      <div class="form-group">
        <label for="oauth_call_url">OAuth Redirect URL:</label>
        <input type="text" id="oauth_call_url" v-model="oauthCallbackUrl" />
        <button @click="copyToClipboard">Copy to Clipboard</button>
      </div>
      <br>
      <button @click="connectMyAccount" v-if="clientId && clientSecret">Connect my Account</button>
      <br>
      <br>
      <button @click="closeModal">Fermer</button>
    </div>
  </div>
</template>

<script>
import openPopupAuthUrl from './popup.js';

export default {
  data() {
    return {
      cred_name: "Spotify OAuth2 Account", //TODO: put  number if is already exist
      clientId: "b38d12624e954e1a9771bfe1c88095d6",
      clientSecret: "759d8f0902df4ddcbca57077ed6cf3f9",
      oauthCallbackUrl: "http://localhost:8080/list-credentials",
    };
  },

  methods: {
    closeModal() {
      this.$emit('close'); // Émet un événement pour fermer la modal
      window.location.reload(); // Recharge la page pour afficher les nouvelles credentials
    },
    copyToClipboard() {
      if (this.clientId && this.clientSecret) {
        const textToCopy = this.oauthRedirectUrl;
        const textArea = document.createElement("textarea");
        textArea.value = textToCopy;
        document.body.appendChild(textArea);
        textArea.select();
        document.execCommand("copy");
        document.body.removeChild(textArea);
      }
    },
    connectMyAccount() {
      if (this.clientId && this.clientSecret) {
        const authorizationUrl = `https://accounts.spotify.com/authorize?client_id=${this.clientId}&scope=user-read-private, user-read-email, playlist-read-private, playlist-read-collaborative, playlist-modify-private, playlist-modify-public, user-follow-modify, user-follow-read, user-library-modify, user-library-read, user-read-playback-state, user-modify-playback-state, user-read-currently-playing, user-read-playback-position, user-top-read, user-read-recently-played&redirect_uri=${this.oauthCallbackUrl}&response_type=code`;
        console.log("URL d'authentification : ", authorizationUrl);
        openPopupAuthUrl(authorizationUrl)
          .then((code) => {
            console.log("Code d'authentification spotify : ", code);
            //TODO: sent it  to back end
          })
          .catch((error) => {
            console.error("Erreur : ", error);
          });
      }
    },

  },
};
</script>


<style scoped>
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
}

.modal-content {
  background: #ffffff;
  padding: 20px;
  border-radius: 10px;
  width: 40%;
  height: auto;
  color: #333;
  background-color: #2b2d2e;
  border: 2px solid #555555;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.modal-title {
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 20px;
  color: #ebebeb;
}

form {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.form-group {
  margin: 10px 0;
  width: 80%;
}

label {
  font-size: 1rem;
  margin-bottom: 5px;
  color: #ebebeb;
}

input {
  width: 100%;
  padding: 10px;
  border: 1px solid #555555;
  border-radius: 5px;
  background-color: #333333;
  color: #ebebeb;
  font-size: 1rem;
}

input[readonly] {
  background-color: #2b2d2e;
  color: #ebebeb;
}

button {
  background-color: #0fa552;
  color: #ffffff;
  padding: 10px 20px;
  border: none;
  border-radius: 5px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s;
}

button:hover {
  background-color: #13ce66;
}

.btn {
  margin: 20px 0;
}
</style>
