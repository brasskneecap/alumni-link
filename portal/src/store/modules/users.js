import UsersService from '../../services/users'

const roles = {
  student: "Student"
}

// initial state
const state = () => ({
  users: [],
  students: [],
})

// getters
const getters = {
  users: (state) => state.users,
  students: (state) => state.students,
}

// actions
const actions = {
  async getUsers ({ commit, dispatch, state }, groups) {
    const users = await UsersService.getUsers(groups)
    commit('SET_USERS', users)
    const students = users.filter((user) => user.role === roles.student)
    commit('SET_STUDENTS', students)

  }
}

// mutations
const mutations = {
  SET_USERS (state, users) {
    state.users = users
  },
  SET_STUDENTS (state, students) {
    state.students = students
  },
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}