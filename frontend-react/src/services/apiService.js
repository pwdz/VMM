import axios from 'axios';
import { getAuthToken} from './authService';
import { saveAs } from 'file-saver'; // Import the saveAs function


// Create a custom Axios instance with a base URL
const axiosInstance = axios.create({
  baseURL: 'http://localhost:8000', // Replace with your API URL
});

// Add an interceptor to include the JWT token in the headers
axiosInstance.interceptors.request.use(
  (config) => {
    // Retrieve the JWT token from localStorage
    const token = getAuthToken()

    // If the token exists, add it to the Authorization header
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }

    return config;
  },
  (error) => {
    // Handle any request errors here
    return Promise.reject(error);
  }
);

export async function login(username, password) {
  try {
    const response = await axiosInstance.post('/login', {
      username,
      password,
    });

    return response.data;
  } catch (error) {
    throw error;
  }
}

export async function fetchUsers() {
  try {
    const response = await axiosInstance.get('/admin/users');
    console.log(response.data)
    return response.data;
  } catch (error) {
    throw error;
  }
}

export async function exportUsersAsExcel() {
  try {
    const response = await axiosInstance.get('/admin/export-users', {
      responseType: 'blob', // Set the response type to 'blob' to handle binary data
    });

    // Get the content type from the response headers
    const contentType = response.headers['content-type'];

    if (contentType === 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet') {
      // Create a Blob from the binary data
      const blob = new Blob([response.data], { type: contentType });

      // Save the Blob as a file with the name 'users.xlsx'
      saveAs(blob, 'users.xlsx');
    } else {
      console.error('Invalid content type:', contentType);
      throw new Error('Invalid content type');
    }
  } catch (error) {
    console.error('Error exporting users:', error);
    throw error;
  }
}

export async function fetchServerData() {
  try {
    console.log(axiosInstance)
    const response = await axiosInstance.get('/user/get-vms');
    console.log(response.data)
    return response.data;
  } catch (error) {
    throw error;
  }
}

export async function deleteServerByVmId(vmId) {
  try {
    // Make the DELETE request to your DeleteVMHandler API with vm_id in the request body
    console.log(vmId)
    await axiosInstance.post(`/user/delete-vm`, { vm_id: vmId });
    return true; // Return true if the deletion was successful
  } catch (error) {
    console.error('Error deleting server:', error);
    return false; // Return false if there was an error
  }
}

export async function toggleServerStatusAPI(vmId, status) {
  try {
    await axiosInstance.post(status.toLowerCase() === "off"? "/user/power-on-vm": "/user/power-off-vm", { vm_id: vmId });
    return true; // Return true if the deletion was successful
  } catch (error) {
    console.error('Error toggling server status:', error);
    return false; // Return false if there was an error
  }
}

// API call to save changes to VM settings
export async function saveChangesToVM(vmId, newVmName, ramInMb, numCpus) {
  try {
    const response = await axiosInstance.post('/user/change-vm-settings', {
      vm_id: vmId,
      new_vm_name: newVmName,
      ram_in_mb: ramInMb.toString(),
      num_cpus: numCpus.toString(),
    });
    return response.data; // You can return any data you need from the response
  } catch (error) {
    console.error('Error saving changes to VM settings:', error);
    throw error; // You can handle the error as needed
  }
}

// API call to create a new VM
export async function createVM(vmName, ramInMb, numCpus) {
  try {
    const response = await axiosInstance.post('/user/create-vm', {
      vm_name: vmName,
      ram_in_mb: ramInMb,
      num_cpus: numCpus,
      os_type: "ubuntu",
    });
    return response.data; // You can return any data you need from the response
  } catch (error) {
    console.error('Error creating a new VM:', error);
    throw error; // You can handle the error as needed
  }
}
export async function fetchVms() {
  try {
    const response = await axiosInstance.get('/admin/vms');
    console.log(response.data)
    return response.data;
  } catch (error) {
    throw error;
  }
}

export async function exportVmsAsExcel() {
  try {
    const response = await axiosInstance.get('/admin/export-vms', {
      responseType: 'blob', // Set the response type to 'blob' to handle binary data
    });

    // Get the content type from the response headers
    const contentType = response.headers['content-type'];

    if (contentType === 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet') {
      // Create a Blob from the binary data
      const blob = new Blob([response.data], { type: contentType });

      // Save the Blob as a file with the name 'users.xlsx'
      saveAs(blob, 'vms.xlsx');
    } else {
      console.error('Invalid content type:', contentType);
      throw new Error('Invalid content type');
    }
  } catch (error) {
    console.error('Error exporting users:', error);
    throw error;
  }
}
export async function fetchUserProfile() {
  try {
    const response = await axiosInstance.get('/user/profile');
    return response.data;
  } catch (error) {
    console.error('Error fetching user profile:', error);
    throw error;
  }
}
export async function cloneServer(vmName) {
  try {
    const response = await axiosInstance.post('/user/clone-vm', {
      vm_name: vmName,
      new_vm_name: vmName + " Clone"
    });
    return response.data; // You can return any data you need from the response
  } catch (error) {
    console.error('Error cloning server:', error);
    throw error; // You can handle the error as needed
  }
}
export async function executeCommandOnVM(vmId, command) {
  try {
    const response = await axiosInstance.post('/user/execute-command-on-vm', {
      vm_id: vmId,
      command: command,
    });

    // Handle the response as needed
    return response.data;
  } catch (error) {
    console.error('Error executing command on VM:', error);
    throw error; // Handle the error as needed
  }
}
// Import any necessary libraries/modules

export async function uploadServerFile(vmId, file, filePath) {
  return new Promise(async (resolve, reject) => {
    try {
      // Convert the file content to Base64 encoding
      const fileReader = new FileReader();
      fileReader.readAsDataURL(file);

      fileReader.onload = async () => {
        const fileContentBase64 = fileReader.result.split(',')[1]; // Extract Base64 content

        const response = await axiosInstance.post('/user/upload-file-to-vm', {
          vm_id: vmId,
          file_content: fileContentBase64,
          guest_file_path: filePath,
        });

        // Check the response status code and handle it as needed
        if (response.status === 200) {
          resolve({ success: true, data: response.data });
        } else if (response.status === 500) {
          // Handle the 500 status code
          console.error('Server returned a 500 status code.');
          reject({ success: false, error: 'Internal server error' });
        } else {
          // Handle other error status codes
          console.error('Server returned an error status code:', response.status);
          reject({ success: false, error: 'File upload failed' });
        }
      };

      fileReader.onerror = (error) => {
        console.error('Error reading file content:', error);
        reject(error);
      };
    } catch (error) {
      console.error('Error uploading server file:', error);
      reject(error);
    }
  });
}
export async function signup(username, password, email) {
  try {
    const response = await axiosInstance.post('/signup', {
      username,
      email,
      password
    });
    return response.data;
  } catch (error) {
    throw error;
  }
}

export async function fetchPricingData() {
  try {
    const response = await axiosInstance.get('/admin/pricing');
    console.log(response.data)
    return response.data;
  } catch (error) {
    console.error('Error fetching pricing data:', error);
    throw error;
  }
}

export async function updatePricingData(priceConfigs) {
  try {
    console.log(priceConfigs)
    const response = await axiosInstance.post('/admin/pricing-update', priceConfigs);
    return response.data;
  } catch (error) {
    console.error('Error updating pricing data:', error);
    throw error;
  }
}
