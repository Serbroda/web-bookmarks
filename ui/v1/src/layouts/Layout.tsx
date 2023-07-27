import SideNav from '@components/side-nav/SideNav';
import { Navigate, Outlet, useLocation } from 'react-router-dom';
import { useEffect } from 'react';
import useSideNav from '@stores/useSideNav';

const Layout = () => {
  const location = useLocation();
  const { isOpen, setOpen } = useSideNav();

  useEffect(() => {
    if (isOpen) {
      setOpen(false);
    }
  }, [location]);

  return (
    <div className="h-screen flex">
      <SideNav />

      <main className="flex-1 h-full max-h-screen overflow-y-auto">
        <Outlet />
      </main>
    </div>
  );
};

export default Layout;
