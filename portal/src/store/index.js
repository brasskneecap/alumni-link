import { createStore, createLogger } from 'vuex'
import user from './modules/user'
import assignments from './modules/assignments'
const debug = true

export default createStore({
  modules: {
    user,
    assignments,
  },
  strict: debug,
  plugins: debug ? [createLogger()] : []
})