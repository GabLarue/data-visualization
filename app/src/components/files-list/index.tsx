import { SavedFile } from "../../api/types"
import NoSavedFileMessage from "./no-saved-file-message"

type FilesListProps = {
    openFile: (id: string) => Promise<void>
    files: SavedFile[]
}

const FilesList = ({ openFile, files }: FilesListProps) => {
    if (!files) {
        return (
            <NoSavedFileMessage />
        )
    }

    return (
        <>
            {files.map(file => {
                return (
                    <div
                        key={file.key}
                        onClick={() => openFile(file.key)}
                        className="text-grey flex justify-between hover:text-white rounded px-4 py-2 bg-[#0182FF33] hover:bg-[#0182FFB3] cursor-pointer">
                        <div className="flex truncate gap-2">
                            <span className="material-symbols-outlined">folder</span>
                            <span className="truncate">{file.name}</span>
                        </div>
                    </div>
                )
            })}
        </>)
}

export default FilesList;