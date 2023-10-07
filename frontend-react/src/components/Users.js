import React, { useEffect, useState } from 'react';
import { fetchUsers, exportUsersAsExcel } from '../services/apiService';
import './Users.css'; // Import your CSS file for styling
import { FaDownload } from 'react-icons/fa'; // Import the download icon from a library
import { useNavigate } from 'react-router-dom';

function UsersView() {
  const [users, setUsers] = useState([]);
  const [exporting, setExporting] = useState(false);
  const navigate = useNavigate()

  // Function to fetch users from the API
  const fetchUsersData = async () => {
    try {
      const userData = await fetchUsers();
      setUsers(userData);
    } catch (error) {
      console.error('Error fetching users:', error);
    }
  };

  // Function to export users as Excel
  const handleExportUsers = async () => {
    setExporting(true);

    try {
      await exportUsersAsExcel();
      exportUsersAsExcel()
    } catch (error) {
      console.error('Error exporting users:', error);
    } finally {
      setExporting(false);
    }
  };

  // Fetch users when the component mounts
  useEffect(() => {
    fetchUsersData();
  }, []);
  

  const navigateToVMs = () => {
    navigate('/vms');
  };
  
  return (
    <div className="users-container">
           <div className="header-container">
        <h1 className="header">Users</h1>
        <button className="vms-button" onClick={navigateToVMs}>
          [Go to VMs]
        </button>
      </div> 

      {/* Export button with download icon */}
      <button className="export-button" onClick={handleExportUsers} disabled={exporting}>
        <FaDownload className="download-icon" />
        {exporting ? 'Exporting...' : 'Export Users'}
      </button>

      {/* Display the list of users */}
      <table className="user-table">
        <thead>
          <tr>
            <th>ID</th>
            <th>Username</th>
            <th>Email</th>
            <th>Role</th>
            <th>Active VM Count</th>
            <th>Inactive VM Count</th>
            <th>Total Cost</th>
          </tr>
        </thead>
        <tbody>
          {users.map((user, index) => (
            <tr key={user.id} className={index % 2 === 0 ? 'even-row' : 'odd-row'}>
              <td>{user.ID}</td>
              <td>{user.username}</td>
              <td>{user.email}</td>
              <td>{user.role}</td>
              <td>{user.active_vm_count}</td>
              <td>{user.inactive_vm_count}</td>
              <td>{user.total_cost}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default UsersView;
