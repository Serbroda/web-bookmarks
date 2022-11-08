import useMobileNavbar from "../stores/navigation";

const GroupPage = () => {
  const { isOpen, setOpen } = useMobileNavbar();

  return (
    <>
      <button onClick={() => setOpen(true)}>Open nav</button>
      <button onClick={() => setOpen(false)}>Close nav</button>
      <h1>Group</h1>
      <div className="w-64 h-64 bg-green-500 r horizontal rounded-md">Test</div>
    </>
  );
};

export default GroupPage;
