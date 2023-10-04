// Profile.js

import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { fetchUserProfile } from '../services/apiService';

function Profile() {
  const navigate = useNavigate();
  const [userData, setUserData] = useState({});

  useEffect(() => {
    // Fetch user profile data when the component mounts
    fetchUserProfile()
      .then((data) => {
        console.log('profile data:', data)
        setUserData(data);
      })
      .catch((error) => {
        console.error('Error fetching user profile data:', error);
      });
  }, []);

  return (
    <div className="container">
      <button id="backButton" className="back-button" onClick={handleBackClick}>
        <svg className="icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
          <path d="M0 0h24v24H0z" fill="none" />
          <path d="M15.5 5.5l-7 7 7 7" />
        </svg>
      </button>
      <h2>User Profile</h2>
      <div className="profile-info">
        <label htmlFor="username">Username</label>
        <p>{userData.username}</p>
      </div>
      <div className="profile-info">
        <label htmlFor="email">Email</label>
        <p>{userData.email}</p>
      </div>
      <div className="profile-info">
        <label htmlFor="active-vms">Active VMs</label>
        <p>{userData.active_vm_count}</p>
      </div>
      <div className="profile-info">
        <label htmlFor="inactive-vms">Inactive VMs</label>
        <p>{userData.inactive_vm_count}</p>
      </div>
    </div>
  );

  function handleBackClick() {
    navigate('/dashboard');
  }
}

export default Profile;
