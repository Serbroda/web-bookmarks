import { Dialog, Transition } from "@headlessui/react";
import { FC, Fragment, ReactNode, useState } from "react";

export interface ModalProps {
  children: ReactNode;
  title?: string;
  show?: boolean;
  width?: "small" | "medium" | "big";
  postition?: "top" | "center";
  padding?: boolean;
}

const Modal: FC<ModalProps> = ({
  children,
  title,
  show = false,
  width = "small",
  postition = "center",
  padding = true,
}) => {
  let [isOpen, setIsOpen] = useState(show);

  const closeModal = () => {
    setIsOpen(false);
  };

  const openModal = () => {
    setIsOpen(true);
  };

  return (
    <Transition appear show={show !== undefined ? show : isOpen} as={Fragment}>
      <Dialog as="div" className="relative z-10" onClose={closeModal}>
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
              postition === "center" ? "items-center" : "items-start"
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
                    ? "max-w-[80%]"
                    : "max-w-full"
                } ${
                  padding ? "p-6" : "p-1"
                } w-full max-h-screen transform rounded-2xl bg-white text-left align-middle shadow-xl transition-all`}
              >
                {title && (
                  <Dialog.Title
                    as="h3"
                    className="text-lg font-medium leading-6 text-gray-900 mb-2"
                  >
                    {title}
                  </Dialog.Title>
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
