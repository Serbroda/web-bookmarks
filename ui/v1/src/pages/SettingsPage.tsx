import { Link } from 'react-router-dom';

const SettingsPage = () => {
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
      <p>Settings</p>
    </>
  );
};

export default SettingsPage;
