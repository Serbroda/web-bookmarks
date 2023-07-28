import { ClipboardIcon, PencilIcon } from '@heroicons/react/24/outline';
import { FC } from 'react';

export interface LinkItem {
  title: string;
  url: string;
}

export interface LinkCardProps {
  item: LinkItem;
}

const LinkCard: FC<LinkCardProps> = ({ item }) => {
  return (
    <div
      className="card relative bg-white p-4 invisble-hover-container hover:border-indigo-500 hover:cursor-pointer"
      onClick={() => {
        window.open(item.url, '_blank');
      }}>
      <div className="leading-none">
        <span className="font-semibold line-clamp-2">{item.title}</span>{' '}
        <div className="text-sm flex leading-none mt-2">
          <img
            className="favicon favicon-sm pr-1"
            src="https://icons.duckduckgo.com/ip3/www.youtube.com.ico"
            alt="favicon"
          />{' '}
          <div>{item.url}</div>
        </div>{' '}
        <i className="text-xs text-gray-500">18.10.02022, 10:48</i>
      </div>

      <div className="invisble-hover-item flex absolute bottom-1 right-2">
        <button
          className="hover:text-indigo-500"
          onClick={(e) => {
            e.stopPropagation();
            e.nativeEvent.stopImmediatePropagation();

            navigator.clipboard.writeText(item.url);
          }}>
          <ClipboardIcon className="w-6 h-6 md:w-5 md:h-5" />
        </button>{' '}
        <button
          className="hover:text-indigo-500"
          onClick={(e) => {
            e.stopPropagation();
            e.nativeEvent.stopImmediatePropagation();

            console.log('edit click');
          }}>
          <PencilIcon className="w-6 h-6 md:w-5 md:h-5" />
        </button>
      </div>
    </div>
  );
};

export default LinkCard;
