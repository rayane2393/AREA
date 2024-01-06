<template>
  <div  v-if="showCredentialsModal" class="modal">
    <div class="modal-content">
    <button class="close-button" @click="closeModal">×</button>
    <img style="width: 30px; height: 30px;" src="../../assets/github.svg" alt="Logo Gitub" />
      <h2 class="modal-title">GitHub API</h2>
      <div class="form-group">
        <label for="cred_name">Credentials name</label>
        <input type="text" id="cred_name" v-model="cred_name" required />
      </div>
      <div class="form-group">
        <label for="user">User:<span style="color: red; font-size:  20px;">*</span></label>
        <input type="text" id="user" v-model="user" required />
      </div>
      <div class="form-group">
        <label for="token">Access Token: <span style="color: red; font-size:  20px;">*</span></label>
        <input type="text" id="token" v-model="token" required />
      </div>
      <div class="form-group">
        <label for="oauth_redirect_url">Github Server</label>
        <input type="text" id="oauth_redirect_url" v-model="oauthRedirectUrl" />
      </div>


      <button @click="connectMyAccount" v-if="clientId && clientSecret">Connect my Account</button>
      <br>
      <br>
      <button @click="closeModal">Fermer</button>
    </div>
  </div>
</template>

<script>

export default {
  data() {
    return {
      cred_name: "GitHub Api Account", //TODO: put  number if is already exist
      token: '',
      user: '',
      showCredentialsModal: true,
      oauthRedirectUrl: "https://api.github.com",
    };
  },

  methods: {
    closeModal() {
      console.log("close modal");
      this.showCredentialsModal = false;
      this.$emit('close'); // Émet un événement pour fermer la modal
    },
    connectMyAccount() {
      //TODO: sent it  to back end
    },

  },
};
</script>


<style scoped>

.close-button
{
  position: absolute;
  top: 23%;
  right: 30%;
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
