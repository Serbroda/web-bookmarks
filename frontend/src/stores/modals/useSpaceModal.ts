import create from 'zustand'

export interface SpaceModalProps {
  onSave: () => void;
}

export type SpaceModalState = {
    isOpen: boolean;
    props: SpaceModalProps;
    setOpen: (value: boolean) => void;
    openModal: (props: SpaceModalProps) => void;
}

const useGroupModal = create<SpaceModalState>((set, get) => ({
    isOpen: false,
    props: {
        onSave: () => {}
    },
    setOpen: (value: boolean) => set({isOpen: value}),
    openModal: (props: SpaceModalProps) => {
      set({isOpen: true, props})
    }
  }));

export default useGroupModal;
