import React, { useState } from 'react';
import './UploadServerModal.css';
import { uploadServerFile } from '../services/apiService'; // Import the new function

function UploadServerModal({ onClose, vmId }) {
  const [selectedFile, setSelectedFile] = useState(null);
  const [filePath, setFilePath] = useState('');

  const handleFileChange = (event) => {
    const file = event.target.files[0];
    setSelectedFile(file);
  };

  const handleUpload = async () => {
    if (selectedFile && filePath.trim() !== '') {
      try {
        const result = await uploadServerFile(vmId, selectedFile, filePath);
        // Handle the result as needed
        alert(result)
        // Close the modal
        onClose();
      } catch (error) {
        console.error('Error uploading file:', error);
        // Handle the error gracefully
      }
    }
  };

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal" onClick={(e) => e.stopPropagation()}>
        <span className="close-icon" onClick={onClose}>
          &times;
        </span>
        <h2>Upload Server File</h2>
        <input
          type="file"
          className="file-input"
          onChange={handleFileChange}
        />
        <input
          type="text"
          className="file-path-input"
          placeholder="Enter file path"
          value={filePath}
          onChange={(e) => setFilePath(e.target.value)}
        />
        <button onClick={handleUpload}>Upload</button>
      </div>
    </div>
  );
}

export default UploadServerModal;
