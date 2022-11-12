import { Dialog, Transition } from "@headlessui/react";
import { FC, Fragment, ReactNode } from "react";
import { XMarkIcon } from "@heroicons/react/20/solid";

export interface ModalProps {
  children: ReactNode;
  show: boolean;
  width?: "small" | "medium" | "big";
  overflow?: "hidden" | "visible";
  onClose?: () => void;
  onCloseButtonClick?: () => void;
  onOutsideClick?: () => void;
}

const Modal: FC<ModalProps> = ({
  children,
  show = false,
  width = "small",
  overflow = "hidden",
  onClose,
  onCloseButtonClick,
  onOutsideClick,
}) => {
  return (
    <Transition.Root show={show} as={Fragment}>
      <Dialog
        as="div"
        className="relative z-10"
        onClose={() => {
          if (onOutsideClick) {
            onOutsideClick();
          }
          if (onClose) {
            onClose();
          }
        }}
      >
        <Transition.Child
          as={Fragment}
          enter="ease-out duration-300"
          enterFrom="opacity-0"
          enterTo="opacity-100"
          leave="ease-in duration-200"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
        >
          <div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
        </Transition.Child>

        <div className="fixed inset-0 z-10 overflow-y-auto">
          <div
            className={`flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0`}
          >
            <Transition.Child
              as={Fragment}
              enter="ease-out duration-300"
              enterFrom="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
              enterTo="opacity-100 translate-y-0 sm:scale-100"
              leave="ease-in duration-200"
              leaveFrom="opacity-100 translate-y-0 sm:scale-100"
              leaveTo="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
            >
              <Dialog.Panel
                className={`relative w-full transform rounded-lg bg-white px-4 pt-5 pb-4 text-left shadow-xl transition-all sm:my-8 sm:p-6 ${
                  width == "small"
                    ? "max-w-lg"
                    : width == "medium"
                    ? "max-w-4xl"
                    : "max-w-7xl"
                } overflow-${overflow}`}
              >
                <div className="absolute top-0 right-0 pt-4 pr-4 sm:block">
                  <button
                    type="button"
                    className="rounded-md bg-white text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                    onClick={() => {
                      if (onCloseButtonClick) {
                        onCloseButtonClick();
                      }
                      if (onClose) {
                        onClose();
                      }
                    }}
                  >
                    <span className="sr-only">Close</span>
                    <XMarkIcon className="h-6 w-6" aria-hidden="true" />
                  </button>
                </div>

                {children}
              </Dialog.Panel>
            </Transition.Child>
          </div>
        </div>
      </Dialog>
    </Transition.Root>
  );
};

export default Modal;
