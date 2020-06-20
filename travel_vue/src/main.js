import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './plugins/element.js'
import axios from 'axios'

Vue.config.productionTip = false;

var baseUrl = Vue.config.productionTip ? "http://192.168.101.25" : "http://192.168.101.25";

axios.defaults.baseURL = baseUrl;

Vue.prototype.$http = axios;


new Vue({
    router,
    store,
    render: h => h(App)
}).$mount('#app');
