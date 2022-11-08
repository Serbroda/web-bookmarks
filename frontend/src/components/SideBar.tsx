import NavItem, { NavItemData } from "./NavItem";
import { HomeIcon, NewspaperIcon } from "@heroicons/react/20/solid";
import ResizableContainer from "./ResizableContainer";

const navItems: NavItemData[] = [
  { href: "/", label: "Home", active: true, children: [], icon: <HomeIcon /> },
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
          { href: "#", label: "Third level", active: false, children: [] },
        ],
      },
    ],
  },
];

const SideBar = () => {
  return (
    <ResizableContainer conatinerClassName="shrink-0 bg-gray-50 min-w-[200px] max-w-screen-md">
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
    </ResizableContainer>
  );
};

export default SideBar;
