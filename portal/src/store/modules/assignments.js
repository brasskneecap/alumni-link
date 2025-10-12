import AssignmentService from '../../services/assignments'

// initial state
/*[ 
  Example STATE:
  { 
    "id": "4WnBkxBP9YlrfWOYO33c", 
    "name": "Create LinkedIn Profile", 
    "description": "Submit a url link to your LinkedIn Profile", 
    "groupId": "9CrmcRDbrBW4gldZFiV1", 
    "mentorId": "p28tMvMQBZ81cxjACsMx", 
    "allowed_content": [ "url", "text" ],
    "createdAt": "2025-05-31T18:30:04.255Z", 
    "dueDate": "2025-06-06T18:29:18.526Z", 
    "submission": {
      "assignmentId": "4WnBkxBP9YlrfWOYO33c", 
      "studentId": "4iLuNOb4vm8YjxGyUJ0t", 
      "submittedAt": "2025-05-31T20:56:09.142Z", 
      "status": "submitted", 
      "type": "text", 
      "feedback": "", 
      "content": {
        "url": "website.com"
        "text": "this is a test" 
        "file": "start/location/here.pdf" 
      } 
    }
  } 
]*/
const state = () => ({
  groupAssignments: [],
  assignments: []
})

// getters
const getters = {
  groupAssignments: (state) => state.groupAssignments,
  assignments: (state) => state.assignments,
}

// actions
const actions = {
  async getGroupAssignments ({ commit }, {id, mentorId, groupId}) {
    const groupAssignments = await AssignmentService.getGroupAssignments({mentorId, groupId: groupId})
    commit('SET_GROUP_ASSIGNMENTS', groupAssignments)
    console.log('getGroupAssignments', groupAssignments)
  },

  async getStudentAssignments ({ commit }, {id, mentorId, groupId}) {
    const assignments = await AssignmentService.getStudentAssignments({id, groupId: groupId})
    commit('SET_ASSIGNMENTS', assignments)
    console.log('assignments', assignments)
  },

  async createStudentAssignment ({ commit, rootGetters }, assignmentInfo) {
    const user = rootGetters['user/user']
    assignmentInfo.mentorId = user.id
    assignmentInfo.groupId = user.groups[0]
    assignmentInfo.createdAt = new Date()

    const ok = await AssignmentService.createStudentAssignments(assignmentInfo)
    // commit('SET_ASSIGNMENTS', assignments)
  }
}

// mutations
const mutations = {
  SET_ASSIGNMENTS(state, assignments) {
    state.assignments = assignments
  },
  SET_GROUP_ASSIGNMENTS(state, groupAssignments) {
    console.log("Group Assignments", groupAssignments)
    state.groupAssignments = groupAssignments
  },
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}