import { createApp } from 'vue';
import '@vue/compat';
import './style.css'
import App from './App.vue'
import router from './router'
// import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
const app=createApp(App)
app.config.compatConfig={
    MODE:3
};
app.use(router).mount("#app")