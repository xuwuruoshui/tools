import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'
import VuePdf from 'vue3-pdfjs'
import router from './router'


const app = createApp(App)


app.use(createPinia())
app.use(VuePdf)
app.use(ElementPlus)
app.use(router)

app.mount('#app')