import { BootstrapVue } from '@ankurk91/bootstrap-vue'
import  { createApp } from 'vue';
import './style.css'
import App from './App.vue'
import router from './router'
// import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
const app=createApp(App)
app.use(router).use(BootstrapVue).mount("#app")