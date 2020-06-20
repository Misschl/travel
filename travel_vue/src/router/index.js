import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from "../components/Login";
import Register from "../components/Register";
import Index from "../components/Index";
import User from "../components/User";


Vue.use(VueRouter);

const routes = [
    {
        path: "/login",
        component: Login,
        name: "login",
    },
    {
        path: '/register',
        component: Register,
        name: 'register'
    },
    {
        path: "/index",
        component: Index,
        name: "index",
        children: [
            {
                path: '/user',
                component: User,
                name: 'user'
            }
        ]
    }
];

const router = new VueRouter({
    routes
});

export default router
