import create from 'zustand'

const useMobileNavbar = create((set) => ({
    isOpen: true,
    setOpen: (value: boolean) => set({isOpen: value}),
  }));

export default useMobileNavbar;