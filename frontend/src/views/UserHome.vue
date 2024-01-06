<template>
  <div>
    <div class="home">
    <IntroAnimation v-if="showIntro" />
      <header>
        <h1>Area</h1>
        <div class="buttons-container">
          <a href="/app/shared-assets/client.apk" download>
            <button class="btn btn-primary">
              <font-awesome-icon :icon="['fas', 'download']" />Download App
            </button>
          </a>
          <router-link to="/login" class="btn btn-primary">Sign In</router-link>
        </div>
      </header>
      <section class="intro">
        <h2>Automate Your Workflows</h2>
        <p>Area is a cutting-edge automation platform inspired by n8n. Streamline your workflows, integrate your favorite
          tools, and enhance productivity.</p>

        <div class="download-button">
          
        </div>
        <img src="https://unsplash.com/photos/1K9T5YiZ2WU/download?force=true" alt="Automation Image">
      </section>

    <footer>
      <div class="footer-content">
        <p>&copy; 2023 AREA. All rights reserved.</p>
      </div>
    </footer>
    </div>
  </div>
</template>

<script>
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faDownload } from '@fortawesome/free-solid-svg-icons';
import { library } from '@fortawesome/fontawesome-svg-core';
import IntroAnimation from '../component/IntroAnimation.vue';
library.add(faDownload);

export default {
  components: {
    FontAwesomeIcon,
    IntroAnimation,
  },
  data() {
    return {
      showIntro: true,
    };
  },
  mounted() {
    setTimeout(() => {
      this.showIntro = false;
    }, 5000); // 3000 millisecondes (3 secondes)
  },
};
</script>


<style scoped>

.download-button {
  text-align: center;
  margin-top: 20px;
}

.download-link {
  text-decoration: none;
}

.download-btn {
  background-color: #007BFF;
  color: #fff;
  padding: 10px 20px;
  border: none;
  border-radius: 5px;
  font-size: 1.2em;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background-color 0.3s;
}

.download-btn:hover {
  background-color: #0056b3;
}

.fa-download {
  margin-right: 10px;
}
.home {
  font-family: Arial, sans-serif;
  background-color: #f7f7f7;
  color: #333;
  padding: 50px 0;
}

header {
  background-color: #fff;
  padding: 20px 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

h1 {
  margin-left: 50px;
  color: #333;
  font-size: 2em;
}

.download-button {
  position: fixed;
  top:50px;
  right: 240px;
}


h2 {
  color: #333;
  font-size: 2.5em;
  margin-bottom: 20px;
}

.intro {
  text-align: center;
  padding: 50px 0;
}

p {
  max-width: 600px;
  margin: 20px auto;
  line-height: 1.5;
  font-size: 1.2em;
}

img {
  margin-top: 30px;
  max-width: 100%;
  height: auto;
}

.buttons-container {
  margin-right: 50px;
}

.btn {
  margin-left: 10px;
  padding: 10px 15px;
  border: none;
  background-color: #eee;
  color: #333;
  border-radius: 5px;
  text-decoration: none;
  transition: background-color 0.3s;
}

.btn-primary {
  background-color: #007BFF;
  color: #fff;
}

.btn:hover {
  background-color: #ddd;
}

.btn-primary:hover {
  background-color: #0056b3;
}


.logo {
  font-size: 2em;
}


footer {
  background-color: #333;
  color: #fff;
  padding: 10px 0;
  text-align: center;
}

.footer-content {
  font-size: 1em;
}

</style>