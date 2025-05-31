import utils from "./utils.js";

const BASE_URL = "assignments";

const getStudentAssignments = async ({id, groupId}) => {
  try {
      const request = {
        url: `${BASE_URL}/${groupId}/${id}`,
      }
      const response = utils.request(request)

      return response;
  } catch (error) {
      console.error("Error fetching assignments:", error);
      throw error;
  }
};

export default {
  getStudentAssignments
}