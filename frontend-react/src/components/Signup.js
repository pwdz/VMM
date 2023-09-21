import React from 'react';
import './Signup.css';

function Signup() {
  return (
    <div className="container">
      <h2>Sign Up</h2>
      <form>
        <div className="form-group">
          <label htmlFor="name">Name</label>
          <input type="text" id="name" placeholder="Enter your name" />
        </div>
        <div className="form-group">
          <label htmlFor="email">Email</label>
          <input type="email" id="email" placeholder="Enter your email" />
        </div>
        <div className="form-group">
          <label htmlFor="password">Password</label>
          <input type="password" id="password" placeholder="Enter your password" />
        </div>
        <div className="form-group">
          <input type="submit" value="Sign Up" />
        </div>
      </form>
    </div>
  );
}

export default Signup;
