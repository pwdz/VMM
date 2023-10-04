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
  removeUserRole()
}
export function setUserRole(role){
  localStorage.setItem('role', role)
}
export function getUserRole(){
  return localStorage.getItem('role')
}
export function removeUserRole(){
  localStorage.removeItem('role')
}
export async function getUserRoleAPI() {
  const token = getAuthToken();
  // console.log(token)

  if (!token) {
    // Handle the case where the user is not authenticated
    return Promise.reject('User is not authenticated');
  }

  try {
    // Make an API call to your backend to get the user's role
    const response = await axios.get('http://127.0.0.1:8000/api/get-role', {
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