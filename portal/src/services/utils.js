import axios from "axios";

// const API_BASE_URL = "https://alumni-link-api.onrender.com";
// const API_BASE_URL = "http://localhost:8080";

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL;

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