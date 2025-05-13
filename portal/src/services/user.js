import axios from "axios";

const API_BASE_URL = "http://localhost:8080";

export const getUserWithCredentials = async (credentials) => {
  try {
      const response = await axios.post(
        `${API_BASE_URL}/user/login`, // REQUEST URL
        credentials, // BODY
        {
          headers: { // Additional information for the request
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