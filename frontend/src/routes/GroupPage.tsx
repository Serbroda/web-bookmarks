import LinkCard from "../components/LinkCard";
import GroupModal from "../components/modals/GroupModal";
import LinkModal from "../components/modals/LinkModal";
import TopNav from "../components/TopNav";
import useLinkModal from "../stores/modals/useLinkModal";
import useSideNav from "../stores/useSideNav";
import {EllipsisVerticalIcon} from "@heroicons/react/20/solid";
import {useEffect} from "react";
import {spacesApi} from "../services/config";
import useAlertModal from "../stores/modals/useAlertModal";
import {} from "../gen";

const GroupPage = () => {
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
        openModal({title: 'Lol', message: 'Create', confirmButtonMessage: 'Ok', onConfirm: () => console.log('Bla')})
    }, [])

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

            <div className="p-4 h-full overflow-auto">
                <div className="flex flex-wrap gap-4">
                    <LinkCard
                        onClick={() => {
                            openLinkModal({
                                mode: "edit",
                                onSave: () => {
                                    console.log("Link saved");
                                },
                            });
                        }}
                    />
                    <LinkCard
                        onClick={() => {
                            openLinkModal({
                                mode: "edit",
                                onSave: () => {
                                    console.log("Link saved");
                                },
                            });
                        }}
                    />

                    <LinkCard
                        onClick={() => {
                            openLinkModal({
                                mode: "edit",
                                onSave: () => {
                                    console.log("Link saved");
                                },
                            });
                        }}
                    />
                    <LinkCard
                        onClick={() => {
                            openLinkModal({
                                mode: "edit",
                                onSave: () => {
                                    console.log("Link saved");
                                },
                            });
                        }}
                    />
                    <LinkCard
                        onClick={() => {
                            openLinkModal({
                                mode: "edit",
                                onSave: () => {
                                    console.log("Link saved");
                                },
                            });
                        }}
                    />
                </div>
            </div>

            <GroupModal/>
            <LinkModal/>
        </div>
    );
};

export default GroupPage;
