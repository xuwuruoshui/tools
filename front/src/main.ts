import {createApp, toRaw} from 'vue'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'
import VuePdf from 'vue3-pdfjs'
import router from './router'







const app = createApp(App)
// 持久化插件
const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

app.use(pinia)
app.use(VuePdf)
app.use(ElementPlus)
app.use(router)

app.mount('#app')