import { createApp } from "vue";
import { initHttp } from "./api";
import { createPinia } from 'pinia'
import App from "./App.vue";

const app = createApp(App)

const pinia = createPinia()
app.use(pinia)

initHttp()

app.mount("#app");