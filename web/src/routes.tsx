import React from 'react';
import DashboardLayout from './layouts/Dashboard';
import ConfigView from './views/ConfigView';
import MetricsView from './views/MetricsView';

const routes = [
  // {
  //   path: 'app',
  //   element: <DashboardLayout />,
  //   children: [
  //     { path: 'account', element: <AccountView /> },
  //     { path: 'customers', element: <CustomerListView /> },
  //     { path: 'dashboard', element: <DashboardView /> },
  //     { path: 'products', element: <ProductListView /> },
  //     { path: 'settings', element: <SettingsView /> },
  //     { path: '*', element: <Navigate to="/404" /> }
  //   ]
  // },
  {
    path: '/',
    element: <DashboardLayout />,
    children: [
      { path: '/', element: <MetricsView /> },
      { path: 'config', element: <ConfigView /> },
      // { path: 'register', element: <RegisterView /> },
      // { path: '404', element: <NotFoundView /> },
      // { path: '/', element: <Navigate to="/app/dashboard" /> },
      // { path: '*', element: <Navigate to="/404" /> }
    ]
  }
];

export default routes;
