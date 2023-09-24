import axios from 'axios';

// Create a custom Axios instance with a base URL
const axiosInstance = axios.create({
  baseURL: 'http://localhost', // Replace with your API URL
});

// Add an interceptor to include the JWT token in the headers
axiosInstance.interceptors.request.use(
  (config) => {
    // Retrieve the JWT token from localStorage
    const token = localStorage.getItem('jwtToken');

    // If the token exists, add it to the Authorization header
    if (token) {
      config.headers['Authorization'] = `${token}`;
    }

    return config;
  },
  (error) => {
    // Handle any request errors here
    return Promise.reject(error);
  }
);

export default axiosInstance;
