import { createStore, createLogger } from 'vuex'
import user from './modules/user'

const debug = true

export default createStore({
  modules: {
    user,
  },
  strict: debug,
  plugins: debug ? [createLogger()] : []
})