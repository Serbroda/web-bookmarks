import { Dialog, Transition } from "@headlessui/react";
import { FC, Fragment, ReactNode, useState } from "react";
import { XMarkIcon } from "@heroicons/react/20/solid";

export interface ModalBaseProps {
  width?: "small" | "medium" | "big";
  postition?: "top" | "center";
  padding?: boolean;
}

export interface ModalProps extends ModalBaseProps {
  children: ReactNode;
  show: boolean;
  showCloseButton?: boolean;
  onClose?: () => void;
  onCloseButtonClick?: () => void;
  onOutsideClick?: () => void;
}

const Modal: FC<ModalProps> = ({
  children,
  show = false,
  width = "small",
  postition = "center",
  padding = true,
  showCloseButton = true,
  onClose,
  onCloseButtonClick,
  onOutsideClick,
}) => {
  return (
    <Transition appear show={show} as={Fragment}>
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
          <div className="fixed inset-0 bg-black/30" aria-hidden="true" />
        </Transition.Child>

        <div className="fixed inset-0">
          <div
            className={`flex min-h-full justify-center p-4 text-center ${
              postition === "center"
                ? "items-center"
                : "items-center md:items-start md:mt-16"
            }`}
          >
            <Transition.Child
              as={Fragment}
              enter="ease-out duration-300"
              enterFrom="opacity-0 scale-95"
              enterTo="opacity-100 scale-100"
              leave="ease-in duration-200"
              leaveFrom="opacity-100 scale-100"
              leaveTo="opacity-0 scale-95"
            >
              <Dialog.Panel
                className={`${
                  width == "small"
                    ? "max-w-md"
                    : width == "medium"
                    ? "max-w-4xl"
                    : "max-w-full"
                } ${
                  padding ? "p-6" : "p-1"
                } w-full max-h-screen transform rounded-2xl bg-white text-left align-middle shadow-xl transition-all`}
              >
                {showCloseButton && (
                  <div className="absolute top-0 right-0 hidden pt-4 pr-4 sm:block">
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
                )}
                {children}
              </Dialog.Panel>
            </Transition.Child>
          </div>
        </div>
      </Dialog>
    </Transition>
  );
};

export default Modal;
