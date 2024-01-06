<template>
  <div class="background">
    <div class="home">
      <router-link to="/settings" class="settings-link">
        <img src="../assets/HomePage_Web/settings.png" alt="Logo de votre site" class="settings" />
      </router-link>
      <div class="title-page"> List of your areas.</div>
      <div class="area-container">
        <div class="area" v-for="area in areas" :key="area.id">
          <div class="if-then-container">
            <div class="service-image">
              <img :src="getImage(area.action)" />
            </div>
            <div class="service-image">
              <img :src="getImage(area.reaction)" />
            </div>
          </div>
          <div class="title-rectangle">
            {{ area.name }}
          </div>
          <div class="toggle-switch">
            <input type="checkbox" :id="'toggle-' + area.id" :checked="area.toggled" @change="toggleArea(area.id, $event.target.checked)" class="toggle-input">
            <label :for="'toggle-' + area.id" class="toggle-label"></label>
          </div>
          <button class="delete-button" @click="prepareDelete(area.id, area.action)">
            <img src="../assets/HomePage_Web/trash.png" alt="Delete" class="trash-icon">
          </button>
          <div v-if="showDeleteModal" class="delete-modal">
            <div class="modal-content">
              <p>Êtes-vous sûr de vouloir supprimer cet area ?</p>
              <button @click="confirmDelete">Confirmer</button>
              <button @click="showDeleteModal = false">Annuler</button>
            </div>
          </div>


        </div>
        <router-link to="/createArea">
          <img src="../assets/HomePage_Web/plus.png" class="button-add-area"/>
        </router-link>
      </div>

    </div>
  </div>
</template>

<script>
import { globalVariable } from '../global.js';

export default {
  data() {
    return {
      areas: [],
      globalVariable: localStorage.getItem('token'),
      showDeleteModal: false,
      areaToDelete: null
    };
  },
  methods: {
  getImage(serviceName) {
    try {
      return require(`../assets/logo_services/${serviceName}.png`);
    } catch {
      return require(`../assets/logo_services/none.png`);
    }
  },
  async getAreas() {
    try {
      const token = localStorage.getItem('token');
      const url = 'https://back-area.lekmax.fr/get-area';
      const headers = {
        'Content-Type': 'application/json',
        'x-api-key': 'ef4f85d0-6128-11ee-8c99-0242ac120002',
        'bearer': token
      };
      const response = await fetch(url, {
        method: 'POST',
        headers: headers,
      });
      const json = await response.json();
      this.areas = json.areas;
      console.log(this.areas);
    } catch (error) {
      console.error('Erreur lors de la récupération des areas:', error);
    }
  },
    created() {
      this.getAreas();
    },
    async deleteArea(id, action) {
      const token = globalVariable.value;
      console.log(token);
      console.log(action);
      try {
        const response = await fetch('https://back-area.lekmax.fr/delete-area', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            'x-api-key': 'ef4f85d0-6128-11ee-8c99-0242ac120002',
            'bearer': token
          },
          body: JSON.stringify({
            area_id: id,
            action: action
          })
        });
        const json = await response.json();
        console.log(json);
        return json;
      } catch (error) {
        console.error(error);
      }
    },

    prepareDelete(id, action) {
      this.areaActionToDelete = action;
      this.areaToDelete = id;
      this.showDeleteModal = true;
    },

    async confirmDelete() {
      if (this.areaToDelete !== null) {
        const deleteResult = await this.deleteArea(this.areaToDelete, this.areaActionToDelete);
        if (deleteResult) {
          this.areas = this.areas.filter(area => area.id !== this.areaToDelete);
        }
        this.areaToDelete = null;
      }
      this.showDeleteModal = false;
    },
    async toggleArea(id, toggled) {
      const token = localStorage.getItem('token');
      const area = this.areas.find(area => area.id === id);
      console.log(`Toggle area ${id} to ${toggled}`);
      console.log(area);
      console.log(token);

      const url = 'https://back-area.lekmax.fr/modify-area';
      const headers = {
        'Content-Type': 'application/json',
        'x-api-key': 'ef4f85d0-6128-11ee-8c99-0242ac120002',
        'bearer': token
      };
      try {
        const response = await fetch(url, {
          method: 'POST',
          headers: headers,
          body: JSON.stringify({
            name: area.name,
            color: area.color,
            action: area.action,
            reaction: area.reaction,
            toggled: toggled,
            id: id,
            action_name: area.action_name,
            reaction_name: area.reaction_name
          })
        });
        const json = await response.json();
        if(response.ok) {
          console.log('Toggle updated successfully:', json);
          area.toggled = toggled;
        } else {
          console.log('Toggle not uptated:', json);
          console.error('Erreur lors de la mise à jour du toggle:', json.message);
        }
      } catch (error) {
        console.error('Erreur lors de la mise à jour du toggle:', error);
      }
    }
  },
  mounted() {
    console.log('Component mounted.', this.globalVariable.value);
    this.getAreas();
    localStorage.removeItem('selectedServiceAction');
    localStorage.removeItem('selectedServiceReaction');
    localStorage.removeItem('selectedAction');
    localStorage.removeItem('selectedReaction');
    localStorage.removeItem('areaTitle');
  },
};
</script>

