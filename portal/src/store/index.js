import { createStore, createLogger } from 'vuex'
import assignments from './modules/assignments'
import blasts from './modules/blasts'
import user from './modules/user'
import users from './modules/users'
const debug = true

export default createStore({
  modules: {
    assignments,
    blasts,
    user,
    users,
  },
  strict: debug,
  plugins: debug ? [createLogger()] : []
})