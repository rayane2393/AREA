<template>
  <div class="background">
    <div class="home">
      <div class="title-page">
        <router-link to="/areas" class="return-link">
          <img src="../assets/HomePage_Web/return.png" alt="Logo de votre site" class="return" />
        </router-link>
        <div class="title"> Create a new area.</div>
      </div>
      <div class="area-container">
        <div class="area">
          <div class="if-then-container">
            <router-link to="/select-service-if" class="if-rectangle"
              @mouseenter="setIfHover(true)"
              @mouseleave="setIfHover(false)"
              :class="{ 'if-hover': isIfHover }"
            >
              <img v-if="selectedServiceAction" :src="actionImage" alt="Action Image" class="action-service" />
              <span v-else>Select an If</span>
            </router-link>
            <router-link to="/select-service-then" class="then-rectangle"
              @mouseenter="setThenHover(true)"
              @mouseleave="setThenHover(false)"
              :class="{ 'then-hover': isThenHover }"
            >
              <img v-if="selectedServiceReaction" :src="reactionImage" alt="Reaction Image" class="reaction-service" />
              <span v-else>Select a Then</span>
            </router-link>
          </div>
          <div class="title-rectangle">
            <input
            v-if="editTitle"
              ref="titleInput"
              class="title-input"
              type="text"
              v-model="Title"
              @blur="saveTitle"
              @keyup.enter="disableTitleEdit"
            />
            <div v-else class="title-display">
              {{ Title || 'Click to set a title for your area' }}
            </div>
          </div>
        </div>
        <div class="create-area-container">
          <div class="title-create-area">Fill your area and click on create.</div>
          <button 
            @click="CreateArea" 
            :class="{ 'button-create-area': true, 'disabled': isCreateDisabled }"
            :disabled="isCreateDisabled"
            @mouseover="tooltip = isCreateDisabled"
            @mouseleave="tooltip = false"
          >
            Create area
            <div v-show="tooltip" class="tooltip">{{ tooltipText }}</div>
          </button>
          <div v-if="tooltip" class="tooltip">{{ tooltipText }}</div>
        </div>
      </div>
    </div>
  </div>
</template>


<script>

export default {
  name: "HomePage",
  data() {
    return {
      tooltip: false,
      tooltipText: 'You have to choose an action, a reaction and a title to create your area',
      Title: localStorage.getItem('areaTitle') || '',
      editTitle: !localStorage.getItem('areaTitle'),
      isTitleHover: false,
      isIfHover: false,
      isThenHover: false,
      selectedServiceAction: localStorage.getItem('selectedServiceAction'),
      selectedServiceReaction: localStorage.getItem('selectedServiceReaction'),
    };
  },
  computed: {
    isCreateDisabled() {
      return !this.Title || !this.selectedServiceAction || !this.selectedServiceReaction;
    },
    actionImage() {
      const action = localStorage.getItem('selectedServiceAction');
      return action ? this.getImage(action) : require(`../assets/logo_services/none.png`);
    },
    reactionImage() {
      const reaction = localStorage.getItem('selectedServiceReaction');
      return reaction ? this.getImage(reaction) : require(`../assets/logo_services/none.png`);
    },
  },
  methods: {
    CreateArea() {
      if (this.isCreateDisabled) {
        this.showError = true;
        setTimeout(() => {
          this.showError = false;
        }, 3000);
        return;
      }
      const token = localStorage.getItem('token');
      const action = localStorage.getItem('selectedServiceAction');
      const reaction = localStorage.getItem('selectedServiceReaction');
      const action_name = localStorage.getItem('selectedAction');
      const reaction_name = localStorage.getItem('selectedReaction');
      const arg_action_name = localStorage.getItem('selectedActionArguments');
      const arg_action_title = localStorage.getItem(arg_action_name);
      const arg_reaction_name = localStorage.getItem('selectedReactionArguments');
      const arg_reaction_title = localStorage.getItem(arg_reaction_name);

      console.log(token)
      console.log(action, reaction, action_name, reaction_name);
      console.log(arg_action_name, arg_action_title);
      const url = 'https://back-area.lekmax.fr/create-area';
      const headers = {
        'Content-Type': 'application/json',
        'x-api-key': 'ef4f85d0-6128-11ee-8c99-0242ac120002',
        'bearer': token
      };
      const body = JSON.stringify({
        name : this.Title,
        color : "#FF0000",
        action : action,
        reaction : reaction,
        toggled: true,
        action_name: action_name,
        reaction_name: reaction_name,
        [arg_action_name]: arg_action_title,
        [arg_reaction_name]: arg_reaction_title,
      });
      console.log(body)
      fetch(url, {
        method: 'POST',
        headers: headers,
        body: body
      })
      .then(response => response.json())
      .then(data => { 
          console.log('Réponse du serveur:(malek)', data);
          if (data.message === 'AREA created succesfully') {
            localStorage.removeItem('selectedServiceAction');
            localStorage.removeItem('selectedServiceReaction');
            localStorage.removeItem('selectedAction');
            localStorage.removeItem('selectedReaction');
            localStorage.removeItem('areaTitle');
            this.$router.push({
              name: 'HomePage',
            });
          }
          else {
            console.error('Erreur lors de la connexion(malek):', data.message);
          }
        })
        .catch(error => {
          console.error('Erreur lors de la connexion:(malek)', error);
        });
    },
    setTitleHover(value) {
      this.isTitleHover = value;
    },
    setIfHover(value) {
      this.isIfHover = value;
    },
    setThenHover(value) {
      this.isThenHover = value;
    },
    getImage(serviceName) {
      try {
        return require(`../assets/logo_services/${serviceName}.png`);
      } catch {
        return require(`../assets/logo_services/none.png`);
      }
    },
    saveTitle() {
      localStorage.setItem('areaTitle', this.Title);
    },
    enableTitleEdit() {
      this.editTitle = true;
      this.$nextTick(() => {
        const input = this.$refs.titleInput;
        input.focus();
        input.setSelectionRange(0, input.value.length);
      });
    },
    disableTitleEdit() {
      this.editTitle = false;
      this.saveTitle();
    },
  },
};
</script>

