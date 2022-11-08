import NavItem, { NavItemData } from "./NavItem";
import { HomeIcon, NewspaperIcon } from "@heroicons/react/20/solid";
import ResizableContainer from "./ResizableContainer";

const navItems: NavItemData[] = [
  { href: "/", label: "Home", active: true, children: [], icon: <HomeIcon /> },
  {
    href: "groups/0",
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
          { href: "#", label: "Third level", active: false, children: [] },
        ],
      },
    ],
  },
];

const SideBar = () => {
  return (
    <ResizableContainer
      width={312}
      conatinerClassName="shrink-0 bg-gray-50 min-w-[256px] max-w-screen-md"
    >
      <div className="flex h-screen">
        <div className="flex flex-col gap-2 justify-items-center w-11 h-screen overflow-x-hidden overflow-y-auto border-r border-gray-200">
          <div className="h-[8px]" />

          <button className="rounded-md flex justify-center items-center mx-1 hover:bg-gray-200 h-8">
            <HomeIcon className="w-4 h-4 m-1 text-gray-700" />
          </button>
          <button className="rounded-md flex justify-center items-center mx-1 hover:bg-gray-200 h-8">
            <HomeIcon className="w-4 h-4 m-1 text-gray-700" />
          </button>
        </div>
        <div className="grow h-screen overflow-x-hidden overflow-y-auto">
          <div className="px-4">
            <a href="#" className="inline-block">
              Logo
            </a>
          </div>

          <nav className="mt-2 px-2 flex-1">
            {navItems.map((item) => (
              <NavItem item={item} />
            ))}
          </nav>
        </div>
      </div>
    </ResizableContainer>
  );
};

export default SideBar;