<style scoped>

.service-image img {
  width: 50px;
  height: 50px;
  display: flex;
  justify-content: center;
  align-items: center;
}
.settings-link {
  margin-top: 1%;
  margin-left: 1%;
  display: inline-block;
  background-color: white;
  padding: 1%;
  width: 2%;
  height: 4%;
  border-radius: 10px;
  color:#9d9d9d;
}

.settings {
  transition: filter 0.3s ease;
  filter: brightness(0.7);
  height: 100%;
  width: 100%;
}

.settings-link:hover {
  background-color:#CCCCCC;
  color: white
}


.title-page {
  text-align: center;
  font-weight: bold;
  color: #CB6CE6;
  font-size: 400%;
  margin-top: 0%;
  margin-bottom: 2%;
  background-color: white;
  padding: 3%;
  margin-left: 28%;
  margin-right: 28%;
  border-radius: 10px;
}
.title-input.hoverTitle::placeholder {
  color: #78CBF9;
}

.logo {
  margin-top: 1%;
  margin-left: 1%;
  height: 8%;
  width: 10%;
  border-radius: 10px;
  margin-bottom: 1%;
}
.home {
  position : relative;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: #CB6CE6;
  overflow-x: hidden;
  border-radius: 0px;
  text-align: left;

}

.background {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

.button-add-area {
  transition: filter 0.3s ease;

  margin-top: 200%;
  padding-left: 100%;
  padding-right: 100%;
  padding-top: 100%;
  padding-bottom: 100%;
  height: 6%;
  width: 100%;
  background-color: #CCCCCC;
  border-radius: 10px;
}
.button-add-area:hover {
  background-color: #9d9d9d;
}
.area-container {
  display: flex;
  flex-wrap: wrap;
  width: 100%;
  margin-left: 1%;
  gap: 1%;
  min-height: 100vh;
  overflow-y: auto;
}

.area {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  flex: 0 0 calc(19.8% - 1%);
  height: 50vh;
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
  margin-bottom: 5%;
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

/* toggle */
.toggle-switch {
  margin: 5%;
  position: relative;
  width: 50px;
  height: 24px;
  align-self: flex-end;
  margin-top: auto;
}

.toggle-input {
  display: none;
}

.toggle-label {
  position: absolute;
  top: 0;
  left: 0;
  width: 50px;
  height: 24px;
  background-color: #e6e6e6;
  border-radius: 15px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.toggle-label::after {
  content: "";
  position: absolute;
  top: 2px;
  left: 2px;
  width: 20px;
  height: 20px;
  background-color: white;
  border-radius: 50%;
  transition: left 0.3s;
}

.toggle-input:checked + .toggle-label {
  background-color: #4CAF50;
}

.toggle-input:checked + .toggle-label::after {
  left: 28px;
}

.delete-button {
  background: transparent;
  border: none;
  cursor: pointer;
}

.trash-icon {
  width: 20px;
  transition: color 0.3s;
}

.delete-button:hover .trash-icon {
  color: red;
}

.delete-modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-content {
  background-color: white;
  padding: 20px;
  border-radius: 10px;
  text-align: center;
}

</style>
