import {createApp} from 'vue'
import App from './App.vue'
import '@/config/style/global.less'
import router from "@/router/index";
import store from "@/stores/index"

const app = createApp(App)
app.use(store)
app.use(router)
app.mount('#app')

