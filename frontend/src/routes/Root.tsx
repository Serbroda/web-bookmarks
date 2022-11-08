import { Outlet, useLoaderData } from "react-router-dom";
import NavItem, { NavItemData } from "../components/NavItem";
import { HomeIcon, NewspaperIcon } from "@heroicons/react/20/solid";
import SideBar from "../components/SideBar";

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
  console.log(groups);

  return (
    <div className="App">
      <div className="flex min-h-screen">
        <SideBar />

        <main className="p-4">
          <Outlet />
        </main>
      </div>
    </div>
  );
};

export default Root;
