<template>
    <div>
      <!-- Contenu du composant -->
    </div>
  </template>
  
  <script>
  export default {
    mounted() {
      // Générez et renvoyez le fichier APK ici
      this.generateAndSendAPK();
    },
    methods: {
      generateAndSendAPK() {
        // Code pour générer et renvoyer le fichier APK
        // Vous pouvez utiliser des bibliothèques ou des outils pour générer le fichier APK, puis l'envoyer comme réponse HTTP
        const apkPath = '/app/shared-assets/client.apk';
        //const apkPath = '/Users/malekgatoufi/Project/Epitech/MyArea/frontend/client.apk';
        fetch(apkPath)
          .then(response => response.blob())
          .then(blob => {
            const url = window.URL.createObjectURL(blob);
            const a = document.createElement('a');
            a.style.display = 'none';
            a.href = url;
            a.download = 'client.apk';
            document.body.appendChild(a);
            a.click();
            window.URL.revokeObjectURL(url);
            this.$router.push({
                    name: 'home',
                });
          });
      }
    }
  }
  </script>
  