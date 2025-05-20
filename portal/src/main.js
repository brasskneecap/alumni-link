import { createApp } from 'vue'
import store from "./store"
import router from "./router/index"
import './style.css'
import './variables.scss'
import App from './App.vue'
import { createVuetify } from 'vuetify'
import 'vuetify/styles'
import '@mdi/font/css/materialdesignicons.css'

const vuetify = createVuetify({
  icons: {
    defaultSet: 'mdi',
  },
})


const app = createApp(App)
app.use(vuetify)
app.use(store)
app.use(router)
app.mount('#app')
