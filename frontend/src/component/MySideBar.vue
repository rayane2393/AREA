<template>
    <aside class="sidebar" :class="{ 'is-expanded': is_expanded }">
        <!-- LOGO -->
        <div class="logo-t" v-if="is_expanded">
            <img :src="logoURLS" alt="Vue" />
        </div>
        <div class="logo" v-else>
            <img :src="logoURL" alt="Vue" />
        </div>

        <!-- MENU TOGGLE -->
        <div class="menu-toggle-wrap">
            <button class="menu-toggle" @click="ToggleMenu">
                <span>></span>
            </button>
        </div>

        <!-- MENU ITEMS -->
        <div class="menu">

            <router-link to="/malek" class="button" :class="{ 'active': $route.path === '/malek' }">
                <font-awesome-icon class="icon" :icon="faListCheck" />
                <span class="text text-menu">Malek Test</span>
            </router-link>
            <router-link to="/areas" class="button"
                :class="{ 'active': $route.path === '/areas', 'button-hover': !is_expanded }">
                <font-awesome-icon class="icon" :icon="faNetworkWired" />
                <span class="text text-menu">Areas</span>
            </router-link>
            <router-link to="/list-credentials" class="button" :class="{ 'active': $route.path === '/list-credentials' }">
                <font-awesome-icon class="icon" :icon="faKey" />
                <span class="text text-menu">Credentials</span>
            </router-link>
            <router-link to="/executions" class="button" :class="{ 'active': $route.path === '/executions' }">
                <font-awesome-icon class="icon" :icon="faListCheck" />
                <span class="text text-menu">All execution</span>
            </router-link>
        </div>

        <div class="flex"></div>

        <!-- SETTINGS LINK -->
        <div class="menu">
            <router-link to="/settings" class="button" :class="{ 'active': $route.path === '/settings' }">
                <font-awesome-icon class="icon" :icon="faGears" />
                <span class="text text-menu">Settings</span>
            </router-link>
        </div>
    </aside>
</template>
  
<script setup>
import { ref } from 'vue';
import logoURLS from '../assets/Logo_copy.png';
import logoURL from '../assets/Logo.png';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faNetworkWired, faKey, faListCheck, faGears } from '@fortawesome/free-solid-svg-icons';

const is_expanded = ref(localStorage.getItem("is_expanded") === "true");

const ToggleMenu = () => {
    is_expanded.value = !is_expanded.value;
    localStorage.setItem("is_expanded", is_expanded.value);
};
</script>
  
<style lang="scss" scoped>
aside {
    display: flex;
    flex-direction: column;
    background-color: var(--side-bar);
    color: var(--text-color);
    width: 55px;
    overflow: hidden;
    min-height: 100vh;
    padding: 1rem;
    transition: 0.2s ease-in-out;
    transform: translateX(0);


    &.is-expanded {
        transform: translateX(0); // Change to 0 to remove explicit width
    }



    .text-menu {
        margin-left: 10px;
        color: var(--text-color);
        font-size: 15px;
        font-weight: 400;
    }

    .icon {
        margin: 10px;
        font-size: 20px;
        color: var(--text-color);
    }

    .flex {
        flex: 1 1 0%;
    }

    .logo {
        margin-bottom: 1rem;

        img {
            width: 3rem;
        }
    }

    .logo-t {
        margin-bottom: 1rem;

        img {
            width: 10rem;
        }
    }

    .button-small {
        border-radius: 5px;
        height: 10%;
        margin: 10px;
        width: 15%;

    }


    .button {
        border-radius: 5px;
        height: 10%;
        margin: 10px;
        width: 75%;

    }

    .sidebar {
        border: 1px solid var(--border-color);
        border-right: 10px solid black;
        box-shadow: 0 0 5px rgba(0, 0, 0, 0.2);
    }

    .menu-toggle-wrap {
        position: fixed;
        top: 50%;
        right: -10px;
        transform: translateY(-50%);

        .menu-toggle {
            left: --sidebar-width;
            transition: 0.2s ease-in-out;
            border-radius: 50%;

            .material-icons {
                font-size: 10px;
                color: var(--text-color);
                transition: 0.2s ease-in-out;
            }

            &:hover {
                background-color: var(--hover-color);

                .material-icons {
                    color: var(--hover-text);
                    transform: rotate(-180deg);
                }
            }
        }
    }

    // ... Rest of your styles
    h3,
    .button .text {
        opacity: 0;
        transition: opacity 0.3s ease-in-out;
    }



    h3 {
        color: var(--text-color);
        font-size: 10px;
        margin-bottom: 0.5rem;
        text-transform: uppercase;
    }


    .menu {
        margin: 0 -1rem;

        .button-hover {
            &:hover {
                background-color: red;
            }
        }

        .button {
            display: flex;
            align-items: center;
            text-decoration: none;
            transition: 0.2s ease-in-out;
            padding: 0.5rem 1rem;

            .material-icons {
                font-size: 10px;
                color: var(--text-color);
                transition: 0.2s ease-in-out;
            }

            .text {
                color: var(--text-color);
                transition: 0.2s ease-in-out;
            }

            &:hover {
                background-color: var(--hover-color);

            }

            &.active {
                background-color: var(--hover-color);

            }

        }

    }

    .footer {
        opacity: 0;
        transition: opacity 0.3s ease-in-out;

        p {
            font-size: 10px;
            color: var(--text-color);
        }
    }

    &.is-expanded {
        width: var(--sidebar-width);

        .menu-toggle-wrap {
            position: fixed;
            top: 50%;
            right: -10px;
            transform: translateY(-50%);

            .menu-toggle {
                transform: rotate(-180deg);
            }
        }

        h3,
        .button .text {
            opacity: 1;
        }

        .button {
            .material-icons {
                margin-right: 1rem;
            }
        }

        .footer {
            opacity: 0;
        }
    }

    @media (max-width: 1024px) {
        position: absolute;
        z-index: 99;
    }
}
</style>
  