import { createRouter, createWebHistory } from 'vue-router';

import UserHome from './views/UserHome.vue';
import UserLogin from './views/UserLogin.vue';
import CreateArea from './views/CreateArea.vue';
import UserDashboard from './views/UserDashboard.vue';
import ForgotPassword from './views/ForgotPassword.vue';
import HomePage from './views/HomePage.vue';
import SelectServiceIf from './views/SelectServiceIf.vue';
import SelectServiceThen from './views/SelectServiceThen.vue';
import ThenSpotify from './views/ThenSpotify.vue';
import SettingsPage from './views/SettingsPage.vue';
import MalekPage from './views/MalekPage.vue';
import callbackGithub from './views/callbackOAuth/callbackGithub.vue';
import callbackGoogle from './views/callbackOAuth/callbackGoogle.vue';
import SelectAction from './views/SelectAction.vue';
import SelectReaction from './views/SelectReaction.vue';
import CreateCredentials from './views/CreateCredentials.vue';
import ListCredentials from './views/ListCredentials.vue';
import ClientApkDownload from './component/ClientApkDownload.vue';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/client.apk',
      name: 'client-apk',
      component: ClientApkDownload,
    },
    {
      path: '/callbackgithub',
      name: 'callbackGithub',
      component: callbackGithub
    },
    {
      path: '/callbackgoogle',
      name: 'callbackGoogle',
      component: callbackGoogle
    },
    {
      path: '/',
      name: 'home',
      component: UserHome
    },
    {
      path: '/createArea',
      name: 'createArea',
      component: CreateArea,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/create-credentials',
      name: 'createCredentials',
      component: CreateCredentials,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/list-credentials',
      name: 'listCredentials',
      component: ListCredentials,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/login',
      name: 'login',
      component: UserLogin
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: UserDashboard,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/forgot-password',
      name: 'ForgotPassword',
      component: ForgotPassword
    },
    {
      path: '/areas',
      name: 'HomePage',
      component: HomePage,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/select-service-if',
      name: 'SelectServiceIf',
      component: SelectServiceIf,
      meta: {
        requiresAuth: true
      }

    },
    {
      path: '/select-service-then',
      name: 'SelectServiceThen',
      component: SelectServiceThen,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/then-spotify',
      name: 'ThenSpotify',
      component: ThenSpotify,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/settings',
      name: 'SettingsPage',
      component: SettingsPage,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/malek',
      name: 'malekpage',
      component: MalekPage
    },
    {
      path: '/select-action',
      name: 'SelectAction',
      component: SelectAction,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/select-reaction',
      name: 'SelectReaction',
      component: SelectReaction,
      meta: {
        requiresAuth: true
      }
    }
  ]
});

router.beforeEach((to, from, next) => {
  const loggedIn = localStorage.getItem('user');

  if (to.matched.some(record => record.meta.requiresAuth) && !loggedIn) {
    next('/login');
  } else {
    next();
  }
});

export default router;
