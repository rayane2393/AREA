<template>
    <div>
    <PopUp :show="showPopup" :message="Message" @close="closePopup" :type="typeMessage" />
    <div class="container">
      <div class="form-container sign-up">
        <h1>Forgotten password</h1>
        <br />
        <br />
        <form @submit.prevent="sendResetInstructions">
          <input type="email" id="reset-email" v-model="forgot_email" required placeholder="Email" />
          <br />
          <br />
          <button class="btn">Reset</button>
        </form>
      </div>

      <div class="toggle-container">
        <div class="toggle">
          <div class="toggle-panel toggle-left">
            <img class="rounded-image image-logo" src="../assets/Logo_copy.png" alt="Description de l'image">
            <h1>Welcome Back!</h1>
            <p>An email to change your password will be sent to your email address</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>


<script>
import PopUp from '../component/PopUp.vue';
export default {
  components: {
    PopUp,
  },
  data() {
    return {
      forgot_email: '',
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
    sendResetInstructions() {
      this.showMessage('An email to change your password will be sent to your email address', 'info');
      console.log('Email:', this.forgot_email);
      //TODO: send reset instructions
    }
  }
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


.container form {
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
  padding: 80px 40px;
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

.container .sign-in {
  transform: translateX(100%);
  opacity: 0;
  z-index: 5;
  animation: move 0.6s;
}

.container .sign-up {
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

.container .toggle-container {
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


.container .toggle {
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

.container .toggle-left {
  transform: translateX(0);
}

.toggle-right {
  right: 0;
  transform: translateX(0);
}

.container .toggle-right {
  transform: translateX(200%);
}
</style>


  