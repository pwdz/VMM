import React from 'react';
import { useRoutes, Navigate } from 'react-router-dom';

// Import your components here
import Dashboard from '../components/Dashboard';
import VMDetails from '../components/VMDetails';
import Profile from '../components/Profile';
import Users from '../components/Users';
import VMs from '../components/VMs';
import Login from '../components/Login';
import Signup from '../components/Signup';
import NotFound from '../components/NotFound';
import Pricing from '../components/Pricing';
import AccessDenied from '../components/AccessDenied';
import { getUserRole } from '../services/authService';

function RouteGuard() {
  const routes = useRoutes([
    {
      path: '/',
      element: !getUserRole() ? <Navigate to="/login" /> : (
        getUserRole() === 'user' ? <Navigate to="/dashboard" /> : (
          getUserRole() === 'admin' ? <Navigate to="/users" /> : <Navigate to="/accessdenied" />
        )
      ),
    },
    {
      path: '/dashboard',
      element: getUserRole() && getUserRole() === 'user' ? <Dashboard /> : <Navigate to="/" />,
    },
    {
      path: '/vm-details',
      element: getUserRole() && getUserRole() === 'user' ? <VMDetails /> : <Navigate to="/" />,
    },
    {
      path: '/profile',
      element: getUserRole() && getUserRole() === 'user' ? <Profile /> : <Navigate to="/" />,
    },
    {
      path: '/users',
      element: getUserRole() && getUserRole() === 'admin' ? <Users /> : <Navigate to="/accessdenied" />,
    },
    {
      path: '/vms',
      element: getUserRole() && getUserRole() === 'admin' ? <VMs /> : <Navigate to="/accessdenied" />,
    },
    {
      path: '/pricing',
      element: getUserRole() && getUserRole() === 'admin' ? <Pricing /> : <Navigate to="/accessdenied" />,
    },
    {
      path: '/login',
      element: !getUserRole() ? <Login /> : <Navigate to="/" />,
    },
    {
      path: '/signup',
      element: !getUserRole() ? <Signup /> : <Navigate to="/" />,
    },
    {
      path: '/accessdenied',
      element: <AccessDenied />,
    },
    {
      path: '*',
      element: <NotFound />,
    },
    // Add other public routes here
  ]);
  return routes;
}

export default RouteGuard;