import './main.css'
import router from './router'
import { createPinia } from 'pinia'
import { createApp } from 'vue'
import App from './App.vue'

// create app
const app = createApp(App)
// create pinia
const pinia = createPinia()
// use router
app.use(router)
app.use(pinia)

// mount app
app.mount('#app')
