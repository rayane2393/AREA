// Définissez la fonction openPopupAuthUrl en dehors de votre composant Vue
function openPopupAuthUrl(url) {
    // Ouvrir l'URL dans une nouvelle fenêtre contextuelle
    const popup = window.open(url, 'Authorization Popup', 'width=600,height=600');
  
    // Retournez une promesse qui résoudra avec le code d'authentification ou rejettera avec une erreur
    return new Promise((resolve, reject) => {
      const checkPopupClosed = setInterval(() => {
        if (popup.closed) {
          clearInterval(checkPopupClosed);
          reject(new Error('Popup closed before authentication was completed.'));
        }
      }, 1000);
  
      const checkUrlInterval = setInterval(() => {
        try {
          if (popup.location.href.indexOf('http://localhost:8080/list-credentials') === 0) {
            clearInterval(checkUrlInterval);
            const url = popup.location.href;
            popup.close();
  
            // Extraire le code d'authentification de l'URL
            const code = new URL(url).searchParams.get('code');
            if (code) {
              resolve(code);
            } else {
              reject(new Error('Code not found in the URL.'));
            }
          }
        } catch (error) {
          // Ignorer les erreurs qui se produisent lorsque le domaine de l'URL dans la fenêtre contextuelle est différent
        }
      }, 1000);
    });
  }
  
  export default openPopupAuthUrl; // Exportez la fonction pour l'utiliser dans votre composant Vue
  