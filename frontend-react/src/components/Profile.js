import React from 'react';

function Profile() {
  return (
    <div className="container">
      <h2>User Profile</h2>
      <div className="profile-info">
        <label htmlFor="name">Name</label>
        <p>John Doe</p>
      </div>
      <div className="profile-info">
        <label htmlFor="email">Email</label>
        <p>johndoe@example.com</p>
      </div>
      <div className="profile-info">
        <label htmlFor="username">Username</label>
        <p>johndoe123</p>
      </div>
      <div className="dashboard-button">
        <a href="dashboard.html">Go to Dashboard</a>
      </div>
    </div>
  );
}

export default Profile;