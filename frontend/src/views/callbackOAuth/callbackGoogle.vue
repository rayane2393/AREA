
<template>
    <div>
        <PopUp :show="showPopup" :message="Message" @close="closePopup" :typeMessage="Type" />
    </div>
</template>
<script>
import PopUp from '../../component/PopUp.vue';
import { globalVariable } from '../../global.js';
import { HOST_URL, GOOGLE_CLIENT_ID, GOOGLE_SECRET_CLIENT } from '../../config.js';

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
    mounted() {
        this.extractAuthorizationCode();
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
        extractAuthorizationCode() {
            const urlParams = new URLSearchParams(window.location.search);
            const authorizationCode = urlParams.get('code');
            if (authorizationCode) {
                const tokenEndpoint = 'https://accounts.google.com/o/oauth2/token';
                const clientId = GOOGLE_CLIENT_ID;
                const redirectUri = HOST_URL + '/callbackgoogle';
                const clientSecret = GOOGLE_SECRET_CLIENT;
                const postData = {
                    code: authorizationCode,
                    client_id: clientId,
                    client_secret: clientSecret,
                    redirect_uri: redirectUri,
                    grant_type: 'authorization_code',
                };

                tokenEndpoint;
                postData;
                //TODO: envoyer le code au back
                globalVariable.value = "false token";
                localStorage.setItem('user', 'authenticated');
                this.$router.push({
                    name: 'HomePage',
                });
            }
            /*
            const postData = {
                code: authorizationCode,
                client_id: clientId,
                client_secret: clientSecret,
                redirect_uri: redirectUri,
                grant_type: 'authorization_code',
            };
            fetch(tokenEndpoint, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: new URLSearchParams(postData),
            })
                .then(response => response.json())
                .then(data => {
                    const accessToken = data.access_token;
                    console.log('Code OAuth (jeton d\'accès) obtenu :', accessToken); //FIXME: remove
                    // TODO: Envoyer le code OAuth au back
                    // message gg vous etes connecté
                    globalVariable.value = "false token";
                    localStorage.setItem('user', 'authenticated');
                    this.$router.push({
                        name: 'HomePage',
                    });
                })
                .catch(error => {
                    console.error('Erreur lors de la récupération du code OAuth :', error);
                    this.showPopup('Erreur lors de la récupération du code OAuth', "error");
                });*/

        }
    },
};
</script>
