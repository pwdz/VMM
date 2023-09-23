import React, { useEffect, useState } from 'react';
import { useRoutes, Outlet, Navigate } from 'react-router-dom';
import * as authService from './services/authService';

// Import your components here
import Dashboard from './components/Dashboard';
import CreateVM from './components/CreateVM';
import Profile from './components/Profile';
import Users from './components/Users';
import VMs from './components/VMs';
import Login from './components/Login';
import NotFound from './components/NotFound';
import AccessDenied from './components/AccessDenied';

function RouteGuard() {
  const [userRole, setuserRole] = useState(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    // Make an API call to get user data, including authentication and role
    authService.getUserRole()
      .then((data) => {
        setuserRole(data);
        setIsLoading(false);
      })
      .catch((error) => {
        console.error('Error fetching user data:', error);
        setIsLoading(false);
      });
  }, []);

  const routes = useRoutes([
    {
      path: '/',
      element: isLoading ? <div>Loading...</div> : !userRole ? <Navigate to="/login" /> : (
        userRole.role === 'user' ? <Navigate to="/dashboard" /> : (
          userRole.role === 'admin' ? <Navigate to="/users" /> : <Navigate to="/accessdenied" />
        )
      ),
    },
    {
      path: '/dashboard',
      element: userRole && userRole.role === 'user' ? <Dashboard /> : <Navigate to="/" />,
    },
    {
      path: '/createvm',
      element: userRole && userRole.role === 'user' ? <CreateVM /> : <Navigate to="/" />,
    },
    {
      path: '/profile',
      element: userRole && userRole.role === 'user' ? <Profile /> : <Navigate to="/" />,
    },
    {
      path: '/users',
      element: userRole && userRole.role === 'admin' ? <Users /> : <Navigate to="/accessdenied" />,
    },
    {
      path: '/vms',
      element: userRole && userRole.role === 'admin' ? <VMs /> : <Navigate to="/accessdenied" />,
    },
    {
      path: '/login',
      element: !userRole ? <Login /> : <Navigate to="/" />,
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