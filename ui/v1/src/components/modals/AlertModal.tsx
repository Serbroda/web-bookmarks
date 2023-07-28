import { Dialog } from '@headlessui/react';
import { ExclamationTriangleIcon } from '@heroicons/react/20/solid';
import Modal from './Modal';
import useAlertModal from '../../stores/modals/useAlertModal';

const AlertModal = () => {
  const { isOpen, props, setOpen } = useAlertModal();

  return (
    <Modal show={isOpen} onClose={() => setOpen(false)}>
      <div className="sm:flex sm:items-start">
        <div className="mx-auto flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full bg-red-100 sm:mx-0 sm:h-10 sm:w-10">
          <ExclamationTriangleIcon className="h-6 w-6 text-red-600" aria-hidden="true" />
        </div>
        <div className="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
          <Dialog.Title as="h3" className="text-lg font-medium leading-6 text-gray-900">
            {props.title}
          </Dialog.Title>
          <div className="mt-2">
            <p className="text-sm text-gray-500">{props.message}</p>
          </div>
        </div>
      </div>
      <div className="mt-5 sm:mt-4 sm:flex sm:flex-row-reverse">
        <button
          type="button"
          className="btn btn-danger w-full sm:ml-3 sm:w-auto"
          onClick={() => {
            if (props.onConfirm) {
              props.onConfirm();
            }
            setOpen(false);
          }}
          tabIndex={2}>
          {props.confirmButtonMessage}
        </button>
        <button
          type="button"
          className="btn mt-4 w-full sm:mt-0 sm:w-auto sm:text-sm"
          onClick={() => {
            setOpen(false);
          }}
          autoFocus
          tabIndex={1}>
          Cancel
        </button>
      </div>
    </Modal>
  );
};

export default AlertModal;
