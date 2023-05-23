import {createRouter, createWebHistory} from "vue-router";

import {useUserStore} from "@/stores/userStore.js"

const routes = [
    // {
    //     path: "/root",
    //     name: 'name',
    //     redirect: "/system/user"
    // },
    {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/login/index.vue')
    },
    {
        name: '404',
        path: '/:catchAll(.*)',
        component: () => import('@/views/404.vue')
    },
    // {
    //     path: '/system',
    //     component: () => import('@/views/index.vue'),
    //     children: [
    //         // {
    //         //     path: 'user',
    //         //     component: () => import("@/views/system/user/index.vue"),
    //         // },
    //     ]
    // }

];

const router = createRouter({
    history: createWebHistory('/admin'),
    routes,
});

router.beforeEach(async (to, from) => {
    //console.log(router.getRoutes())
    const store = useUserStore()
    if (router.hasRoute(to.name) && to.name != "404") {
        return true
    } else {
        if (store.token) {
            if (store.routes.length == 0) {
                let newRoutes = await store.getRoutes
                newRoutes.forEach(newRoute => {
                    router.addRoute(newRoute)
                })
                if (!store.token) {
                    return {path: "/login"}
                } else if (to.name === '404') {
                    return {path: to.path, query: to.query}
                } else {
                    return {...to, replace: true}
                }
            }
        } else {
            return {path: "/login"}
        }
    }

})

export default router