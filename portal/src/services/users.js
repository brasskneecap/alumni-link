import utils from "./utils.js";

const BASE_URL = "user";

const getUsers = async (groups) => {
  try {
      const request = {
        url: `${BASE_URL}/users`,
        body: {groups},
      }
      const response = utils.request(request)

      return response;
  } catch (error) {
      console.error("Error logging in:", error);
      throw error;
  }
};

export default {
  getUsers
}