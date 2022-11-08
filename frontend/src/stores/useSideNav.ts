import create from 'zustand'

export interface UseMovileNavbarState {
    isOpen: boolean;
    setOpen: (value: boolean) => void;
}

const useSideNav = create<UseMovileNavbarState>((set) => ({
    isOpen: true,
    setOpen: (value: boolean) => set({isOpen: value}),
  }));

export default useSideNav;