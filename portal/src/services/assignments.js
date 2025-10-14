import utils from "./utils.js";

const BASE_URL = "assignments";

const getGroupAssignments = async ({mentorId, groupId}) => {
  try {
      console.log(`${BASE_URL}/${groupId}/mentor/${mentorId}`)
      const request = {
        url: `${BASE_URL}/${groupId}/mentor/${mentorId}`,
        method: 'GET',
      };
      const response = await utils.request(request)

      return response;
  } catch (error) {
      console.error("Error fetching assignments:", error);
      throw error;
  }
};

const getStudentAssignments = async ({id, groupId}) => {
  try {
      const request = {
        url: `${BASE_URL}/${groupId}/${id}`,
        method: 'GET',
      }
      const response = utils.request(request)

      return response;
  } catch (error) {
      console.error("Error fetching assignments:", error);
      throw error;
  }
};

const createStudentAssignments = async (assignmentInfo) => {
  try {
      console.log(assignmentInfo)
      const request = {
        url: `${BASE_URL}/createAssignment`,
        method: 'POST',
        body: JSON.stringify(assignmentInfo),
      }
      const response = utils.request(request)

      return response;
  } catch (error) {
      console.error("Error fetching assignments:", error);
      throw error;
  }
};

export default {
  getStudentAssignments,
  createStudentAssignments,
  getGroupAssignments
}