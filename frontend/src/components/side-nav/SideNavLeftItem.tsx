import {FC, ReactNode} from "react";
import Tippy from "@tippyjs/react";

export interface SideSideBarItem {
    content: ReactNode;
    active: boolean;
    tooltip?: string;
    href?: string;
    onClick?: () => void;
}

const SideNavLeftItem: FC<{ item: SideSideBarItem }> = ({item}) => {
    return (
        <Tippy content={item.tooltip} placement="right">
            <div className="flex items-center">
                <div
                    className={`w-1 h-7 mr-1 rounded-r-md ${
                        item.active ? "bg-indigo-700" : ""
                    }`}
                />
                <button
                    className={`${
                        item.active ? "" : ""
                    } rounded-full h-9 w-11 mr-2 flex justify-center items-center hover:bg-gray-300`}
                    onClick={item.onClick}
                >
                    {item.content}
                </button>
            </div>
        </Tippy>
    )
}

export default SideNavLeftItem;
