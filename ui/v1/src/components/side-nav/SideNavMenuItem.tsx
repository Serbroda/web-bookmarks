import { FC, ReactElement } from "react";
import { Disclosure } from "@headlessui/react";
import { ChevronUpIcon } from "@heroicons/react/20/solid";
import { Link } from "react-router-dom";

export interface SideNavMenuItemData {
  href: string;
  label: string;
  active: boolean;
  children: SideNavMenuItemData[];
  icon?: ReactElement;
}

const SideNavMenuItem: FC<{ item: SideNavMenuItemData }> = ({ item }) => {
  const hasActiveChild = (): boolean => {
    const hasActiveItems = (items: SideNavMenuItemData[]): boolean => {
      return items.some((item) => item.active || hasActiveItems(item.children));
    };
    return hasActiveItems(item.children);
  };

  const createLink = (item: SideNavMenuItemData) => {
    return (
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
            } w-5 h-5 shrink-0 group-hover:text-gray-600 mr-2`}
          >
            {item.icon}
          </span>
        ) : (
          <></>
        )}
        <span>{item.label}</span>
      </Link>
    );
  };

  return (
    <>
      {item.children.length == 0 ? createLink(item) : <></>}

      {item.children.length > 0 ? (
        <Disclosure defaultOpen={hasActiveChild()}>
          {({ open }) => (
            <>
              <div className="flex w-full items-center p-2 text-sm font-medium text-gray-600 mr-2 hover:bg-gray-100 rounded-md text-left">
                <Disclosure.Button className="rounded-md hover:bg-gray-200">
                  <ChevronUpIcon
                      className={`${
                          open
                              ? "rotate-180 text-gray-600"
                              : "rotate-90 text-gray-400"
                      } h-6 w-6`}
                  />
                </Disclosure.Button>

                <Link to={item.href} className="group flex flex-1 ml-1">
                  {item.icon ? (
                    <span className="w-5 h-5 text-gray-400 group-hover:text-gray-700 mr-2">
                      {item.icon}
                    </span>
                  ) : (
                    <></>
                  )}
                  <span className="group-hover:text-gray-800">
                    {item.label}
                  </span>
                </Link>
              </div>

              <Disclosure.Panel className="ml-4">
                {item.children.map((child, idx) => (
                  <SideNavMenuItem item={child} key={`${idx}_${item.label}`} />
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

export default SideNavMenuItem;
