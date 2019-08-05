import Vue from 'vue'
import Router from 'vue-router'
import Layout from './layouts/BasicLayout.vue'

Vue.use(Router);

const routes = [
    {
        path: '/dashboard',
        component: Layout,
        children: [
            {
                path: '',
                name: 'default',
                component: () => import('./views/dashboard/Workplace.vue')
            },
            {
                path: 'workplace',
                name: 'workplace',
                component: () => import('./views/dashboard/Workplace.vue')
            },
            {
                path: 'config',
                name: 'config',
                component: () => import('./views/config/Config.vue')
            },
            {
                path: 'category',
                name: 'category',
                component: () => import('./views/config/Category.vue')
            },
            {
                path: 'tag',
                name: 'tag',
                component: () => import('./views/config/Tag.vue')
            },
            {
                path: 'add-commission-product',
                name: 'addCommissionProduct',
                component: () => import('./views/media/AddCommissionProduct.vue')
            },
        ],
    },
    {
        path: '/dashboard/login',
        name: 'login',
        component:  () => import('./views/user/Login.vue'),
    },
];


const router = new Router({
    mode: 'history',
    routes
});

router.beforeEach((to, from, next) => {
    /*const title = to.meta && to.meta.title;
    if (title) {
      document.title = title
    }*/
    next();
});
export {
    router
};
