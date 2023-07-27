import Layout from '@layouts/Layout';
import ErrorPage from '@pages/ErrorPage';
import GroupPage from '@pages/GroupPage';
import HomePage from '@pages/HomePage';
import SettingsPage from '@pages/SettingsPage';
import { Route, RouterProvider, createHashRouter, createRoutesFromElements } from 'react-router-dom';

const router = createHashRouter(
  createRoutesFromElements(
    <>
      <Route element={<Layout />} errorElement={<ErrorPage />}>
        <Route path="/" element={<HomePage />} />
        <Route path="/groups/:id" element={<GroupPage />} />
        <Route path="/settings" element={<SettingsPage />} />
      </Route>
    </>
  )
);

const App = () => {
  return <RouterProvider router={router} />;
};

export default App;
