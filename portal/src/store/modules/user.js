import UserService from '../../services/user'

// initial state
const state = () => ({
  user: {
    name: '',
    school: '',
    id: '',
    email: '',
    role: '',
    mentorId: '',
    groups: [],
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
  async login ({ commit, dispatch, state }, { email, password }) {
    const user = await UserService.getUserWithCredentials({ email, password })
    commit('SET_USER', user)
    commit('SET_LOGGED_IN', true)

    const assignmentPayload = {
      id: state.user.id, 
      mentorId: state.user.mentorId,
      groups: state.user.groups,
    }


    await dispatch('assignments/getStudentAssignments', assignmentPayload, {root: true})
    await dispatch('blasts/getBlasts', state.user.groups, {root: true})
    await dispatch('users/getUsers', state.user.groups, {root: true})
  }
}

// mutations
const mutations = {
  SET_USER (state, user) {
    state.user = user
    state.mentorId = user.mentorId
    state.groups = user.groups
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