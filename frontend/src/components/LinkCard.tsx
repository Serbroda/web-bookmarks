import { Dialog, Transition } from "@headlessui/react";
import { FC, Fragment, ReactNode } from "react";
import { XMarkIcon } from "@heroicons/react/20/solid";

export interface LinkItem {
  title: string;
}

export interface LinkCardProps {
  onClick: () => void;
}

const LinkCard: FC<LinkCardProps> = ({ onClick }) => {
  return (
    <div
      className="card relative w-80 bg-white p-4 invisble-hover-container hover:border-indigo-500 hover:cursor-pointer"
      onClick={onClick}
    >
      <div className="leading-none">
        <span className="font-semibold line-clamp-2">
          WebSocket Recipe | Echo - High performance, minimalist Go web
          framework
        </span>{" "}
        <div className="text-sm flex leading-none mt-2">
          <img
            className="favicon favicon-sm pr-1"
            src="https://icons.duckduckgo.com/ip3/www.youtube.com.ico"
            alt="favicon"
          />{" "}
          <div>www.youtube.com</div>
        </div>{" "}
        <i className="text-xs text-gray-500">18.10.02022, 10:48</i>
      </div>

      <div className="invisble-hover-item absolute bottom-1 right-2">
        <button className="hover:text-primary">
          <svg
            className="w-5 h-5"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"
            ></path>
          </svg>
        </button>{" "}
        <button className="hover:text-primary">
          <svg
            className="w-5 h-5"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"
            ></path>
          </svg>
        </button>
      </div>
    </div>
  );
};

export default LinkCard;