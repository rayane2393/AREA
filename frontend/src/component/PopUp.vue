<template>
    <div v-if="show" class="popup" @click="closePopup">
        <div class="popup-content"
            :class="[type === 'error' ? 'error-popup' : 'info-popup', { 'slide-in': showPopup, 'slide-out': hidePopup }]">
            <p class="error-message">
                <span :class="type === 'error' ? 'error-icon' : 'info-icon'">
                    {{ type === 'error' ? '⚠' : '✔' }}
                </span>
                {{ message }}
            <span class="close" @click="closePopup">×</span>
            </p>
        </div>
    </div>
</template>


<script>
export default {
    props: {
        show: Boolean,
        message: String,
        type: String,
    },
    data() {
        return {
            showPopup: true,
            hidePopup: false,
        };
    },
    methods: {
        closePopup() {
            this.hidePopup = true;
            setTimeout(() => {
                this.showPopup = false;
                this.$emit('close');
            }, 500); // 500ms for the exit animation
        },
    },
};
</script>

<style scoped>
/* Style CSS pour votre pop-up */
.popup {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 9999;
}

.popup-content {
    border-radius: 5px;
    padding: 0 20px;
    text-align: center;
    position: relative;
    transition: all 0.5s;
    /* Animation de 0.5 seconde pour l'entrée et la sortie */
    transform: translateY(-100%);
    /* Déplace la popup en dehors de l'écran par défaut */
}

.error-popup {
    background-color: #ffc0cb;
    border: 1px solid #ff0000;
}

.info-popup {
    background-color: #b9ffaf;
    /* Vert clair */
    border: 1px solid #00ff00;
    /* Contour vert vif */
}

.slide-in {
    transform: translateY(0);
    /* Rétablit la position de la popup (entrée) */
}

.slide-out {
    transform: translateY(-100%);
    /* Déplace la popup en dehors de l'écran (sortie) */
}

.close {
    position: absolute;
    top: 2px;
    right: 5px;
    font-size: 20px;
    cursor: pointer;
}



.error-message {
    font-weight: bold;
    font-size: 12px;
}

.error-icon {
    color: #ff0000;
    /* Rouge vif */
    font-size: 20px;
}

.info-icon {
    color: #00ff00;
    /* Vert vif */
    font-size: 20px;
}</style>
