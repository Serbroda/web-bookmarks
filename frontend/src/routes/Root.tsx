import { useEffect } from "react";
import {
  Outlet,
  useLoaderData,
  useLocation,
  useNavigation,
} from "react-router-dom";
import SideNav from "../components/side-nav/SideNav";
import useSideNav from "../stores/useSideNav";

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

  const location = useLocation();
  const { isOpen, setOpen } = useSideNav();

  useEffect(() => {
    if (isOpen) {
      setOpen(false);
    }
  }, [location]);

  return (
    <div className="App">
      <div className="flex h-screen">
        <SideNav />

        <main className="p-4 h-screen w-full overflow-auto">
          <Outlet />
        </main>
      </div>
    </div>
  );
};

export default Root;
