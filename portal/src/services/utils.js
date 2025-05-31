import axios from "axios";

const API_BASE_URL = "http://localhost:8080";

const request = async ({ url, body, options }) => {
  try {
      const response = await axios.post(
        `${API_BASE_URL}/${url}`, // REQUEST URL
        body, // BODY
        {
          ...options,
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

export default {
  request
}