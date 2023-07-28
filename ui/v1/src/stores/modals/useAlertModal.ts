import create from 'zustand'

export interface AlertModalProps {
  title: string;
  message: string;
  confirmButtonMessage: string;
  onConfirm: () => void;
}

export type AlertModalState = {
    isOpen: boolean;
    props: AlertModalProps;
    setOpen: (value: boolean) => void;
    openModal: (props: AlertModalProps) => void;
}

const useAlertModal = create<AlertModalState>((set, get) => ({
    isOpen: false,
    props: {
        title: "Confirm",
        message: "",
        confirmButtonMessage: "Yes",
        onConfirm: () => {}

    },
    setOpen: (value: boolean) => set({isOpen: value}),
    openModal: (props: AlertModalProps) => {
      set({isOpen: true, props})
    }
  }));

export default useAlertModal;