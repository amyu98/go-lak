import { createApp } from "vue";
import { initHttp } from "./api";
import App from "./App.vue";

initHttp()
createApp(App).mount("#app");