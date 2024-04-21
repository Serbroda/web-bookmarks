import useSideNav from '../stores/useSideNav';
import { Bars3Icon } from '@heroicons/react/20/solid';
import { FC, ReactNode } from 'react';

export interface TopNavProps {
  children: ReactNode;
  showMenuButton?: boolean;
  navClassNames?: string;
  containerClassNames?: string;
}

const TopNav: FC<TopNavProps> = ({
  children,
  showMenuButton = true,
  navClassNames = 'bg-white border-b border-gray-300',
  containerClassNames = 'px-2 sm:px-6 lg:px-8',
}) => {
  const { toggle } = useSideNav();

  return (
    <nav className={`sticky top-0 h-12 z-10 ${navClassNames}`}>
      <div className={`mx-auto max-w-7xl ${containerClassNames}`}>
        <div className="relative flex h-12 items-center justify-between">
          {showMenuButton && (
            <div className="inset-y-0 left-0 flex items-center md:hidden">
              <button type="button" className="inline-flex btn btn-ghost" onClick={toggle}>
                <span className="sr-only">Open main menu</span>
                <Bars3Icon className="w-5 h-5" />
              </button>
            </div>
          )}

          {children}
        </div>
      </div>
    </nav>
  );
};

export default TopNav;
