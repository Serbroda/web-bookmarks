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

const navItems: NavItemData[] = [
  { href: "/", label: "Home", active: false, children: [], icon: <HomeIcon /> },
  {
    href: "#",
    label: "Groups",
    active: false,
    children: [
      { href: "groups/1", label: "All posts", active: false, children: [] },
      { href: "groups/2", label: "Add new", active: false, children: [] },
      { href: "groups/3", label: "Categories", active: false, children: [] },
    ],
    icon: <NewspaperIcon />,
  },
  {
    href: "#",
    label: "Media",
    active: false,
    children: [
      { href: "#", label: "Library", active: false, children: [] },
      {
        href: "#",
        label: "Add new",
        active: false,
        children: [
          { href: "#", label: "Third level", active: true, children: [] },
        ],
      },
    ],
  },
];

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