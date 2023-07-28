import { Fragment, useEffect, useState } from 'react';
import SideNavMenuItem, { SideNavMenuItemData } from './SideNavMenuItem';
import {
  ArrowLeftOnRectangleIcon,
  Cog6ToothIcon,
  FolderPlusIcon,
  HomeIcon,
  NewspaperIcon,
  PencilSquareIcon,
  PlusIcon,
  QuestionMarkCircleIcon,
  WrenchScrewdriverIcon,
  XMarkIcon,
} from '@heroicons/react/24/outline';
import Tippy from '@tippyjs/react';
import Logo from '../../assets/react.svg';
import { Dialog, Transition } from '@headlessui/react';
import useSideNav from '../../stores/useSideNav';
import TopNav from '../TopNav';
import SideNavLeftItem, { SideSideBarItem } from './SideNavLeftItem';

const navItems: SideNavMenuItemData[] = [
  {
    href: 'groups/0',
    label: 'Groups',
    active: false,
    children: [
      { href: 'groups/1', label: 'All posts', active: false, children: [] },
      { href: 'groups/2', label: 'Add new', active: false, children: [] },
      { href: 'groups/3', label: 'Categories', active: false, children: [] },
    ],
    icon: <NewspaperIcon />,
  },
  {
    href: '#',
    label: 'Media',
    active: false,
    children: [
      { href: '#', label: 'Library', active: false, children: [] },
      {
        href: '#',
        label: 'Add new',
        active: false,
        children: [
          {
            href: '#',
            label: 'Third level dsadsa dsa dsa dsa dsa dsa dsa dsa',
            active: false,
            children: [],
          },
        ],
      },
    ],
  },
];

const bottomSideDideBarItems: SideSideBarItem[] = [
  {
    tooltip: 'Help',
    content: <QuestionMarkCircleIcon className="w-6 h-6 text-gray-700" />,
    active: false,
  },
  {
    tooltip: 'Settings',
    content: <Cog6ToothIcon className="w-6 h-6 text-gray-700" />,
    active: false,
  },
  {
    tooltip: 'Logout',
    content: <ArrowLeftOnRectangleIcon className="w-6 h-6 text-gray-700" />,
    active: false,
    onClick: () => console.log('logout'),
  },
];

const SideNav = () => {
  const { isOpen, setOpen } = useSideNav();

  return (
    <>
      <Transition.Root show={isOpen} as={Fragment}>
        <Dialog as="div" className="relative z-40 md:hidden" onClose={setOpen}>
          <Transition.Child
            as={Fragment}
            enter="transition-opacity ease-linear duration-300"
            enterFrom="opacity-0"
            enterTo="opacity-100"
            leave="transition-opacity ease-linear duration-300"
            leaveFrom="opacity-100"
            leaveTo="opacity-0">
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
              leaveTo="-translate-x-full">
              <Dialog.Panel className="relative flex w-full max-w-xs flex-1 flex-col bg-white focus:outline-none">
                <Transition.Child
                  as={Fragment}
                  enter="ease-in-out duration-300"
                  enterFrom="opacity-0"
                  enterTo="opacity-100"
                  leave="ease-in-out duration-300"
                  leaveFrom="opacity-100"
                  leaveTo="opacity-0">
                  <div className="absolute top-0 right-0 -mr-12 pt-4">
                    <button
                      type="button"
                      className="ml-1 flex h-10 w-10 items-center justify-center rounded-full focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white"
                      onClick={() => setOpen(false)}>
                      <span className="sr-only">Close sidebar</span>
                      <XMarkIcon className="h-6 w-6 text-white" aria-hidden="true" />
                    </button>
                  </div>
                </Transition.Child>

                <div className="h-full max-h-screen">
                  <TopNav showMenuButton={false}>Hi</TopNav>

                  <nav className="flex flex-col h-[calc(100vh-3rem-3rem)] overflow-y-auto border-r">
                    <div className="p-3">
                      <SideNavMenuItem
                        item={{
                          href: '/',
                          label: 'Home',
                          active: true,
                          children: [],
                          icon: <HomeIcon />,
                        }}
                      />

                      <h2 className="w-full py-4 text-sm font-semibold text-gray-400">Groups</h2>

                      {navItems.map((item, idx) => (
                        <SideNavMenuItem key={idx} item={item} />
                      ))}
                    </div>
                  </nav>
                  <footer className="h-12 border-t border-r">test</footer>
                </div>
              </Dialog.Panel>
            </Transition.Child>
            <div className="w-14 flex-shrink-0" aria-hidden="true">
              {/* Force sidebar to shrink to fit close icon */}
            </div>
          </div>
        </Dialog>
      </Transition.Root>

      <div className="h-full max-h-screen hidden md:block w-60">
        <TopNav showMenuButton={false}>Hi</TopNav>

        <nav className="flex flex-col h-[calc(100vh-3rem-3rem)] overflow-y-auto border-r">
          <div className="p-3">
            <SideNavMenuItem
              item={{
                href: '/',
                label: 'Home',
                active: true,
                children: [],
                icon: <HomeIcon />,
              }}
            />

            <h2 className="w-full py-4 text-sm font-semibold text-gray-400">Groups</h2>

            {navItems.map((item, idx) => (
              <SideNavMenuItem key={idx} item={item} />
            ))}
          </div>
        </nav>
        <footer className="h-12 border-t border-r">test</footer>
      </div>
    </>
  );
};

export default SideNav;
