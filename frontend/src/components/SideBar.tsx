import NavItem, { NavItemData } from "./NavItem";
import {
  HomeIcon,
  Cog6ToothIcon,
  NewspaperIcon,
  QuestionMarkCircleIcon,
} from "@heroicons/react/24/outline";
import ResizableContainer from "./ResizableContainer";
import Tippy from "@tippyjs/react";

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
        <div className="flex flex-col gap-1 justify-items-center h-screen overflow-x-hidden overflow-y-auto border-r border-gray-200">
          <div />

          <Tippy content="Home" placement="right">
            <button className="rounded-md flex justify-center items-center mx-1 hover:bg-gray-200 h-7 w-7 cursor-default">
              <HomeIcon className="w-5 h-5 m-1 text-gray-700" />
            </button>
          </Tippy>

          <div className="flex-1"></div>

          <Tippy content="Help" placement="right">
            <button className="rounded-md flex justify-center items-center mx-1 hover:bg-gray-200 h-7 w-7 cursor-default">
              <QuestionMarkCircleIcon className="w-5 h-5 m-1 text-gray-700" />
            </button>
          </Tippy>
          <Tippy content="Settings" placement="right">
            <button className="rounded-md flex justify-center items-center mx-1 hover:bg-gray-200 h-7 w-7 cursor-default">
              <Cog6ToothIcon className="w-5 h-5 m-1 text-gray-700" />
            </button>
          </Tippy>

          <div />
        </div>

        <div className="grow h-screen overflow-x-hidden overflow-y-auto">
          <div className="px-4">
            <a href="#" className="inline-block">
              Logo
            </a>
          </div>

          <nav className="mt-2 px-2 flex-1">
            <div>
              {navItems.map((item) => (
                <NavItem item={item} />
              ))}
            </div>
          </nav>
        </div>
      </div>
    </ResizableContainer>
  );
};

export default SideBar;
