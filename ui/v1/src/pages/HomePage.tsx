import Layout from '@layouts/Layout';
import LinkCard from '@components/LinkCard';
import { Menu, Transition } from '@headlessui/react';
import { ChevronDownIcon, EllipsisVerticalIcon } from '@heroicons/react/24/outline';
import { Fragment } from 'react';
import { classNames } from '@utils/dom.utils';

const userNavigation = [
  { name: 'Your profile', href: '#' },
  { name: 'Sign out', href: '#' },
];

const HomePage = () => {
  return (
    <Layout
      title="Home"
      subTitle="Description"
      navbarRightElements={[
        <>
          {/* Profile dropdown */}
          <Menu as="div" className="relative">
            <Menu.Button className="-m-1.5 flex items-center p-1.5">
              <span className="sr-only">More</span>
              <span className="flex items-center">
                <EllipsisVerticalIcon className="ml-2 h-5 w-5 text-gray-700" aria-hidden="true" />
              </span>
            </Menu.Button>
            <Transition
              as={Fragment}
              enter="transition ease-out duration-100"
              enterFrom="transform opacity-0 scale-95"
              enterTo="transform opacity-100 scale-100"
              leave="transition ease-in duration-75"
              leaveFrom="transform opacity-100 scale-100"
              leaveTo="transform opacity-0 scale-95">
              <Menu.Items className="absolute right-0 z-10 mt-2.5 w-32 origin-top-right rounded-md bg-white py-2 shadow-lg ring-1 ring-gray-900/5 focus:outline-none">
                {userNavigation.map((item) => (
                  <Menu.Item key={item.name}>
                    {({ active }) => (
                      <a
                        href={item.href}
                        className={classNames(
                          active ? 'bg-gray-50' : '',
                          'block px-3 py-1 text-sm leading-6 text-gray-900'
                        )}>
                        {item.name}
                      </a>
                    )}
                  </Menu.Item>
                ))}
              </Menu.Items>
            </Transition>
          </Menu>
        </>,
      ]}>
      <h1 className="text-xl">Home</h1>
      <LinkCard
        item={{
          title: 'www.test.de',
          url: 'www.test.de',
        }}
        onClick={(item) => console.log(item)}
      />
    </Layout>
  );
};

export default HomePage;
