import { useEffect } from "react";
import { Outlet, useLoaderData, useLocation } from "react-router-dom";
import SideNav from "../components/side-nav/SideNav";
import usePreferences from "../stores/usePreferences";
import useSideNav from "../stores/useSideNav";
import AlertModal from "../components/modals/AlertModal";
import TopNav from "../components/TopNav";
import { authApi } from "../services/api.service";

export interface GroupDto {
  id: number;
  name: string;
}

export const loader = async () => {
  return {
    groups: [
      { id: 1, name: "Rezepte" },
      { id: 2, name: "Videos" },
    ],
  };
};

const Root = () => {
  const { groups } = useLoaderData() as { groups: GroupDto[] };
  const { init } = usePreferences();

  const location = useLocation();
  const { isOpen, setOpen } = useSideNav();

  useEffect(() => {
    authApi
      .login({
        loginDto: { username: "danny@rottstegge.net", password: "mekahesh9*" },
      })
      .then((r) => console.log(r))
      .catch((e) => console.error(e));
  }, []);

  useEffect(() => {
    init();
  }, [init]);

  useEffect(() => {
    if (isOpen) {
      setOpen(false);
    }
  }, [location]);

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

      <AlertModal />
    </>
  );
};

export default Root;
