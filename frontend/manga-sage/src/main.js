import { createApp } from 'vue'
import { createPinia } from 'pinia'
import FeedLayout from './pages/FeedLayout.vue'
import App from './App.vue'
import router from './router'
import('preline')

import './assets/main.css'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.component("FeedLayout",FeedLayout)

app.mount('#app')
