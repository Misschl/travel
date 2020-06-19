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
        component: Login
    },
    {
        path: '/register',
        component: Register
    },
    {
        path: "/index",
        component: Index,
        children: [
            {
                path: '/user',
                component: User
            }
        ]
    }
];

const router = new VueRouter({
    routes
});

export default router
