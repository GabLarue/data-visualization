import { Dialog } from "@headlessui/react"

type FileViewerModalProps = {
    isFileOpen: boolean
    setIsFileOpen: React.Dispatch<React.SetStateAction<boolean>>
    file: string
}

const FileViewerModal = ({ isFileOpen, setIsFileOpen, file }: FileViewerModalProps) => {
    return (
        <Dialog
            open={isFileOpen}
            onClose={() => setIsFileOpen(false)}
            className="absolute w-full h-full inset-0">
            <Dialog.Panel className="w-full h-full p-10 overflow-y-auto">
                <div className="h-full relative flex flex-col p-10 overflow-y-auto bg-white rounded-lg shadow-lg">
                    <div className="break-all">
                        <span>{file}</span>
                    </div>
                    <button className="text-white w-[fit-content] flex items-center justify-center rounded px-4 py-2 bg-[#4E2ECFE6] hover:bg-[#4E2ECF] cursor-pointer sticky bottom-10 right-10" onClick={() => setIsFileOpen(false)}>Close</button>
                </div>
            </Dialog.Panel>
        </Dialog>
    )
}

export default FileViewerModal