import create from 'zustand';

export interface LinkModalProps {
  mode: 'new' | 'edit';
  url?: string;
  onSave: () => void;
}

export type LinkModalState = {
  isOpen: boolean;
  props: LinkModalProps;
  setOpen: (value: boolean) => void;
  openModal: (props: LinkModalProps) => void;
};

const useLinkModal = create<LinkModalState>((set) => ({
  isOpen: false,
  props: {
    mode: 'new',
    url: '',
    onSave: () => {},
  },
  setOpen: (value: boolean) => set({ isOpen: value }),
  openModal: (props: LinkModalProps) => {
    set({ isOpen: true, props });
  },
}));

export default useLinkModal;
