import NavItem, { NavItemData } from "./NavItem";
import {
  HomeIcon,
  Cog6ToothIcon,
  NewspaperIcon,
  QuestionMarkCircleIcon,
  PencilSquareIcon,
  FolderPlusIcon,
  PlusIcon,
} from "@heroicons/react/24/outline";
import ResizableContainer from "./ResizableContainer";
import Tippy from "@tippyjs/react";
import { ReactNode } from "react";
import Logo from "../assets/logo.svg";

const navItems: NavItemData[] = [
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

interface SideSideBarItem {
  content: ReactNode;
  active: boolean;
  tooltip?: string;
  href?: string;
  onClick?: () => void;
}

const topSideDideBarItems: SideSideBarItem[] = [
  {
    tooltip: "Home",
    content: <HomeIcon className="w-5 h-5 m-1 text-gray-700" />,
    active: true,
  },
  {
    tooltip: "New Space",
    content: <PlusIcon className="w-5 h-5 m-1 text-purple-700" />,
    active: false,
  },
];

const bottomSideDideBarItems: SideSideBarItem[] = [
  {
    tooltip: "Help",
    content: <QuestionMarkCircleIcon className="w-5 h-5 m-1 text-gray-700" />,
    active: false,
  },
  {
    tooltip: "Settings",
    content: <Cog6ToothIcon className="w-5 h-5 m-1 text-gray-700" />,
    active: false,
  },
];

const SideBar = () => {
  const sideSideBarItem = (item: SideSideBarItem) => {
    const btn = (
      <button
        className={`${
          item.active ? "bg-gray-200" : ""
        } rounded-full flex justify-center items-center mx-1 hover:bg-gray-200 h-7 w-7 cursor-default`}
      >
        {item.content}
      </button>
    );

    return (
      <>
        {item.tooltip ? (
          <Tippy content={item.tooltip} placement="right">
            {btn}
          </Tippy>
        ) : (
          { btn }
        )}
      </>
    );
  };
  return (
    <ResizableContainer
      width={312}
      conatinerClassName="shrink-0 bg-gray-50 min-w-[256px] max-w-[80%]"
    >
      <div className="flex h-full">
        <div className="flex flex-col gap-0.5 justify-items-center border-r border-gray-200 overflow-x-hidden overflow-y-auto">
          <div />
          {topSideDideBarItems.map((item) => sideSideBarItem(item))}
          <div className="flex-1" />
          {bottomSideDideBarItems.map((item) => sideSideBarItem(item))}
          <div />
        </div>

        <div className="flex flex-col h-full w-full">
          <div className="px-4 sticky">
            <a
              href="/"
              className="px-2 flex-0 btn btn-ghost md:px-4"
              aria-label="Homepage"
            >
              <div className="flex inline-block text-3xl font-semibold">
                <img src={Logo} className="w-9 h-9 mr-3" alt="Logo" />{" "}
                <span className="lowercase  text-red-700">rag</span>
                <span className="uppercase text-gray-700">bag</span>
              </div>
            </a>
          </div>

          <nav className="flex flex-col mt-2 px-2 h-full overflow-x-hidden overflow-y-auto">
            <div>
              <NavItem
                item={{
                  href: "/",
                  label: "Home",
                  active: true,
                  children: [],
                  icon: <HomeIcon />,
                }}
              />

              <h2 className="w-full py-4 px-1 text-sm font-semibold text-gray-400">
                My Groups
              </h2>

              {navItems.map((item) => (
                <NavItem item={item} />
              ))}
            </div>
          </nav>

          <footer className="sticky inset-x-0 bottom-0 bg-base-200 border-t border-base-100 py-1">
            <div className="flex gap-0.5 justify-center">
              <Tippy content="New Group" placement="bottom">
                <button className="rounded-md flex justify-center items-center hover:bg-gray-200 h-7 w-7 cursor-default">
                  <FolderPlusIcon className="w-5 h-5 m-1 text-gray-700" />
                </button>
              </Tippy>

              <Tippy content="New Link" placement="bottom">
                <button className="rounded-md flex justify-center items-center hover:bg-gray-200 h-7 w-7 cursor-default">
                  <PencilSquareIcon className="w-5 h-5 m-1 text-gray-700" />
                </button>
              </Tippy>
            </div>
          </footer>
        </div>
      </div>
    </ResizableContainer>
  );
};

export default SideBar;
