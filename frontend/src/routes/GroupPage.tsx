import MyModal from "../components/modals/AlertModal";
import useAlertModal from "../stores/modals/useAlertModal";
import useSideNav from "../stores/useSideNav";

const GroupPage = () => {
  const { isOpen, setOpen } = useSideNav();
  const { openModal: openAlertModal } = useAlertModal();

  return (
    <>
      <button onClick={() => setOpen(true)}>Open nav</button>
      <h1>Group</h1>
      <div className="w-64 h-64 bg-green-500 r horizontal rounded-md">Test</div>

      <MyModal />

      <button
        className="btn btn-danger"
        onClick={() =>
          openAlertModal({
            title: "Delete Account",
            message: "Do you really want to delete this account?",
            confirmButtonMessage: "Yes, delete account",
            onConfirm: () => {
              console.log("Confirmed account deletion");
            },
          })
        }
      >
        Delete
      </button>

      <br />
      <span className="btn-group">
        <button className="btn btn-xs">xs</button>
        <button className="btn btn-primary btn-xs">xs</button>
        <button className="btn btn-primary btn-xs">xs</button>
      </span>
    </>
  );
};

export default GroupPage;
