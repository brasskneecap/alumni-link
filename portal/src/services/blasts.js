import utils from "./utils.js";

const BASE_URL = "blasts";

const getBlasts = async ({groupId}) => {
  try {
      const request = {
        url: `${BASE_URL}/${groupId}/`,
      }
      const response = utils.request(request)

      return response;
  } catch (error) {
      console.error("Error fetching blasts:", error);
      throw error;
  }
};

export default {
  getBlasts
}