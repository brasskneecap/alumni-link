import UsersService from '../../services/users'

// initial state
const state = () => ({
  users: []
})

// getters
const getters = {
  users: (state) => state.users,
}

// actions
const actions = {
  async getUsers ({ commit, dispatch, state }, groups) {
    const users = await UsersService.getUsers(groups)
    // commit('SET_USERS', user)
    console.log('Users', users)
    commit('SET_USERS', users)
  }
}

// mutations
const mutations = {
  SET_USERS (state, users) {
    state.users = users
  },
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}