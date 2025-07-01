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
  assignments: []
})

// getters
const getters = {
  assignments: (state) => state.assignments,
}

// actions
const actions = {
  async getStudentAssignments ({ commit }, {id, mentorId, groups}) {
    const assignments = await AssignmentService.getStudentAssignments({id, groupId: groups[0]})
    commit('SET_ASSIGNMENTS', assignments)
  }
}

// mutations
const mutations = {
  SET_ASSIGNMENTS(state, assignments) {
    state.assignments = assignments
  },
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}