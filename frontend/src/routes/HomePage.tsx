import logo from "../assets/logo.svg";

const HomePage = () => {
  return (
    <>
      <div>
        <a href="/" target="_blank">
          <img src={logo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1 className="text-3xl font-bold underline text-red-600">
        Simple React Typescript Tailwind Sample
      </h1>
    </>
  );
};

export default HomePage;
