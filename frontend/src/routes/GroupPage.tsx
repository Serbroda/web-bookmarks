import LinkCard, {LinkItem} from "../components/LinkCard";
import GroupModal from "../components/modals/GroupModal";
import LinkModal from "../components/modals/LinkModal";
import TopNav from "../components/TopNav";
import useLinkModal from "../stores/modals/useLinkModal";
import useSideNav from "../stores/useSideNav";
import {EllipsisVerticalIcon} from "@heroicons/react/20/solid";
import {useEffect, useState} from "react";
import useAlertModal from "../stores/modals/useAlertModal";
import {formatUrl, isValidHttpUrl} from "../utils/url.utils";
import {ArrowRightIcon} from '@heroicons/react/24/outline'

const mockItems: LinkItem[] = [{
    title: "www.google.de",
    url: "www.google.de"
}, {
    title: "heise",
    url: "https://heise.de"
}, {
    title: "golem.de",
    url: "golem.de"
}, {
    title: "Entwickler Info",
    url: "dev.to"
}];

const GroupPage = () => {
    const [search, setSearch] = useState<string>("")
    const [hasValidUrl, setHasValidUrl] = useState(false)
    const {isOpen, setOpen} = useSideNav();
    const {openModal: openLinkModal} = useLinkModal();

    const {openModal} = useAlertModal();

    useEffect(() => {
        /*const cre = async () => {
            const res = await spacesApi.v1SpacesGet({
                createSpaceDto: {name: "Test", description: "A test space", visibility: SpaceVisibility.Private}
            })
            console.log(res)
        }*/
        //cre()
        //openModal({title: 'Lol', message: 'Create', confirmButtonMessage: 'Ok', onConfirm: () => console.log('Bla')})
    }, [])

    useEffect(() => {
        setHasValidUrl(false)
        let formattedUrl: URL | undefined;

        if (search !== '') {
            formattedUrl = formatUrl(search);
        }

        if (isValidHttpUrl(formattedUrl?.href)) {
            setHasValidUrl(true)
        }
    }, [search])

    return (
        <div className="flex flex-col relative h-full">
            <TopNav>
                <div className="flex flex-1 items-stretch justify-start">
                    <div className="block">
                        <div className="flex space-x-4">Title</div>
                    </div>
                </div>
                <div
                    className="absolute inset-y-0 right-0 flex items-center pr-2 sm:static sm:inset-auto sm:ml-6 sm:pr-0">
                    <button type="button" className="btn btn-ghost rounded-full ">
                        <EllipsisVerticalIcon className="h-5 w-5"/>
                    </button>
                </div>
            </TopNav>

            <div className="p-4 h-full overflow-auto flex flex-col items-center">
                <div className="w-full lg:w-[80%] pb-4 relative">
                    <input
                        id="username"
                        name="username"
                        type="search"
                        autoComplete="username"
                        required
                        className={`input w-full ${hasValidUrl ? 'input-success' : ''}`}
                        placeholder="Search or add link"
                        onChange={event => setSearch(event.target.value)}
                    />
                    {search && hasValidUrl &&
                        <button type="submit"
                                className="btn btn-xs btn-ghost text-gray-400 hover:text-indigo-400 hover:bg-gray-50 absolute right-8 top-1 w-24">
                            <span className="pr-2">Add link</span> <ArrowRightIcon className="h-4 w-4"/>
                        </button>
                    }
                </div>

                <div className="w-full grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                    {mockItems.filter((item) => {
                        if (!search || !search.trim()) {
                            return true;
                        }
                        const text = search.toLowerCase().trim();
                        return item.title.toLowerCase().includes(text) || item.url.toLowerCase().includes(text)
                    }).map((item, idx) => <LinkCard key={idx}
                                                    item={item}
                                                    onClick={() => {
                                                        openLinkModal({
                                                            mode: "edit",
                                                            onSave: () => {
                                                                console.log("Link saved");
                                                            },
                                                        });
                                                    }}
                    />)}
                </div>
            </div>

            <GroupModal/>
            <LinkModal/>
        </div>
    );
};

export default GroupPage;
