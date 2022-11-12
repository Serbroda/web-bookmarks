import create from 'zustand'

export interface LinkModalProps {
  mode: "new" | "edit";
  onSave: () => void;
}

export type LinkModalState = {
    isOpen: boolean;
    props: LinkModalProps;
    setOpen: (value: boolean) => void;
    openModal: (props: LinkModalProps) => void;
}

const useLinkModal = create<LinkModalState>((set, get) => ({
    isOpen: false,
    props: {
        mode: "new",
        onSave: () => {}

    },
    setOpen: (value: boolean) => set({isOpen: value}),
    openModal: (props: LinkModalProps) => {
      set({isOpen: true, props})
    }
  }));

export default useLinkModal;