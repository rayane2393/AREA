<template>
  <div>
    <!-- Le contenu du composant est vide car nous n'avons besoin que des variables CSS ici. -->
  </div>
</template>

<script>
export default {
  setup() {
    // Définissez vos variables CSS ici
    document.documentElement.style.setProperty('--primary', '#4ade80');
    document.documentElement.style.setProperty('--primary-alt', '#22c55e');
    document.documentElement.style.setProperty('--grey', '#64748b');
    document.documentElement.style.setProperty('--dark', '#1e293b');
    document.documentElement.style.setProperty('--dark-alt', '#334155');
    document.documentElement.style.setProperty('--light', '#f1f5f9');
    document.documentElement.style.setProperty('--sidebar-width', '300px');
  }
};
</script>
