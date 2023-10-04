import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { fetchVms, exportVmsAsExcel } from '../services/apiService';
import './VMs.css'; // Import your CSS file for styling
import { FaDownload } from 'react-icons/fa'; // Import the download icon from a library

function Vms() {
  const [vms, setVms] = useState([]);
  const [exporting, setExporting] = useState(false);
  const navigate = useNavigate(); // Create a navigate function

  // Function to fetch VMs from the API
  const fetchVMSData = async () => {
    try {
      const vmData = await fetchVms();
      setVms(vmData);
    } catch (error) {
      console.error('Error fetching VMs:', error);
    }
  };

  // Function to export VMs as Excel
  const handleExportVMs = async () => {
    setExporting(true);

    try {
      await exportVmsAsExcel();
    } catch (error) {
      console.error('Error exporting VMs:', error);
    } finally {
      setExporting(false);
    }
  };

  // Function to navigate to the "Users" page
  const navigateToUsers = () => {
    navigate('/users');
  };

  // Fetch VMs when the component mounts
  useEffect(() => {
    fetchVMSData();
  }, []);

  return (
    <div className="vms-container">
     <div className="header-container">
        <h1 className="header">VMs</h1>
        <button className="users-button" onClick={navigateToUsers}>
          [Go to Users]
        </button>
      </div> 

      {/* Export button with download icon */}
      <button className="export-button" onClick={handleExportVMs} disabled={exporting}>
        <FaDownload className="download-icon" />
        {exporting ? 'Exporting...' : 'Export VMs'}
      </button>

      {/* Display the list of VMs */}
      <table className="vm-table">
        <thead>
          <tr>
            <th>User ID</th>
            <th>Username</th>
            <th>VM ID</th>
            <th>VM Name</th>
            <th>OS Type</th>
            <th>RAM</th>
            <th>CPU</th>
            <th>Status</th>
            <th>Is Deleted</th>
          </tr>
        </thead>
        <tbody>
          {vms.map((vm, index) => (
            <tr key={vm.id} className={index % 2 === 0 ? 'even-row' : 'odd-row'}>
              <td>{vm.user_id}</td>
              <td>{vm.username}</td>
              <td>{vm.ID}</td>
              <td>{vm.name}</td>
              <td>{vm.os_type}</td>
              <td>{vm.ram}</td>
              <td>{vm.cpu}</td>
              <td>{vm.status}</td>
              <td>{vm.is_deleted ? 'Yes' : 'No'}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default Vms;
