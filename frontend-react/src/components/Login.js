import React, { useState } from 'react';
import './Login.css'; // Import the CSS file
import { useNavigate } from "react-router-dom";
import axios from 'axios'; // Import Axios

function Login() {
    const navigate = useNavigate();

    const handleSignupClick = () => {
        navigate('/signup');
    };

    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    // Define a function to save the JWT token to localStorage
    const saveTokenToLocalStorage = (token) => {
        localStorage.setItem('jwt_token', token);
    };


    async function handleSubmit(event) {
      event.preventDefault();
  
      const formData = {
          username,
          password,
      };
  
      try {
          // Make the POST request using Axios
          const response = await axios.post("http://127.0.0.1:8000/login", formData, {
              headers: {
                  "Content-Type": "application/json",
              },
          });
  
          const data = response.data;
  
          if (response.status === 200) {
              // Save the JWT token to localStorage
              saveTokenToLocalStorage(data.data);
              console.log(data.data)
  
              // Redirect or perform other actions
              navigate("/")
              console.log(":|")
          } else {
              alert("Login failed. Please check your credentials.");
          }
      } catch (error) {
          console.error("Error:", error);
          alert("An error occurred during login.");
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
