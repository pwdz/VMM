import React, { useState } from 'react';
import './Login.css'; // Import the CSS file

function Login() {
  // Define state variables for username and password
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  // Define a function to handle form submission
  function handleSubmit(event) {
    event.preventDefault();

    // Create a JSON object with the form data
    const formData = {
      username,
      password,
    };

    // Send a POST request to your API endpoint
    fetch("http://127.0.0.1:8000/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(formData),
    })
      .then((response) => response.json())
      .then((data) => {
        // Handle the API response here
        if (data.success) {
          // Login was successful, you can redirect or perform other actions
          alert("Login successful!");
        } else {
          // Login failed, display an error message
          alert("Login failed. Please check your credentials.");
        }
      })
      .catch((error) => {
        console.error("Error:", error);
        alert("An error occurred during login.");
      });
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
          <a href="sign-up.html">Sign Up</a>
        </div>
      </form>
    </div>
  );
}

export default Login;
