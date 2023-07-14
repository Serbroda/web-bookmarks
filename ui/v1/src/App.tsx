import HomePage from './pages/HomePage';
import { createHashRouter, RouterProvider } from 'react-router-dom';
import SettingsPage from './pages/SettingsPage';

const router = createHashRouter([
  {
    path: '/',
    Component: HomePage,
  },
  {
    path: 'settings',
    Component: SettingsPage,
  },
  {
    path: 'about',
    element: <div>About</div>,
  },
]);

const App = () => {
  return <RouterProvider router={router} />;
};

export default App;
