// authService.js
import axios from 'axios';

export function setAuthToken(token) {
  localStorage.setItem('jwtToken', token);
}

export function getAuthToken() {
  return localStorage.getItem('jwtToken');
}

export function removeAuthToken() {
  localStorage.removeItem('jwtToken');
}

export async function getUserRole() {
  const token = getAuthToken();

  if (!token) {
    // Handle the case where the user is not authenticated
    return Promise.reject('User is not authenticated');
  }

  try {
    // Make an API call to your backend to get the user's role
    const response = await axios.get('/api/get-role', {
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    });

    if (response.status === 200) {
      const data = response.data;
      return data.role; // Return the user's role from the backend
    } else {
      // Handle access denied or other errors
      return Promise.reject('Access denied'); // Customize this message as needed
    }
  } catch (error) {
    console.error('Get User Role API call error:', error);
    return Promise.reject('Error fetching user role');
  }
}