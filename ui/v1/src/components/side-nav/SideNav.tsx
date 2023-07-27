import { Cog6ToothIcon, HomeIcon, NewspaperIcon } from '@heroicons/react/24/outline';
import SideNavMenuItem, { SideNavMenuItemData } from '@components/side-nav/SideNavMenuItem';
import { Link } from 'react-router-dom';

const navigation = [{ name: 'Home', href: '#', icon: HomeIcon, current: true }];
const userNavigation = [
  { name: 'Your profile', href: '#' },
  { name: 'Sign out', href: '#' },
];

function classNames(...classes: string[]) {
  return classes.filter(Boolean).join(' ');
}

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

const SideNav = () => {
  return (
    <>
      <div className="flex h-16 shrink-0 items-center">
        <img
          className="h-8 w-auto"
          src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600"
          alt="Your Company"
        />
      </div>
      <nav className="flex flex-1 flex-col">
        <ul role="list" className="flex flex-1 flex-col gap-y-7">
          <li>
            <ul role="list" className="-mx-2 space-y-1">
              {navigation.map((item) => (
                <li key={item.name}>
                  <a
                    href={item.href}
                    className={classNames(
                      item.current
                        ? 'bg-gray-50 text-indigo-600'
                        : 'text-gray-700 hover:text-indigo-600 hover:bg-gray-50',
                      'group flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold'
                    )}>
                    <item.icon
                      className={classNames(
                        item.current ? 'text-indigo-600' : 'text-gray-400 group-hover:text-indigo-600',
                        'h-6 w-6 shrink-0'
                      )}
                      aria-hidden="true"
                    />
                    {item.name}
                  </a>
                </li>
              ))}
            </ul>
          </li>
          <li>
            <div className="text-xs font-semibold leading-6 text-gray-400">Your Links</div>
            <ul role="list" className="-mx-2 mt-2 space-y-1">
              {navItems.map((item) => (
                <SideNavMenuItem item={item} />
              ))}
            </ul>
          </li>
          <li className="mt-auto">
            <Link
              to="/settings"
              className="group -mx-2 flex gap-x-3 rounded-md p-2 text-sm font-semibold leading-6 text-gray-700 hover:bg-gray-50 hover:text-indigo-600">
              <Cog6ToothIcon
                className="h-6 w-6 shrink-0 text-gray-400 group-hover:text-indigo-600"
                aria-hidden="true"
              />
              Settings
            </Link>
          </li>
        </ul>
      </nav>
    </>
  );
};

export default SideNav;
