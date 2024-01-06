
import { h } from 'vue';

import { HOST_URL } from '@/config';
<template>
    <div>
        <PopUp :show="showPopup" :message="Message" @close="closePopup" :typeMessage="Type" />
    </div>
</template>
<script>
import PopUp from '../../component/PopUp.vue';
import {HOST_URL, GITHUB_CLIENT_ID, GITHUB_SECRET_CLIENT } from '../../config.js';


export default {
    components: {
        PopUp,
    },
    data() {
        return {
            Message: '',
            showPopup: false,
            typeMessage: '',
        };
    },
    created() {
        const urlParams = new URLSearchParams(window.location.search);
        const authorizationCode = urlParams.get('code');
        console.log("Authorization Code", authorizationCode);

        if (authorizationCode) {
            const clientID = GITHUB_CLIENT_ID;
            const clientSecret = GITHUB_SECRET_CLIENT;
            const redirectURI = HOST_URL + '/callbackgithub';
            const tokenUrl = `https://github.com/login/oauth/access_token`;

            const requestBody = {
                client_id: clientID,
                client_secret: clientSecret,
                code: authorizationCode,
                redirect_uri: redirectURI
            };

            fetch(tokenUrl, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Accept': 'application/json',
                },
                body: JSON.stringify(requestBody)
            })
                .then(response => response.json())
                .then(data => {
                    if (data.access_token) {
                        // Vous avez récupéré l'access token ici, vous pouvez l'utiliser comme nécessaire
                        const accessToken = data.access_token;
                        console.log("Access Token", accessToken);

                        // Redirigez ensuite l'utilisateur vers votre page d'accueil ou une autre page
                        this.$router.push({
                            name: 'HomePage',
                        });
                    }
                })
                .catch(error => {
                    console.error("Erreur lors de la récupération de l'access token :", error);
                });
        }
    },
    methods: {
        showMessage(message, type) {
            this.Message = message;
            this.showPopup = true;
            this.typeMessage = type;
        },
        closePopup() {
            this.showPopup = false;
            this.Message = '';
            this.typeMessage = '';
        },
    },
};
</script>
