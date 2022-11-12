import { Fragment } from "react";
import SideNavMenuItem, { SideNavMenuItemData } from "./SideNavMenuItem";
import {
  HomeIcon,
  Cog6ToothIcon,
  NewspaperIcon,
  QuestionMarkCircleIcon,
  PencilSquareIcon,
  FolderPlusIcon,
  PlusIcon,
  XMarkIcon,
} from "@heroicons/react/24/outline";
import Tippy from "@tippyjs/react";
import { ReactNode } from "react";
import Logo from "../../assets/logo.svg";
import { Dialog, Transition } from "@headlessui/react";
import useSideNav from "../../stores/useSideNav";
import TopNav from "../TopNav";

const navItems: SideNavMenuItemData[] = [
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
          {
            href: "#",
            label: "Third level dsadsa dsa dsa dsa dsa dsa dsa dsa",
            active: false,
            children: [],
          },
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
    content: <HomeIcon className="w-6 h-6 text-gray-700" />,
    active: true,
  },
  {
    tooltip: "New Space",
    content: <PlusIcon className="w-6 h-6 text-purple-700" />,
    active: false,
  },
];

const bottomSideDideBarItems: SideSideBarItem[] = [
  {
    tooltip: "Help",
    content: <QuestionMarkCircleIcon className="w-6 h-6 text-gray-700" />,
    active: false,
  },
  {
    tooltip: "Settings",
    content: <Cog6ToothIcon className="w-6 h-6 text-gray-700" />,
    active: false,
  },
];

const SideNav = () => {
  const { isOpen, setOpen } = useSideNav();

  const createSideBarItem = (item: SideSideBarItem) => {
    const btn = (
      <button
        className={`${
          item.active ? "bg-gray-200" : ""
        } rounded-full flex justify-center items-center mx-2 hover:bg-gray-200 h-8 w-8 cursor-default`}
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

  const content = () => {
    return (
      <div className="flex flex-col w-full h-full">
        <TopNav containerClassNames="px-2">
          <a
            href="/"
            className="px-2 py-2 flex-0 inline-flex items-center"
            aria-label="Homepage"
          >
            <div className="flex text-2xl font-semibold">
              <img src={Logo} className="w-8 h-8 mr-3" alt="Logo" />{" "}
              <span className="lowercase text-red-700">rag</span>
              <span className="uppercase text-gray-700">bag</span>
            </div>
          </a>
        </TopNav>

        <div className="flex h-full w-full bg-white border-r border-r-gray-200">
          <div className="flex flex-col gap-1 justify-items-center border-r border-gray-200 overflow-x-hidden overflow-y-auto">
            <div />
            {topSideDideBarItems.map((item) => createSideBarItem(item))}
            <div className="flex-1" />
            {bottomSideDideBarItems.map((item) => createSideBarItem(item))}
            <div />
          </div>

          <div className="flex flex-col h-full w-full overflow-x-hidden overflow-y-auto">
            <nav className="flex flex-col mt-8 px-2 h-full">
              <div>
                <SideNavMenuItem
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
                  <SideNavMenuItem item={item} />
                ))}
              </div>
            </nav>

            <footer className="sticky inset-x-0 bottom-0 border-t py-1">
              <div className="flex gap-0.5 justify-center">
                <Tippy content="New Group" placement="bottom">
                  <button className="rounded-md flex justify-center items-center hover:bg-gray-200 h-8 w-8 cursor-default">
                    <FolderPlusIcon className="w-6 h-6 m-1 text-gray-700" />
                  </button>
                </Tippy>

                <Tippy content="New Link" placement="bottom">
                  <button className="rounded-md flex justify-center items-center hover:bg-gray-200 h-8 w-8 cursor-default">
                    <PencilSquareIcon className="w-6 h-6 m-1 text-gray-700" />
                  </button>
                </Tippy>
              </div>
            </footer>
          </div>
        </div>
      </div>
    );
  };

  return (
    <>
      {/* Menu for mobile */}
      <Transition.Root show={isOpen} as={Fragment}>
        <Dialog as="div" className="relative z-40 md:hidden" onClose={setOpen}>
          <Transition.Child
            as={Fragment}
            enter="transition-opacity ease-linear duration-300"
            enterFrom="opacity-0"
            enterTo="opacity-100"
            leave="transition-opacity ease-linear duration-300"
            leaveFrom="opacity-100"
            leaveTo="opacity-0"
          >
            <div className="fixed inset-0 bg-gray-600 bg-opacity-75" />
          </Transition.Child>

          <div className="fixed inset-0 z-40 flex">
            <Transition.Child
              as={Fragment}
              enter="transition ease-in-out duration-300 transform"
              enterFrom="-translate-x-full"
              enterTo="translate-x-0"
              leave="transition ease-in-out duration-300 transform"
              leaveFrom="translate-x-0"
              leaveTo="-translate-x-full"
            >
              <Dialog.Panel className="relative flex w-full max-w-xs flex-1 flex-col bg-white focus:outline-none">
                <Transition.Child
                  as={Fragment}
                  enter="ease-in-out duration-300"
                  enterFrom="opacity-0"
                  enterTo="opacity-100"
                  leave="ease-in-out duration-300"
                  leaveFrom="opacity-100"
                  leaveTo="opacity-0"
                >
                  <div className="absolute top-0 right-0 -mr-12 pt-4">
                    <button
                      type="button"
                      className="ml-1 flex h-10 w-10 items-center justify-center rounded-full focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white"
                      onClick={() => setOpen(false)}
                    >
                      <span className="sr-only">Close sidebar</span>
                      <XMarkIcon
                        className="h-6 w-6 text-white"
                        aria-hidden="true"
                      />
                    </button>
                  </div>
                </Transition.Child>
                {content()}
              </Dialog.Panel>
            </Transition.Child>
            <div className="w-14 flex-shrink-0" aria-hidden="true">
              {/* Force sidebar to shrink to fit close icon */}
            </div>
          </div>
        </Dialog>
      </Transition.Root>

      {/* Static and resizable menu for desktop */}
      <div className="bg-gray-50 min-w-min w-96 hidden md:flex md:lex-shrink-0 ">
        {content()}
      </div>
    </>
  );
};

export default SideNav;
