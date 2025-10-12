import axios from "axios";

// const API_BASE_URL = "https://alumni-link-api.onrender.com";
// const API_BASE_URL = "http://localhost:8080";

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL;

const request = async ({ url, method = 'GET', body, options }) => {
  try {
    console.log('url', url, method)
    const response = await axios({
      url: `${API_BASE_URL}/${url}`,
      method,                     // now method can be GET, POST, PUT, etc.
      data: body,                 // axios uses `data` for request bodies
      headers: {
        "Content-Type": "application/json",
        ...(options?.headers || {}),
      },
      ...options,                 // any other options (timeout, params, etc.)
    });

      return response.data;
  } catch (error) {
      console.error("Error logging in:", error);
      throw error;
  }
};

export default {
  request
}