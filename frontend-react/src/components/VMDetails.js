
import React, { useState } from 'react';
import './VMDetails.css';
import { useLocation, useNavigate } from "react-router-dom";
import { createVM, saveChangesToVM } from '../services/apiService';

function VMDetails() {
  const navigate = useNavigate();
  const location = useLocation();
  const isEdit = location.state && location.state.isEdit;

  // Set initial values based on whether it's an edit or create operation
  const initialVmNameValue = isEdit ? location.state.vmName : '';
  const initialRamValue = isEdit ? location.state.ram : '512'; // Default to the first option
  const initialCpuCoresValue = isEdit ? location.state.cpuCores : '1'; // Default to the first option

  // State to track changes in the form
  const [vmName, setVmName] = useState(initialVmNameValue);
  const [ram, setRam] = useState(initialRamValue);
  const [cpuCores, setCpuCores] = useState(initialCpuCoresValue);
  const [isLoading, setIsLoading] = useState(false); // Add a loading state variable
  const [progress, setProgress] = useState(0);
  const [progressInterval, setProgressInterval] = useState(null);



  // Set the button text and submit callback based on whether it's an edit or create operation
  const buttonText = isEdit ? "Save Changes" : "Create";
  const handleSubmit = isEdit ? handleSaveChanges : handleCreate;

  return (
    <div className="container">
      <button id="backButton" className="back-button" onClick={handleBackClick}>
        <svg className="icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
          <path d="M0 0h24v24H0z" fill="none" />
          <path d="M15.5 5.5l-7 7 7 7" />
        </svg>
      </button>
      <h2>{isEdit ? "Edit VM" : "Create New VM"}</h2>
      <form id="createVmForm" onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="vmName">VM Name:</label>
          <input type="text" id="vmName" name="vmName" required value={vmName} onChange={(e) => setVmName(e.target.value)} />
        </div>
        <div className="form-group">
          <label>RAM:</label>
          <select id="ramSelect" name="ramCores" value={ram} onChange={(e) => setRam(e.target.value)}>
            <option value="512">512</option>
            <option value="1024">1024</option>
            <option value="2048">2048</option>
            <option value="4096">3096</option>
          </select>
        </div>
        <div className="form-group">
          <label>CPU Cores:</label>
          <select id="cpuSelect" name="cpuCores" value={cpuCores} onChange={(e) => setCpuCores(e.target.value)}>
            <option value="1">1 Core</option>
            <option value="2">2 Cores</option>
            <option value="3">3 Cores</option>
            <option value="4">4 Cores</option>
          </select>
        </div>
        <button type="submit" className="save-button">
          {buttonText}
        </button>
      </form>
      <div id="alertMessage" className="alert" style={{ display: 'none' }}></div>
      {isLoading && (
        <div className="progress">
          <div
            className="progress-bar"
            role="progressbar"
            style={{ width: `${progress}%` }}
            aria-valuenow={progress}
            aria-valuemin="0"
            aria-valuemax="100"
          >
            {progress}%
          </div>
        </div>
      )}
    </div>
  );

  // Function to handle back button click
  function handleBackClick() {
    navigate("/dashboard");
  }
 // Function to handle Save Changes
  function handleSaveChanges(e) {
    e.preventDefault();
    setIsLoading(true);
    setProgress(0);

    const apiCall = saveChangesToVM(location.state.vmID, vmName, ram, cpuCores);

    const timeout = setTimeout(() => {
      clearInterval(progressInterval); // Stop incrementing progress
      alert('API call is taking longer than expected. Please check your connection.');
      setIsLoading(false);
    }, 30000); // Adjust the timeout value as needed

    apiCall
      .then((success) => {
        clearTimeout(timeout); // Clear the timeout if the API call completes before the timeout
        if (success) {
          navigate('/dashboard');
        } else {
          alert('Error saving changes. Please try again.');
        }
      })
      .catch((error) => {
        console.error('Error saving:', error);
        alert('An error occurred while saving changes.');
      })
      .finally(() => {
        setIsLoading(false);
        clearInterval(progressInterval); // Stop incrementing progress
      });

    // Simulate progress incrementing every second
    const interval = setInterval(() => {
      setProgress((prevProgress) => {
        if (prevProgress < 100) {
          return prevProgress + 1;
        }
        return prevProgress;
      });
    }, 1000);

    setProgressInterval(interval);
  }

  // Function to handle Create
  function handleCreate(e) {
    e.preventDefault();
    setIsLoading(true); // Set isLoading to true when the API call starts
    setProgress(0); // Reset the progress bar
  
    const progressInterval = setInterval(() => {
      if (progress < 100) {
        setProgress((prevProgress) => prevProgress + 1); // Increment progress by 1
      }
    }, 3000); // Adjust the interval time as needed
  
    createVM(vmName, ram, cpuCores)
      .then((success) => {
        clearInterval(progressInterval); // Clear the progress interval when the API call completes
  
        if (success) {
          navigate("/dashboard");
        } else {
          alert('Error creating vm. Please try again.');
        }
      })
      .catch((error) => {
        clearInterval(progressInterval); // Clear the progress interval in case of an error
        console.error('Error creating vm:', error);
        alert('An error occurred while creating the vm.');
      })
      .finally(() => {
        setIsLoading(false); // Set isLoading to false when the API call is completed
      });
  }
}

export default VMDetails;