import { createApp } from 'vue'
import { createPinia } from 'pinia'
import FeedLayout from './components/FeedLayout.vue'
import App from './App.vue'
import router from './router'
import('preline')

import './assets/main.css'

const app = createApp(App)

app.config.globalProperties.$hostname = 'http://localhost:3000'
const pinia = createPinia();
pinia.use(({ store }) => {
  store.$hostname = app.config.globalProperties.$hostname;
});
app.use(pinia);
app.use(router)
app.component("FeedLayout",FeedLayout)

app.mount('#app')
