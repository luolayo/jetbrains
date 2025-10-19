import {createRouter, createWebHistory} from "vue-router";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [{
        path: '/',
        name: 'Index',
        component: () => import('../views/IndexView.vue')
    },
        {path: '/download', name: 'Download', component: () => import('../views/DownloadView.vue')},
        {path: '/activation', name: 'Activation', component: () => import('../views/ActivationView.vue')},
    ],
})

export default router