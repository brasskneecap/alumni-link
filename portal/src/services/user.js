import utils from "./utils.js";

const BASE_URL = "user";

const getUserWithCredentials = async (credentials) => {
  try {
      const request = {
        url: `${BASE_URL}/login`,
        body: credentials,
      }
      const response = utils.request(request)

      return response;
  } catch (error) {
      console.error("Error logging in:", error);
      throw error;
  }
};

export default {
  getUserWithCredentials
}