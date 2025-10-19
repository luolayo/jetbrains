import {createApp} from 'vue'
import App from './App.vue'
import './style.css';
import {createPinia} from "pinia";

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from "./router";


const app = createApp(App)
const pinia = createPinia();
app.use(pinia);
app.use(ElementPlus)
app.use(router)
app.mount('#app')
