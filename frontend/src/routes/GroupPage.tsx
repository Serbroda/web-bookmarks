import useSideNav from "../stores/useSideNav";

const GroupPage = () => {
  const { isOpen, setOpen } = useSideNav();

  return (
    <>
      <button onClick={() => setOpen(true)}>Open nav</button>
      <h1>Group</h1>
      <div className="w-64 h-64 bg-green-500 r horizontal rounded-md">Test</div>
    </>
  );
};

export default GroupPage;
