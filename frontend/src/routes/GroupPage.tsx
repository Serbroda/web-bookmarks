import GroupModal from "../components/modals/GroupModal";
import LinkModal from "../components/modals/LinkModal";
import useAlertModal from "../stores/modals/useAlertModal";
import useGroupModal from "../stores/modals/useGroupModal";
import useLinkModal from "../stores/modals/useLinkModal";
import useSideNav from "../stores/useSideNav";

const GroupPage = () => {
  const { isOpen, setOpen } = useSideNav();
  const { openModal: openAlertModal } = useAlertModal();
  const { openModal: openGroupModal } = useGroupModal();
  const { openModal: openLinkModal } = useLinkModal();

  return (
    <>
      <button onClick={() => setOpen(true)}>Open nav</button>
      <h1>Group</h1>
      <div className="w-64 h-64 bg-green-500 r horizontal rounded-md">Test</div>

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

      <button
        className="btn btn-primary"
        onClick={() =>
          openGroupModal({
            onSave: () => {
              console.log("Saved");
            },
          })
        }
      >
        Add group
      </button>

      <button
        className="btn btn-primary"
        onClick={() =>
          openLinkModal({
            mode: "edit",
            onSave: () => {
              console.log("Link saved");
            },
          })
        }
      >
        Add Link
      </button>

      <br />
      <span className="btn-group">
        <button className="btn btn-xs">xs</button>
        <button className="btn btn-primary btn-xs">xs</button>
        <button className="btn btn-primary btn-xs">xs</button>
      </span>

      <br />
      <input type="text" className="input" />
      <input type="text" className="input input-error" />

      <GroupModal />
      <LinkModal />
    </>
  );
};

export default GroupPage;
