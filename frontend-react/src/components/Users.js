import React, { useEffect, useState } from 'react';
import axiosInstance from '../services/apiService'; // Import your custom Axios instance

function UsersView() {
  const [users, setUsers] = useState([]);
  const [exporting, setExporting] = useState(false);

  // Function to fetch users from the API
  const fetchUsers = () => {
    axiosInstance.get('/admin/users')
      .then((response) => {
        setUsers(response.data);
      })
      .catch((error) => {
        console.error('Error fetching users:', error);
      });
  };

  // Function to export users as Excel
  const exportUsers = () => {
    setExporting(true);

    // Make a request to export users as Excel
    axiosInstance.get('/admin/export-users')
      .then((response) => {
        // Handle the response, e.g., initiate a download
        // This depends on how your server responds with the Excel file
        // You may need to implement a file download feature
        // Example: window.location.href = response.data.downloadLink;
      })
      .catch((error) => {
        console.error('Error exporting users:', error);
      })
      .finally(() => {
        setExporting(false);
      });
  };

  // Fetch users when the component mounts
  useEffect(() => {
    fetchUsers();
  }, []);

  return (
    <div>
      <h1>Users</h1>

      {/* Export button */}
      <button onClick={exportUsers} disabled={exporting}>
        {exporting ? 'Exporting...' : 'Export Users'}
      </button>

      {/* Display the list of users */}
      <ul>
        {users.map((user) => (
          <li key={user.ID}>{user.Username}</li>
        ))}
      </ul>
    </div>
  );
}

export default UsersView;
