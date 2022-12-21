import SideNav from "../side-nav/SideNav";
import {Navigate, Outlet, useLocation} from "react-router-dom";
import useAuthentication from "../../stores/useAuthentication";
import {useEffect} from "react";
import useSideNav from "../../stores/useSideNav";
import AlertModal from "../modals/AlertModal";
import SpaceModal from "../modals/SpaceModal";

const ProtectedLayout = () => {
    const location = useLocation();
    const { isOpen, setOpen } = useSideNav();

    useEffect(() => {
        if (isOpen) {
            setOpen(false);
        }
    }, [location]);

    const {authenticated} = useAuthentication();
    if(!authenticated) {
        return <Navigate to="/login" />;
    }

    return (
        <>
            <div className="App bg-gray-100 dark:bg-gray-800">
                <div className="flex h-screen">
                    <SideNav />

                    <main className="h-screen w-full overflow-auto">
                        <Outlet />
                    </main>
                </div>
            </div>

            <AlertModal/>
            <SpaceModal/>
        </>
    )
}

export default ProtectedLayout
