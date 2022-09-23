import { useState } from "react"

import Api from "../../api/api"
import FileInput from "./file-input"

const FileUploadForm = () => {
    const [fileToUpload, setFileToUpload] = useState<File | null>(null)
    const api = new Api

    const handleUpload = async () => {
        if (fileToUpload === null) {
            alert("No file selected for upload!")
        } else {
            var formData = new FormData()
            formData.append('file', fileToUpload)

            await api.uploadFile(formData)
            setFileToUpload(null)
        }
    }

    const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        if (e.target && e.target.files) {
            setFileToUpload(e.target.files[0])
        }
    }

    return (
        <form id={"form"} className="flex flex-col gap-2 w-1/2" onSubmit={handleUpload}>
            <FileInput handleFileChange={handleFileChange} />
            <button type="submit"
                disabled={fileToUpload ? false : true}
                className={`${!fileToUpload && "opacity-20"} truncate text-white flex items-center justify-center rounded px-4 py-2 bg-[#0182FFE6] hover:bg-[#0182FF]`}>
                {fileToUpload ? `Upload ${fileToUpload.name}` : "Upload CSV"}
            </button>
        </form>
    )
}

export default FileUploadForm;