import { FC, ReactElement } from "react";
import { Disclosure } from "@headlessui/react";
import { ChevronUpIcon } from "@heroicons/react/20/solid";
import { Link } from "react-router-dom";

export interface NavItemData {
  href: string;
  label: string;
  active: boolean;
  children: NavItemData[];
  icon?: ReactElement;
}

const NavItem: FC<{ item: NavItemData }> = ({ item }) => {
  const hasActiveChild = (): boolean => {
    const hasActiveItems = (items: NavItemData[]): boolean => {
      return items.some((item) => item.active || hasActiveItems(item.children));
    };
    return hasActiveItems(item.children);
  };

  return (
    <>
      {item.children.length == 0 ? (
        <Link
          to={item.href}
          className={`${
            item.active
              ? "text-gray-800 font-semibold"
              : "text-gray-600 font-medium"
          } group flex w-full items-center py-2 px-3 text-sm font-medium mr-2 hover:bg-gray-100 rounded-md`}
          key={item.label}
        >
          {item.icon ? (
            <span
              className={`${
                item.active
                  ? "text-gray-800 font-semibold"
                  : "text-gray-600 font-medium"
              } w-6 h-6 shrink-0 group-hover:text-gray-600 mr-2`}
            >
              {item.icon}
            </span>
          ) : (
            <></>
          )}
          <span>{item.label}</span>
        </Link>
      ) : (
        <></>
      )}

      {item.children.length > 0 ? (
        <Disclosure defaultOpen={hasActiveChild()}>
          {({ open }) => (
            <>
              <Disclosure.Button className="group flex w-full items-center py-2 px-3 text-sm font-medium text-gray-600 mr-2 hover:bg-gray-100 rounded-md text-left">
                {item.icon ? (
                  <span className="w-6 h-6 shrink-0 text-gray-400 group-hover:text-gray-600 mr-2">
                    {item.icon}
                  </span>
                ) : (
                  <></>
                )}
                <span className="flex-1">{item.label}</span>
                <ChevronUpIcon
                  className={`${
                    open
                      ? "-rotate-180 transform text-gray-600"
                      : "text-gray-400"
                  } h-6 w-6 text-purple-500`}
                />
              </Disclosure.Button>
              <Disclosure.Panel className="ml-4">
                {item.children.map((child) => (
                  <NavItem item={child} key={item.label} />
                ))}
              </Disclosure.Panel>
            </>
          )}
        </Disclosure>
      ) : (
        <></>
      )}
    </>
  );
};

export default NavItem;
