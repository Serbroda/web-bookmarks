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
  const [spaces, setSpaces] = useState<SideSideBarItem[]>([]);

  useEffect(() => {
    loadSpaces();
  }, []);

  const loadSpaces = async () => {
    const result: any = [{}];
    let items: SideSideBarItem[] = result.map((i: any) => {
      return {
        content: <HomeIcon className="w-6 h-6 text-gray-700" />,
        active: false,
        tooltip: i.name,
      } as SideSideBarItem;
    });
    setSpaces(items);
  };

  return (
    <div className="h-full max-h-screen w-60">
      <TopNav showMenuButton={false}>Hi</TopNav>

      <nav className="flex flex-col h-[calc(100vh-3rem-3rem)] overflow-y-auto border-r">
        <div>
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
  );
};

export default SideNav;
