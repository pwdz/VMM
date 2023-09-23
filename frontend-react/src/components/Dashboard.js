import React from 'react';
import './Dashboard.css';
import { useNavigate } from "react-router-dom";

function Dashboard() {
  const navigate = useNavigate();

  // Example data (replace with your backend data)
  const servers = [
    { name: "Server 1", cpu: 4, ram: "8GB", status: "On" },
    { name: "Server 2", cpu: 8, ram: "16GB", status: "Off" },
    { name: "Server 3", cpu: 2, ram: "4GB", status: "On" }
  ];

  // Function to dynamically generate server cards
  function generateServerCards() {
    return servers.map((server, index) => (
      <div key={index} className={`server-card ${server.status === "On" ? "on" : "off"}`}>
        <h3>{server.name}</h3>
        <p>CPU: {server.cpu}</p>
        <p>RAM: {server.ram}</p>
        <button className="action-button" onClick={() => toggleServerStatus(index)}>
          {server.status === "On" ? "Turn Off" : "Turn On"}
        </button>
        <img
          className="edit-icon"
          src="http://cdn.onlinewebfonts.com/svg/img_354025.png"
          alt={`Edit Server ${index}`}
          onClick={() => handleEditServer(index)}
        />
      </div>
    ));
  }

  // Function to toggle server status
  function toggleServerStatus(index) {
    // Replace this with your logic to toggle server status
    const updatedServers = [...servers];
    updatedServers[index].status = updatedServers[index].status === "On" ? "Off" : "On";
    // Update state or send a request to your backend here
  }

  // Function to handle editing a server
  function handleEditServer(index) {
    // Replace this with your logic to handle editing a server
    const selectedServer = servers[index];
    localStorage.setItem("selectedServer", JSON.stringify(selectedServer));
    // Redirect to the appropriate page, e.g., create-vm.html
  }

  return (
    <div className="container">
      <div className="header">
        <div className="dropdown profile-button">
          <svg className="icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
            <path d="M12 12c2.76 0 5-2.24 5-5s-2.24-5-5-5-5 2.24-5 5 2.24 5 5 5zm0 2c-3.86 0-7 3.14-7 7v1h14v-1c0-3.86-3.14-7-7-7zm0 2c2.67 0 8 1.34 8 4v2H4v-2c0-2.66 5.33-4 8-4z"/>
            <path d="M0 0h24v24H0z" fill="none"/>
          </svg>
          <span className="username">John Doe</span>
          <div className="dropdown-content">
            <button onClick={() => handleDropdownItemClick("profile")}>Profile</button>
            <button onClick={() => handleDropdownItemClick("create-vm")}>Create new VM</button>
            <button onClick={() => handleDropdownItemClick("login")}>Logout</button>
          </div>
        </div>
        <h2>Welcome to the Dashboard</h2>
      </div>
      <div className="server-list" id="serverList">
        {generateServerCards()}
      </div>
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
        navigate("/createvm")
        // Redirect to the create-vm page
        // You can use React Router for this
        break;
      case "logout":

        // navigate("/profile")
        // Redirect to the login page
        // You can use React Router for this
        break;
      default:
        // Handle other dropdown items if needed
    }
  }
}

export default Dashboard;
