import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from "../components/Login";
import Register from "../components/Register";
import Index from "../components/Index";
import User from "../components/User";
import Home from "../components/Home";
import Behavior from "../components/Behavior";
import Friends from "../components/Friends";


Vue.use(VueRouter);

const routes = [
    {
        path: "/",
        redirect: "/login"
    },
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
            },
            {
                path: "/home",
                component: Home,
            },
            {
                path: "/behavior",
                component: Behavior
            },
            {
                path: "/friends",
                component: Friends
            }
        ]
    }
];

const router = new VueRouter({
    routes
});




export default router
