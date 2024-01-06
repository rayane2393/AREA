<template>
  <div class="background">
    <div class="home">
      <router-link to="/areas" class="return-link">
        <img src="../assets/HomePage_Web/return.png" alt="Logo de votre site" class="return" />
      </router-link>
    </div>
    <div class="title-page">Settings.</div>
    <div class="page-button" @click="handleLogout">
      <router-link to="/login" class="logout-link">
        Logout
      </router-link>
    </div>
  </div>
</template>

<script>

export default {
  methods: {
    handleLogout() {
      localStorage.removeItem('token');
    }
  }
};

</script>


<style scoped>
.logout-link {
  margin-top: 1%;
  display: inline-block;
  background-color: white;
  padding: 1%;
  padding-inline: 6%;
  border-radius: 10px;
  transition: filter 3.3s ease;
  color:#9d9d9d;
}

.logout-link:hover {
  background-color: #d52b2b;
  color: white
}

.page-button {
}

.return-link {
  margin-left: -90%;
  margin-top: 1%;
  display: inline-block;
  height: 100%;
  background-color: white;
  padding: 1%;
  border-radius: 10px;
  transition: filter 0.3s ease;
}

.return {
  height: 50%;
  width: 50%;
  transition: filter 0.3s ease;
  filter: brightness(0.7);
}

.return-link:hover {
  background-color: #595959;
}

.return-link:hover .return {
  filter: brightness(1);
}

.background {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color:#c1c2c3;
}

.title-page {
  text-align: center;
  font-weight: bold;
  color: #c1c2c3;
  font-size: 400%;
  margin-top: 2%;
  background-color: white;
  padding: 3%;
  margin-left: 28%;
  margin-right: 28%;
  border-radius: 10px;
}


</style>
