import {Fragment, useEffect, useState} from "react";
import SideNavMenuItem, {SideNavMenuItemData} from "./SideNavMenuItem";
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
    XMarkIcon
} from "@heroicons/react/24/outline";
import Tippy from "@tippyjs/react";
import Logo from "../../assets/logo.svg";
import {Dialog, Transition} from "@headlessui/react";
import useSideNav from "../../stores/useSideNav";
import TopNav from "../TopNav";
import {spacesApi, authService} from "../../services/config";
import useAuthentication from "../../stores/useAuthentication";
import ScrollArea from "../ScrollArea";
import SideNavLeftItem, {SideSideBarItem} from "./SideNavLeftItem";
import useSpaceModal from "../../stores/modals/useSpaceModal";

const navItems: SideNavMenuItemData[] = [
    {
        href: "groups/0",
        label: "Groups",
        active: false,
        children: [
            {href: "groups/1", label: "All posts", active: false, children: []},
            {href: "groups/2", label: "Add new", active: false, children: []},
            {href: "groups/3", label: "Categories", active: false, children: []},
        ],
        icon: <NewspaperIcon/>,
    },
    {
        href: "#",
        label: "Media",
        active: false,
        children: [
            {href: "#", label: "Library", active: false, children: []},
            {
                href: "#",
                label: "Add new",
                active: false,
                children: [
                    {
                        href: "#",
                        label: "Third level dsadsa dsa dsa dsa dsa dsa dsa dsa",
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
        tooltip: "Help",
        content: <QuestionMarkCircleIcon className="w-6 h-6 text-gray-700"/>,
        active: false,
    },
    {
        tooltip: "Settings",
        content: <Cog6ToothIcon className="w-6 h-6 text-gray-700"/>,
        active: false,
    },
    {
        tooltip: "Logout",
        content: <ArrowLeftOnRectangleIcon className="w-6 h-6 text-gray-700"/>,
        active: false,
        onClick: () => authService.logout()
    },
];

const SideNav = () => {
    const {isOpen, setOpen} = useSideNav();
    const {openModal: openSpacesModal} = useSpaceModal();
    const [spaces, setSpaces] = useState<SideSideBarItem[]>([]);

    useEffect(() => {
        loadSpaces();
    }, []);

    const loadSpaces = async () => {
        const result = await spacesApi.getSpaces();
        let items: SideSideBarItem[] = result.map((i) => {
            return {
                content: <HomeIcon className="w-6 h-6 text-gray-700"/>,
                active: false,
                tooltip: i.name
            } as SideSideBarItem
        });
        setSpaces(items);
    }

    const content = () => {
        return (
            <div className="flex flex-col w-full h-full">
                <TopNav containerClassNames="px-2" showMenuButton={false}>
                    <a
                        href="/"
                        className="px-2 py-2 flex-0 inline-flex items-center"
                        aria-label="Homepage"
                    >
                        <div className="flex text-2xl font-semibold">
                            <img src={Logo} className="w-8 h-8 mr-3" alt="Logo"/>{" "}
                            <span className="lowercase text-red-700">rag</span>
                            <span className="uppercase text-gray-700">bag</span>
                        </div>
                    </a>
                </TopNav>

                <div className="flex h-full--3rem w-full bg-white border-r border-r-gray-200">
                    <ScrollArea overflowY="auto"
                                className="border-r border-gray-200 flex flex-col gap-1 justify-items-center">
                        <div/>
                        {spaces.map((item, idx) =>
                            <SideNavLeftItem
                                key={idx}
                                item={item}
                            />)}
                        <SideNavLeftItem
                            item={{
                                tooltip: "New Space",
                                content: <PlusIcon className="w-6 h-6 text-indigo-700"/>,
                                active: false,
                                onClick: () => {
                                    console.log('Clicked')
                                    openSpacesModal({
                                        onSave: () => console.log('saved')
                                    })
                                }
                            }}
                        />
                        <div className="flex-1"/>
                        {bottomSideDideBarItems.map((item, idx) =>
                            <SideNavLeftItem
                                key={idx}
                                item={item}
                            />)}
                        <div/>
                    </ScrollArea>

                    <ScrollArea overflowY="auto" className="flex flex-col w-full">
                        <nav className="flex flex-col mt-8 px-2 h-full">
                            <div>
                                <SideNavMenuItem
                                    item={{
                                        href: "/",
                                        label: "Home",
                                        active: true,
                                        children: [],
                                        icon: <HomeIcon/>,
                                    }}
                                />

                                <h2 className="w-full py-4 px-1 text-sm font-semibold text-gray-400">
                                    My Groups
                                </h2>

                                {navItems.map((item, idx) => (
                                    <SideNavMenuItem key={idx} item={item}/>
                                ))}
                            </div>
                        </nav>

                        <footer className="sticky inset-x-0 bottom-0 border-t py-1 px-2 bg-white">
                            <div className="flex space-x-0.5 justify-center">
                                <div className="flex-1"/>
                                <Tippy content="New Group" placement="bottom">
                                    <button
                                        className="rounded-md flex justify-center items-center hover:bg-gray-200 h-8 w-8">
                                        <FolderPlusIcon className="w-6 h-6 m-1 text-gray-700"/>
                                    </button>
                                </Tippy>

                                <Tippy content="New Link" placement="bottom">
                                    <button
                                        className="rounded-md flex justify-center items-center hover:bg-gray-200 h-8 w-8">
                                        <PencilSquareIcon className="w-6 h-6 m-1 text-gray-700"/>
                                    </button>
                                </Tippy>

                                <div className="flex-1"/>
                                <Tippy content="Edit Space" placement="bottom">
                                    <button
                                        className="rounded-md flex justify-center items-center hover:bg-gray-200 h-8 w-8">
                                        <WrenchScrewdriverIcon className="w-6 h-6 m-1 text-gray-700"/>
                                    </button>
                                </Tippy>
                            </div>
                        </footer>
                    </ScrollArea>
                </div>
            </div>
        );
    };

    return (
        <>
            {/* Menu for mobile */}
            <Transition.Root show={isOpen} as={Fragment}>
                <Dialog as="div" className="relative z-40 md:hidden" onClose={setOpen}>
                    <Transition.Child
                        as={Fragment}
                        enter="transition-opacity ease-linear duration-300"
                        enterFrom="opacity-0"
                        enterTo="opacity-100"
                        leave="transition-opacity ease-linear duration-300"
                        leaveFrom="opacity-100"
                        leaveTo="opacity-0"
                    >
                        <div className="fixed inset-0 bg-gray-600 bg-opacity-75"/>
                    </Transition.Child>

                    <div className="fixed inset-0 z-40 flex">
                        <Transition.Child
                            as={Fragment}
                            enter="transition ease-in-out duration-300 transform"
                            enterFrom="-translate-x-full"
                            enterTo="translate-x-0"
                            leave="transition ease-in-out duration-300 transform"
                            leaveFrom="translate-x-0"
                            leaveTo="-translate-x-full"
                        >
                            <Dialog.Panel
                                className="relative flex w-full max-w-xs flex-1 flex-col bg-white focus:outline-none">
                                <Transition.Child
                                    as={Fragment}
                                    enter="ease-in-out duration-300"
                                    enterFrom="opacity-0"
                                    enterTo="opacity-100"
                                    leave="ease-in-out duration-300"
                                    leaveFrom="opacity-100"
                                    leaveTo="opacity-0"
                                >
                                    <div className="absolute top-0 right-0 -mr-12 pt-4">
                                        <button
                                            type="button"
                                            className="ml-1 flex h-10 w-10 items-center justify-center rounded-full focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white"
                                            onClick={() => setOpen(false)}
                                        >
                                            <span className="sr-only">Close sidebar</span>
                                            <XMarkIcon
                                                className="h-6 w-6 text-white"
                                                aria-hidden="true"
                                            />
                                        </button>
                                    </div>
                                </Transition.Child>
                                {content()}
                            </Dialog.Panel>
                        </Transition.Child>
                        <div className="w-14 flex-shrink-0" aria-hidden="true">
                            {/* Force sidebar to shrink to fit close icon */}
                        </div>
                    </div>
                </Dialog>
            </Transition.Root>

            {/* Static and resizable menu for desktop */}
            <div className="bg-gray-50 min-w-min w-96 hidden md:flex md:lex-shrink-0">
                {content()}
            </div>
        </>
    );
};

export default SideNav;
