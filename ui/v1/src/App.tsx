import { Link, Outlet } from 'react-router-dom';

const App = () => {
  return (
    <>
      <ol>
        <li>
          <Link to="/">Home</Link>
        </li>
        <li>
          <Link to="/settings">Settings</Link>
        </li>
        <li>
          <Link to="/about">About</Link>
        </li>
      </ol>
      <Outlet />
    </>
  );
};

export default App;
