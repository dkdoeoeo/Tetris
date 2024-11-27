import './main.css'
import router from './router'

import { createApp } from 'vue'
import App from './App.vue'

// create app
const app = createApp(App)

// use router
app.use(router)

// mount app
app.mount('#app')
