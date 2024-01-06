<template>
    <div class="intro-animation">
      <div class="intro-content">
        <img src="../assets/animation_area.gif" alt="Votre Logo Animé">
      </div>
    </div>
  </template>
  
  <style scoped>
  .intro-animation {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: #fff;
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 9999; /* Assurez-vous qu'il est au-dessus de votre page d'accueil actuelle */
    animation: fadeOut 5s ease-in-out forwards;
  }
  
  .intro-content {
    text-align: center;
  }
  
  img {
    max-width: 600px; /* Ajustez la taille de votre logo */
  }
  @keyframes fadeOut {
  0% {
    opacity: 1;
  }
  80% {
    opacity: 1;
  }
  100% {
    opacity: 0;
  }
}
  </style>
  