import React, { useState, useEffect } from 'react';
import './Dashboard.css';
import { deleteServerByVmId, fetchServerData, toggleServerStatusAPI, cloneServer } from '../services/apiService'; // Import the fetchServerData function
import { removeAuthToken } from '../services/authService';
import { useNavigate } from 'react-router-dom';
import uploadIcon from '../icons/upload.png';
import cloneIcon from '../icons/clone.png';
import executeIcon from '../icons/command.png';
import ExecuteCommandModal from './ExecuteCommandModal'; // Import the modal component
import UploadServerModal from './UploadServerModal'; // Import the modal component

function Dashboard() {
  const navigate = useNavigate();
  const [servers, setServers] = useState([]); // State to hold server data
  const [isExecuteModalOpen, setIsExecuteModalOpen] = useState(false);
  const [executeModalVmId, setExecuteModalVmId] = useState(null); // State to store vmId for the ExecuteCommandModal
  const [isUploadModalOpen, setIsUploadModalOpen] = useState(false);
  const [uploadModalVmId, setUploadModalVmId] = useState(null); // State to store vmId for the ExecuteCommandModal



  useEffect(() => {
    // Fetch server data when the component mounts
    fetchServerData()
      .then((data) => {
        // Set the fetched data in the state
        console.log("**********************************\n", data)
        setServers(data);
      })
      .catch((error) => {
        console.error('Error fetching server data:', error);
        // Handle error if necessary
      });
  }, []); // The empty dependency array means this effect runs once after mounting

  function generateServerCards() {
    return servers.map((server, index) => (
      <div key={index} className={`server-card ${server.status.toLowerCase() === "on" ? "on" : "off"}`}>
        <div className="icon-container">
          {server.status.toLowerCase() === "off" && (
            <img
              className="icon edit-icon"
              src="http://cdn.onlinewebfonts.com/svg/img_354025.png"
              alt={`Edit Server ${index}`}
              onClick={() => handleEditServer(index)}
            />
          )}
          {server.status.toLowerCase() === "on" && (
            <img
              className="icon upload-icon"
              src={uploadIcon}
              alt={`Upload Server ${index}`}
              onClick={() => handleUploadServer(index)}
            />
          )}
          {server.status.toLowerCase() === "off" &&(
            <img
              className="icon clone-icon"
              src={cloneIcon}
              alt={`Clone VM ${index}`}
              onClick={() => handleCloneServer(index)}
            />
          )}
          {server.status.toLowerCase() === "on" && (
            <img
              className="icon execute-icon"
              src={executeIcon}
              alt={`Exec Command ${index}`}
              onClick={() => handleExecuteCommand(index)}
            />
          )}
        </div>
        <h3>{server.name}</h3>
        <p>CPU: {server.cpu}</p>
        <p>RAM: {server.ram}</p>
        <button className="action-button" onClick={() => toggleServerStatus(index)}>
          {server.status.toLowerCase() === "on" ? "Turn Off" : "Turn On"}
        </button>
        {server.status.toLowerCase() === "off" && (
          <button className="delete-button" onClick={() => handleDeleteServer(server.ID)}>
            Delete
          </button>
        )}
      </div>
    ));
  }

  async function handleCloneServer(index) {
    const selectedServer = servers[index];

    try {
      const clonedData = await cloneServer(selectedServer.name);
      if (clonedData) {
        // Create a new server object using the cloned data
        const clonedServer = {
          name: clonedData.name,
          cpu: selectedServer.cpu, // Assuming you want to keep the same CPU as the original server
          ram: selectedServer.ram, // Assuming you want to keep the same RAM as the original server
          status: selectedServer.status, // Status is set to 'off' for the cloned server
        };

        // Update the UI by adding the cloned server to the servers state
        setServers((prevServers) => [...prevServers, clonedServer]);
      } else {
        // Handle cloning failure
        alert('Failed to clone server. Please try again.');
      }
    } catch (error) {
      console.error('Error cloning server:', error);
      // Handle the error gracefully
      alert('An error occurred while cloning the server.');
    }
  }

  function handleExecuteCommand(index) {
    // Get the selected server and its vm_id
    const selectedServer = servers[index];
    const vmId = selectedServer.ID;

    // Open the Execute Command modal and pass vmId as a prop
    setIsExecuteModalOpen(true);
    setExecuteModalVmId(vmId);
  }
  function handleUploadServer(index) {
      // Get the selected server and its vm_id
      const selectedServer = servers[index];
      const vmId = selectedServer.ID;
  
      // Open the Execute Command modal and pass vmId as a prop
      setIsUploadModalOpen(true);
      setUploadModalVmId(vmId); 
  }
  // Function to toggle server status
  function toggleServerStatus(index) {
    const serverToToggle = servers[index];
    const updatedServers = [...servers];

    toggleServerStatusAPI(serverToToggle.ID, serverToToggle.status)
      .then((success) => {
        if (success) {
          // If the update was successful, update the UI
          updatedServers[index].status = serverToToggle.status === "off"? "on": "off";
          setServers(updatedServers);
        } else {
          alert('Error turning on the server. Please try again.');
        }
      })
      .catch((error) => {
        console.error('Error turning on the server:', error);
        alert('An error occurred while turning on the server.');
      });
  }

  // Function to handle editing a server
  function handleEditServer(index) {
    const selectedServer = servers[index];
    const vmDetailsProps = {
      vmID: selectedServer.ID,
      vmName: selectedServer.name,
      ram: selectedServer.ram,
      cpuCores: selectedServer.cpu,
      isEdit: true, // Indicate that this is an edit operation
    };
    navigate("/vm-details", { state: vmDetailsProps });
  }

  function handleDeleteServer(vmId) {
    // Call the deleteServerByVmId function with the vmId
    deleteServerByVmId(vmId)
      .then((success) => {
        if (success) {
          // If the deletion was successful, update the UI by removing the server card
          const updatedServers = servers.filter((server) => server.ID !== vmId);
          setServers(updatedServers);
        } else {
          alert('Error deleting server. Please try again.');
        }
      })
      .catch((error) => {
        console.error('Error deleting server:', error);
        alert('An error occurred while deleting the server.');
      });
  }

  return (
    <div className="container">
      <div className="header">
        <div className="dropdown profile-button">
          <svg className="icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
            <path d="M12 12c2.76 0 5-2.24 5-5s-2.24-5-5-5-5 2.24-5 5 2.24 5 5 5zm0 2c-3.86 0-7 3.14-7 7v1h14v-1c0-3.86-3.14-7-7-7zm0 2c2.67 0 8 1.34 8 4v2H4v-2c0-2.66 5.33-4 8-4z"/>
            <path d="M0 0h24v24H0z" fill="none"/>
          </svg>
          <div className="dropdown-content">
            <button onClick={() => handleDropdownItemClick("profile")}>Profile</button>
            <button onClick={() => handleDropdownItemClick("create-vm")}>Create new VM</button>
            <button onClick={() => handleDropdownItemClick("logout")}>Logout</button>
          </div>
        </div>
        <h2>Welcome to the Dashboard</h2>
      </div>
      <div className="server-list" id="serverList">
        {generateServerCards()}
      </div>
      {isExecuteModalOpen && (
        <ExecuteCommandModal
          onClose={() => setIsExecuteModalOpen(false)}
          onExecute={(command) => {
            // Handle executing the command here
            alert(`Executing command: ${command}`);
          }}
          vmId={executeModalVmId} // Pass the vmId to the modal
        />
      )}
      {isUploadModalOpen && (
        <UploadServerModal
          onClose={() => setIsUploadModalOpen(false)}
          vmId={uploadModalVmId} // Pass the vmId to the modal
        />
      )}
    </div>
  );

  // Function to handle dropdown item click and redirect accordingly
  function handleDropdownItemClick(item) {
    switch (item) {
      case "profile":
        navigate("/profile")
        // Redirect to the profile page
        // You can use React Router for this
        break;
      case "create-vm":
        navigate("/vm-details")
        // Redirect to the create-vm page
        // You can use React Router for this
        break;
      case "logout":
        console.log("Logout kon")
        removeAuthToken()
        navigate("/login")
        break;
      default:
        // Handle other dropdown items if needed
    }
  }
}

export default Dashboard;
