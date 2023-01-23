import {createRouter, createWebHistory,} from 'vue-router'
import Home from '../views/Home.vue'
import Index from '../views/Index.vue'
import Login from '../views/auth/Login.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/login',
            name: 'login',
            component: Login
        },
        {
            path: '/',
            component: Index,
            children: [
                {
                    path: '/home',
                    name: 'home',
                    component: Home,
                },
                {
                    path: '/test',
                    name: 'test',
                    component: () => import('../views/Test.vue')
                },
                {
                    path: '/about',
                    name: 'about',
                    component: () => import('../views/About.vue')
                }
            ]
        },
        {
            path: '/404',
            name: '404',
            // route level api_resp-splitting
            // this generates a separate chunk (About.[hash].js) for this route
            // which is lazy-loaded when the route is visited.
            component: () => import('../views/404.vue')
        }
    ]
})

router.beforeEach((to, from) => {

    // 没登陆跳转login页
    const token = window.localStorage.getItem("custom-token")
    if (token == null && to.name != 'login') {
        return {name: 'login'}
    }

    // 默认跳转home页
    if (token != null && to.name == 'login') {
        return {name: "home"}
    }

})
export default router
