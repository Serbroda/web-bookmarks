import { ExclamationTriangleIcon } from "@heroicons/react/20/solid";
import Modal from "./Modal";
import useLinkModal from "../../stores/modals/useLinkModal";
import useAlertModal from "../../stores/modals/useAlertModal";

const LinkModal = () => {
  const { openModal: openAlertModal } = useAlertModal();
  const { isOpen, props, setOpen } = useLinkModal();

  return (
    <Modal
      show={isOpen}
      onClose={() => setOpen(false)}
      width="medium"
      overflow="visible"
    >
      <div className="sm:flex sm:items-start w-full overflow-visible">
        <div className="mt-3 text-center sm:mt-0 sm:text-left w-full ">
          <div className=" w-full flex flex-col gap-4 ">
            <input
              type="text"
              id="groupName"
              name="groupName"
              className={`input font-bold w-full hover:border-gray-300 hover:shadow-sm ${
                props.mode === "edit" ? "input-ghost w-[97%]" : "mt-6"
              }`}
              placeholder={props.mode === "new" ? "Name" : undefined}
            />

            <a
              className="flex text-sm underline w-fit hover:text-indigo-600 text-left"
              target="_blank"
              href="https://www.youtube.com/watch?v=iG2jotQo9NI&amp;t=17728s"
            >
              <img
                className="favicon pr-1"
                src="https://icons.duckduckgo.com/ip3/www.youtube.com.ico"
                alt="favicon"
              />{" "}
              <span className="break-all">
                https://www.youtube.com/watch?v=iG2jotQo9NI&amp;t=17728s
              </span>{" "}
              <svg
                className="w-4 h-4 ml-1 invisble-hover-item"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"
                ></path>
              </svg>
            </a>

            <select
              id="location"
              name="location"
              className="input mt-1 w-full"
              defaultValue="Canada"
            >
              <option>United States</option>
              <option>Canada</option>
              <option>Mexico</option>
            </select>

            <textarea
              rows={4}
              name="groupDescription"
              id="groupDescription"
              className="block w-full input"
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

        {props.mode === "edit" && (
          <>
            <div className="block sm:flex-1" />
            <button
              type="button"
              className="btn btn-danger w-full mt-4 sm:w-auto sm:mt-0"
              onClick={() => {
                openAlertModal({
                  title: "Delete link",
                  message: "Do you really want to delete this link?",
                  confirmButtonMessage: "Yes, delete link",
                  onConfirm: () => {
                    console.log("Deleted");
                    setOpen(false);
                  },
                });
              }}
              tabIndex={2}
            >
              Delete
            </button>
          </>
        )}
      </div>
    </Modal>
  );
};

export default LinkModal;
