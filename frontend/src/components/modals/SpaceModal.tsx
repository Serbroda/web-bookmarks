import {Dialog, Tab} from "@headlessui/react";
import Modal from "./Modal";
import useSpaceModal from "../../stores/modals/useSpaceModal";
import {classNames} from "../../utils/dom.utils";
import {FC} from "react";

const TabHeader: FC<{ title: string }> = ({title}) => {
    return (
        <Tab
            className={({selected}) =>
                classNames(
                    'w-full py-2.5 text-sm font-medium leading-5 text-blue-700 rounded-t-md',
                    selected
                        ? 'border-x border-t'
                        : 'border-b text-blue-200 hover:bg-white/[0.12] hover:text-blue-600'
                )
            }
        >{title}
        </Tab>
    )
}

const SpaceModal = () => {
    const {isOpen, props, setOpen} = useSpaceModal();

    return (
        <Modal show={isOpen} onClose={() => setOpen(false)} width="medium">
            <Tab.Group>
                <Tab.List className="flex w-[98%]">
                    <TabHeader title="Create"/>
                    <TabHeader title="Join"/>
                </Tab.List>
                <Tab.Panels>
                    <Tab.Panel>
                        <div className="sm:flex sm:items-start w-full pt-4">
                            <div className="mt-3 text-center sm:mt-0 sm:text-left w-full">
                                <Dialog.Title
                                    as="h3"
                                    className="text-lg font-medium leading-6 text-gray-900"
                                >
                                    Create a new Space
                                </Dialog.Title>

                                <div className="mt-2 w-full pt-4">
                                    <input
                                        type="text"
                                        id="spaceName"
                                        name="spaceName"
                                        className="input w-full"
                                        placeholder="Name"
                                    />
                                    <textarea
                                        rows={4}
                                        name="spaceDescription"
                                        id="spaceDescription"
                                        className="block w-full input mt-4"
                                        placeholder="Description (optional)"
                                        defaultValue={""}
                                    />
                                </div>
                            </div>
                        </div>

                        <div className="mt-5 sm:mt-4 sm:flex sm:flex-row-reverse">
                            <button
                                type="button"
                                className="btn btn-primary w-full sm:ml-3 sm:w-auto"
                                onClick={() => {
                                    if (props.onSave) {
                                        props.onSave();
                                    }
                                    setOpen(false);
                                }}
                                tabIndex={2}
                            >
                                Save
                            </button>
                            <button
                                type="button"
                                className="btn mt-4 w-full sm:mt-0 sm:w-auto sm:text-sm"
                                onClick={() => {
                                    setOpen(false);
                                }}
                                autoFocus
                                tabIndex={1}
                            >
                                Cancel
                            </button>
                        </div>
                    </Tab.Panel>

                    <Tab.Panel>
                        <div className="sm:flex sm:items-start w-full pt-4">
                            <div className="mt-3 text-center sm:mt-0 sm:text-left w-full">
                                <Dialog.Title
                                    as="h3"
                                    className="text-lg font-medium leading-6 text-gray-900"
                                >
                                    Join existing Space
                                </Dialog.Title>

                                <div className="mt-2 w-full pt-4">
                                    <div>
                                        <label htmlFor="spaceInviteLink"
                                               className="block text-sm font-medium text-gray-700">
                                            Invitation Link
                                        </label>
                                        <div className="mt-1">
                                            <input
                                                id="spaceInviteLink"
                                                name="spaceInviteLink"
                                                type="text"
                                                autoComplete="space-inv"
                                                required
                                                className="input w-full"
                                                placeholder="https://app.ragbag.dev/invite/AbC732hjc"
                                            />
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>

                        <div className="mt-5 sm:mt-4 sm:flex sm:flex-row-reverse">
                            <button
                                type="button"
                                className="btn btn-primary w-full sm:ml-3 sm:w-auto"
                                onClick={() => {
                                    if (props.onSave) {
                                        props.onSave();
                                    }
                                    setOpen(false);
                                }}
                                tabIndex={2}
                            >
                                Join Space
                            </button>
                            <button
                                type="button"
                                className="btn mt-4 w-full sm:mt-0 sm:w-auto sm:text-sm"
                                onClick={() => {
                                    setOpen(false);
                                }}
                                autoFocus
                                tabIndex={1}
                            >
                                Cancel
                            </button>
                        </div>
                    </Tab.Panel>
                </Tab.Panels>
            </Tab.Group>
        </Modal>
    );
};

export default SpaceModal;
