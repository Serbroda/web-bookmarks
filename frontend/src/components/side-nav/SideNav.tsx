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
import ResizableContainer from "../ResizableContainer";
import Tippy from "@tippyjs/react";
import { ReactNode } from "react";
import Logo from "../../assets/logo.svg";
import { Dialog, Transition } from "@headlessui/react";
import useSideNav from "../../stores/useSideNav";

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
      <div className="flex h-full w-full">
        <div className="flex flex-col gap-0.5 justify-items-center border-r border-gray-200 overflow-x-hidden overflow-y-auto">
          <div />
          {topSideDideBarItems.map((item) => createSideBarItem(item))}
          <div className="flex-1" />
          {bottomSideDideBarItems.map((item) => createSideBarItem(item))}
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
      <ResizableContainer
        width={312}
        conatinerClassName="shrink-0 bg-gray-50 min-w-[256px] max-w-[80%] hidden md:flex md:lex-shrink-0"
      >
        {content()}
      </ResizableContainer>
    </>
  );
};

export default SideNav;
