import {Navigate, Outlet} from "react-router-dom";
import useAuthentication from "../../stores/useAuthentication";

const Layout = () => {
    const {authenticated} = useAuthentication();
    if (authenticated) {
        return <Navigate to="/"/>;
    }

    return (
        <>
            <Outlet/>
        </>
    )
}

export default Layout