<style scoped>

.button-create-area.disabled {
  background-color: #ccc;
  cursor: not-allowed;
}

.error-message {
  background-color: red;
  color: white;
  padding: 10px;
  margin-top: 10px;
  border-radius: 5px;
  position: absolute;
  top: 10px;
  right: 10px;
  left: 10px;
  text-align: center;
  animation: fadeInOut 3s;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

@keyframes fadeInOut {
  0%, 100% { opacity: 0; }
  10%, 90% { opacity: 1; }
}

.title-display {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 10px;
  cursor: pointer;
  background-color: #FFFFFF;
}
.title-display:hover,
.title-input:focus {
  background-color: #f0f0f0;
}
.action-service, .reaction-service {
  width: 100px;
  height: 100px;
}
.return-link {
  margin-top: 1%;
  margin-left: 1%;
  display: inline-block;
  background-color: white;
  padding: 1%;
  border-radius: 10px;
  transition: filter 0.3s ease;
  height: 3%;
  width: 3%;
}

.return {
  height: 40%;
  width: 50%;
  margin-left: 20%;
  transition: filter 0.3s ease;
  filter: brightness(0.7);
}

.return-link:hover {
  background-color: #bbbbbb;
}

.return-link:hover .return {
  filter: brightness(1);
}
.image-if-then {
  width: 45%;
  height: 50%;
}
.background {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}
.home {
  position: relative;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: #78CBF9;
  overflow: hidden;
  border-radius: 0px;
  text-align: left;
  flex-direction: column;
}
.title {
  text-align: center;
  font-weight: bold;
  color: #78CBF9;
  font-size: 300%;
  margin-top: -1%;
  margin-bottom: 4%;
  padding: 2%;
  margin-left: 28%;
  margin-right: 28%;
  border-radius: 10px;
  flex-direction: column;
  background-color: white;
}
.title-input.hoverTitle::placeholder {
  color: #78CBF9;
}
.area-container {
  display: flex;
  flex-wrap: wrap;
  width: 100%;
  margin-left: 20%;
  gap: 1%;
  min-height: 100vh;
  overflow-y: auto;
  flex-direction: row;
}
.create-area-container {
  display: flex;
  flex-direction: column;
  width: 45%;
  height: 20%;
  margin-bottom: 1%;
  margin-left: 2%;
  margin-right: 1%;
  margin-top: 4%;
  background-color: #78CBF9;
  border-radius: 10px;
  justify-content: center;
  align-items: center;
}

.title-create-area {
  text-align: center;
  font-weight: bold;
  color: #78CBF9;
  font-size: 200%;
  margin-top: 2%;
  margin-bottom: 6%;
  margin-left: 5%;
  margin-right: 5%;
  border-radius: 10px;
  flex-direction: column;
  background-color: white;
  padding: 5%;
}
.button-create-area {
  background-color: #ffffff;
  color: rgb(187, 187, 187);
  text-align: center;
  border: none;
  border-radius: 10px;
  padding: 4% 20%;
  cursor: pointer;
  font-size: 1.5rem;
  position: relative;
}

.tooltip {
  position: absolute;
  bottom: 100%;
  left: 50%;
  transform: translateX(-50%);
  background-color: #333;
  color: #fff;
  padding: 8px 12px;
  border-radius: 6px;
  font-size: 14px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.2);
  white-space: nowrap;
  z-index: 10;
  opacity: 0;
  visibility: hidden;
  transition: opacity 0.3s, visibility 0.3s;
}

.button-create-area:hover .tooltip {
  opacity: 1;
  visibility: visible;
}

.button-create-area.disabled:hover .tooltip {
  opacity: 1;
  visibility: visible;
}


.area {
  flex: 0 0 calc(30% - 1%);
  height: 66vh;
  background-color: #EEEEEE;
  border-radius: 10px;
  margin-bottom: 1%;
}

.if-then-container {
  display: flex;
  justify-content: space-between;
  height: 40%;
  margin-left: 5%;
  margin-right: 5%;
  margin-top: 5%;
}

.if-rectangle, .then-rectangle {
  width: 47%;
  height: 80%;
  background-color: #FFFFFF;
  border-radius: 10px;
  color: #cccccc;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 150%;
  font-weight: bold;
}

.title-rectangle {
  width: 90%;
  height: 55%;
  border-radius: 10px;
  background-color: #FFFFFF;
  margin-left: 5%;
  margin-top: -5%;
}

.title-input {
  width: 100%;
  height: 100%;
  border: none;
  background-color: #FFFFFF;
  font-size: 150%;
  text-align: center;
  border-radius: 10px;
}
.title-input::placeholder {
  color: #CCCCCC;
  font-weight: bold;
  text-decoration: underline;
}


.if-rectangle.if-hover {
  color: #FF66C4
}

.then-rectangle.then-hover {
  color: #FFC144;
}

.title-rectangle.hoverTitle {
  color: #FFFFFF;
}

</style>
