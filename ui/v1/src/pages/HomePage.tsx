import { Link } from 'react-router-dom';

const HomePage = () => {
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
      <p>Home</p>
    </>
  );
};

export default HomePage;
