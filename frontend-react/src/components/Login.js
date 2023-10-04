import React, { useState } from 'react';
import './Login.css'; // Import the CSS file
import { useNavigate } from 'react-router-dom';
import { setAuthToken, getUserRoleAPI, setUserRole } from '../services/authService';
import { login } from '../services/apiService';

function Login() {

    console.log('in login, defining use navigate with user role');
    const navigate = useNavigate();

    const handleSignupClick = () => {
        // navigate('/signup');
    };

    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    async function handleSubmit(event) {
      event.preventDefault();
  
      try {
        const data = await login(username, password);
  
        if (data) {
          setAuthToken(data.data);
          console.log('login done. jwt token:', data.data);
          const roleData = await getUserRoleAPI()
          if(roleData){
            setUserRole(roleData)
            navigate('/');
          }else{
            alert('Login failed')
          }
        } else {
          alert('Login failed. Please check your credentials.');
        }
      } catch (error) {
        console.error('Error:', error);
        alert('An error occurred during login.');
      }
  }

  return (
    <div className="container">
      <h2>Login</h2>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="username">Username</label>
          <input
            type="text"
            id="username"
            placeholder="Enter your username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </div>
        <div className="form-group">
          <label htmlFor="password">Password</label>
          <input
            type="password"
            id="password"
            placeholder="Enter your password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>
        <div className="form-group">
          <button type="submit" className="login-button">
            Login
          </button>
        </div>
        <div className="forgot-password">
          <a href="forgot-password.html">Forgot Password?</a>
        </div>
        <div className="signup-button">
            <button onClick={handleSignupClick}>Sign Up</button>
        </div>
      </form>
    </div>
  );
}

export default Login;
