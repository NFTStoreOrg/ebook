import { BootstrapVue } from '@ankurk91/bootstrap-vue'
import { createApp } from 'vue';
import { createPinia } from 'pinia'
import './style.css'
import App from './App.vue'
import router from './router'
import 'bootstrap/dist/css/bootstrap.css'
const app = createApp(App)
const pinia = createPinia()

app.use(router).use(BootstrapVue).use(pinia).mount("#app")