import { getUserWithCredentials } from '../../services/user'

// initial state
const state = () => ({
  user: {
    name: '',
    school: '',
    id: '',
    email: ''
  },
  loggedIn: false
})

// getters
const getters = {
  user: (state) => state.user,
  loggedIn: (state) => state.loggedIn
}

// actions
const actions = {
  async login ({ commit }, credentials) {
    console.log(credentials)
    const user = await getUserWithCredentials(credentials)
    commit('SET_USER', user)
    commit('SET_LOGGED_IN', true)
  }
}

// mutations
const mutations = {
  SET_USER (state, user) {
    state.user = user
  },
  SET_LOGGED_IN(state, status) {
    state.loggedIn = status;
  }
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}