import React, { useState } from 'react';
import './ExecuteCommandModal.css';
import { executeCommandOnVM } from '../services/apiService'; // Import the new function

function ExecuteCommandModal({ onClose, onExecute, vmId }) {
  const [command, setCommand] = useState('');

  const handleExecute = async () => {
    if (command.trim() !== '') {
      try {
        const result = await executeCommandOnVM(vmId, command);
        // Handle the result as needed
        alert(result.message)

        // Close the modal
        onClose();
      } catch (error) {
        console.error('Error executing command:', error);
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
        <h2>Execute Command</h2>
        <input
          type="text"
          className="command-input"
          placeholder="Enter your command here"
          value={command}
          onChange={(e) => setCommand(e.target.value)}
        />
        <button onClick={handleExecute}>Execute</button>
      </div>
    </div>
  );
}

export default ExecuteCommandModal;