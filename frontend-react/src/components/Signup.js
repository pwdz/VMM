import React, { useState } from 'react';
import { signup } from '../services/apiService'; // Import your signup function from apiService
import { getUserRoleAPI, setAuthToken, setUserRole } from '../services/authService';
import { useNavigate } from 'react-router-dom';

function SignupForm() {
  const navigate = useNavigate();
  
  const [formData, setFormData] = useState({
    username: '',
    password: '',
    email: '',
  });
  const [error, setError] = useState('');

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  async function handleSubmit(event){
    event.preventDefault();
    try {
      // Make the API call to signup
      const data = await signup(formData.username, formData.password, formData.email);
      // Assuming your signup API returns a success flag
      if (data) {
        setAuthToken(data.data);
        const roleData = await getUserRoleAPI()
        if(roleData){
          setUserRole(roleData)
          navigate('/');
        }else{
          alert('Signup failed')
        }

      } else {
        setError('Signup failed. Please check your credentials.');
      }
    } catch (error) {
      setError(''+error.response.data.error);
      console.log(error.response.data.error)
    }
  };

  return (
    <div className="container">
      <h2>Signup</h2>
      {error && <p>{error}</p>}
      <form onSubmit={handleSubmit}> {/* Use onSubmit on the form, not on the button */}
        <div className="form-group">
          <label htmlFor="username">Username:</label>
          <input type="text" name="username" onChange={handleInputChange} />
        </div>
        <div className="form-group">
          <label htmlFor="password">Password:</label>
          <input type="password" name="password" onChange={handleInputChange} />
        </div>
        <div className="form-group">
          <label htmlFor="email">Email:</label>
          <input type="email" name="email" onChange={handleInputChange} />
        </div>
        <div className="form-group">
          <button type="submit"> {/* Change type to "submit" */}
            Signup
          </button>
        </div>
      </form>
    </div>
  );
}

export default SignupForm;

