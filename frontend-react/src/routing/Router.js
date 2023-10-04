import React from 'react';
import { BrowserRouter as Router } from 'react-router-dom';
import RouteGuard from './RouteGuard'; // Import the new RouteGuard component

function AppRouter() {
  return (
    <Router>
      <RouteGuard /> {/* Wrap your routes with the RouteGuard component */}
    </Router>
  );
}

export default AppRouter