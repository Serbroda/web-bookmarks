import create from 'zustand'

export interface GroupModalProps {
  onSave: () => void;
}

export type GroupModalState = {
    isOpen: boolean;
    props: GroupModalProps;
    setOpen: (value: boolean) => void;
    openModal: (props: GroupModalProps) => void;
}

const useGroupModal = create<GroupModalState>((set, get) => ({
    isOpen: false,
    props: {
        onSave: () => {}

    },
    setOpen: (value: boolean) => set({isOpen: value}),
    openModal: (props: GroupModalProps) => {
      set({isOpen: true, props})
    }
  }));

export default useGroupModal;