import create from 'zustand'

export type UseMovileNavbarState = {
    isOpen: boolean;
    setOpen: (value: boolean) => void;
    toggle: () => void;
}

const useSideNav = create<UseMovileNavbarState>((set, get) => ({
    isOpen: true,
    setOpen: (value: boolean) => set({isOpen: value}),
    toggle: () => set({isOpen: !get().isOpen})
  }));

export default useSideNav;