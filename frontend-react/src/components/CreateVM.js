import React from 'react';
import './CreateVM.css';

function CreateVM() {
  // Handle form submission
  const handleSubmit = (e) => {
    e.preventDefault();

    // Replace this with your logic to handle form submission
    const formData = new FormData(e.target);
    const vmData = {
      vmName: formData.get('vmName'),
      status: formData.get('status'),
      ram: formData.get('ram'),
      cpuCores: formData.get('cpuCores'),
    };

    // Perform any API request or state update here
    console.log('Submitted VM Data:', vmData);
  };

  return (
    <div className="container">
      <button id="backButton" className="back-button" onClick={() => handleBackClick()}>
        <svg className="icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
          <path d="M0 0h24v24H0z" fill="none" />
          <path d="M15.5 5.5l-7 7 7 7" />
        </svg>
      </button>
      <h2>Create New VM</h2>
      <form id="createVmForm" onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="vmName">VM Name:</label>
          <input type="text" id="vmName" name="vmName" required />
        </div>
        <div className="form-group">
          <label>Status:</label>
          <div className="radio-group">
            <input type="radio" id="statusOn" name="status" value="on" defaultChecked />
            <label htmlFor="statusOn">On</label>
            <input type="radio" id="statusOff" name="status" value="off" />
            <label htmlFor="statusOff">Off</label>
          </div>
        </div>
        <div className="form-group slider-container">
          <label>RAM:</label>
          <input type="range" id="ramSlider" name="ram" min="512" max="16384" step="512" />
        </div>
        <div className="form-group dropdown">
          <label>CPU Cores:</label>
          <select id="cpuSelect" name="cpuCores">
            <option value="1">1 Core</option>
            <option value="2">2 Cores</option>
            <option value="3">3 Cores</option>
            <option value="4">4 Cores</option>
            <option value="5">5 Cores</option>
            <option value="6">6 Cores</option>
            <option value="7">7 Cores</option>
            <option value="8">8 Cores</option>
          </select>
        </div>
        <button type="submit" className="save-button">
          Save
        </button>
      </form>
      <div id="alertMessage" className="alert" style={{ display: 'none' }}></div>
    </div>
  );

  // Function to handle back button click
  const handleBackClick = () => {
    // Replace with your logic to navigate back to the dashboard or previous page
    window.location.href = "dashboard.html";
  };
}

export default CreateVM;
