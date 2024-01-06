<template>
    <div class="card">
        <div class="card-header">


            <div v-if="getvalue('google')">
            <img class="logo" src="../assets/google.svg" :alt="url_logo_alt" />
                </div>
            <div v-if="getvalue('github')">
                <img class="logo" src="../assets/github.svg" :alt="url_logo_alt" />
            </div>
            <div v-if="getvalue('spotify')">
                <img class="logo" src="../assets/spotify.svg" :alt="url_logo_alt" />
            </div>

            <div class="title">{{ title }}</div>
            <span class="timestamp">{{ name }} <span style="font-size: 20px;">|</span> Last updated {{
                convertTimestampToTimeAgo(updated_at) }} <span style="font-size: 20px;">|</span> Created {{
        convertTimestampToTimeAgo(created_at) }}</span>
        </div>
        <div class="actions">
            <div class="dropdown">
                <button @click="openDropdown"><img style="width: 10px;heght: 10px;" src="../assets/set.svg"
                        :alt="url_logo_alt" /></button>
                <div v-show="isDropdownOpen" class="dropdown-content">
                    <button @click="deleteCard">Delete</button>
                    <button @click="openCard">Open</button>
                </div>
            </div>
        </div>
    </div>
</template>
  
<script>
export default {
    props: {
        types: String,
        name: String,
        title: String,
        updated_at: String,
        created_at: String,
        id: Int16Array,
    },
    data() {
        return {
            isDropdownOpen: false,
            url_logos: { "google": "../../src/assets/google.svg", "github": "../../assets/github.svg", "spotify": "../../assets/spotify.svg", "": "pute" },
            url_logo_alt: this.types || "", // Assurez-vous d'avoir une valeur par défaut
            url_logo: "../assets/google.svg",//this.url_logos[this.types] || "", // Assurez-vous d'avoir une valeur par défaut
            type: Object,
            required: true,
        };
    },
    methods: {
        getvalue(key) {
            console.log(key);
            const value = (key === this.types);
            console.log(value);
            return value;
        },
        convertTimestampToTimeAgo(timestamp) {
            const currentDate = new Date();
            const timestampDate = new Date(timestamp);
            const timeDifference = currentDate - timestampDate;
            const secondsDifference = Math.floor(timeDifference / 1000);

            if (secondsDifference < 60) {
                return `${secondsDifference} seconds ago`;
            } else if (secondsDifference < 3600) {
                const minutesDifference = Math.floor(secondsDifference / 60);
                return `${minutesDifference} minutes ago`;
            } else if (secondsDifference < 86400) {
                const hoursDifference = Math.floor(secondsDifference / 3600);
                return `${hoursDifference} hours ago`;
            } else if (secondsDifference < 2592000) {
                const daysDifference = Math.floor(secondsDifference / 86400);
                return `${daysDifference} days ago`;
            } else if (secondsDifference < 31536000) {
                const monthsDifference = Math.floor(secondsDifference / 2592000);
                return `${monthsDifference} months ago`;
            } else {
                const yearsDifference = Math.floor(secondsDifference / 31536000);
                return `${yearsDifference} years ago`;
            }
        },
        openDropdown() {
            this.isDropdownOpen = !this.isDropdownOpen;
        },
        deleteCard() {
            // Fonction pour supprimer la carte
        },
        openCard() {
            // Fonction pour ouvrir la carte
        },
    },
};
</script>
  
<style scoped>
:root {

    --hover-color-light: #DCDFE7;
    --hover-color-dark: #32394A;
    --hover-text-color-light: #555555;
    --hover-text-color-dark: #B8B2A8;

    --hover-color: --hover-color-light;
    --hover-text-color: --hover-text-color-light;

    --green-dark: #0FA552;
    --green-light: #13CE66;
    --orange-light: #FF6D5A;
    --orange-dark: #FF4D4D;
    --side-bar-dark: #252829;
    --side-bar-light: #FFFFFF;

    --background-dark: #1D1F20;

    --background-light: #FBFCFE;
    --text-color-light: #7D7D87;
    --text-color-dark: #E8E6E3;

    --border-color-dark: #736B5E;
    --border-color-light: #555555;

    --card-color-dark: #252829;
    --card-color-light: --side-bar-light;

    --text-side-bar-dark: #B8B2A8;
    --text-side-bar-light: #7D7D87;

    --search-bar-dark: #141516;
    --search-bar-light: --side-bar-light;

    --sidebar-width: 175px;
    --green: --green-light;
    --orange: --orange-light;
    --background: --background-light;
    --side-bar: var(--side-bar-light) --text-color: --text-color-light;
    --border-color: --border-color-light;
    --card-color: --card-color-light;
    --text-side-bar: --text-side-bar-light;
    --search-bar: --search-bar-light;
}

.logo {
    width: 40px;
    height: 40px;
    margin-left: 10px;
    margin-right: 10px;
    margin-top: 10px;
}

.card {
    width: 75%;
    height: 5%;
    margin-bottom: 10px;
    border: 2px solid var(--border-color);
    border-radius: 10px;
    display: flex;
    align-items: left;
    justify-content: space-between;
    background-color: var(--hover-color);
}

.card-header {
    display: flex;
    flex-direction: center;
}

.title {
    font-weight: bold;
    font-size: 1.2rem;
    text-align: left;
    margin-left: 70px;
}

.timestamp {
    color: var(--text);
    font-size: 1.2rem;
    margin-bottom: 10px;
    margin-top: 18px;
    text-align: center;
}

.actions {
    display: flex;
    align-items: center;
}




.dropdown {
    position: relative;
    display: inline-block;
    color: var(--text-color)
}

.dropdown button {
    background-color: transparent;
    border: none;
    cursor: pointer;
}

.dropdown-content {
    display: none;
    position: absolute;
    background-color: #f9f9f9;
    min-width: 160px;
    box-shadow: 0px 8px 16px 0px rgba(0, 0, 0, 0.2);
    z-index: 1;
    right: 0;
}

.dropdown-content a {
    padding: 10px;
    text-decoration: none;
    display: block;
}

.dropdown-content a:hover {
    background-color: #ddd;
}

.dropdown:hover .dropdown-content {
    display: block;
}

.custom-dropdown {
    /* Ajoutez ici les styles personnalisés de votre choix */
    border: 1px solid #ccc;
    border-radius: 4px;
    padding: 8px;
    width: 200px;
    font-size: 16px;
    background-color: white;
    color: #333;
    /* Ajoutez d'autres styles selon vos besoins */
}


.custom-dropdown {
    /* Ajoutez ici les styles personnalisés de votre choix */
    border: 1px solid #ccc;
    border-radius: 4px;
    padding: 8px;
    width: 200px;
    font-size: 16px;
    background-color: white;
    color: #333;
    /* Ajoutez d'autres styles selon vos besoins */
}

/* Styles pour les options du dropdown */
.custom-dropdown option {
    /* Styles pour les options */
    background-color: white;
    color: #333;
}
</style>
  