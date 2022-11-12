import { Dialog } from "@headlessui/react";
import { ExclamationTriangleIcon } from "@heroicons/react/20/solid";
import Modal from "./Modal";
import useGroupModal from "../../stores/modals/useGroupModal";

const GroupModal = () => {
  const { isOpen, props, setOpen } = useGroupModal();

  return (
    <Modal show={isOpen} onClose={() => setOpen(false)} width="medium">
      <div className="sm:flex sm:items-start w-full">
        <div className="mt-3 text-center sm:mt-0 sm:text-left w-full">
          <Dialog.Title
            as="h3"
            className="text-lg font-medium leading-6 text-gray-900"
          >
            Add Group
          </Dialog.Title>

          <div className="mt-2 w-full">
            <input
              type="text"
              id="groupName"
              name="groupName"
              className="input w-full"
              placeholder="Name"
            />
            <textarea
              rows={4}
              name="groupDescription"
              id="groupDescription"
              className="block w-full input mt-4"
              placeholder="Description (optional)"
              defaultValue={""}
            />
          </div>
        </div>
      </div>

      <div className="mt-5 sm:mt-4 sm:flex sm:flex-row-reverse">
        <button
          type="button"
          className="btn btn-primary w-full sm:ml-3 sm:w-auto"
          onClick={() => {
            if (props.onSave) {
              props.onSave();
            }
            setOpen(false);
          }}
          tabIndex={2}
        >
          Save
        </button>
        <button
          type="button"
          className="btn mt-4 w-full sm:mt-0 sm:w-auto sm:text-sm"
          onClick={() => {
            setOpen(false);
          }}
          autoFocus
          tabIndex={1}
        >
          Cancel
        </button>
      </div>
    </Modal>
  );
};

export default GroupModal;
