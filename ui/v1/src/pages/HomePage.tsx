import LinkCard from '../components/LinkCard';

const HomePage = () => {
  return (
    <>
      <h1 className="text-xl">Home</h1>
      <LinkCard
        item={{
          title: 'www.test.de',
          url: 'www.test.de',
        }}
        onClick={(item) => console.log(item)}
      />
    </>
  );
};

export default HomePage;
