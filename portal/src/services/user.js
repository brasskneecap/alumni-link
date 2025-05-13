import axios from "axios";

const API_BASE_URL = "http://localhost:8080";

export const getUserWithCredentials = async ({email, password}) => {
  try {
      const response = await axios.post(`${API_BASE_URL}/user/login`,
        { email, password },
        {
          headers: {
              "Content-Type": "application/json",
          },
        }
      );

      return response.data;
  } catch (error) {
      console.error("Error logging in:", error);
      throw error;
  }
};