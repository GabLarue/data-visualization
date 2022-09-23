type FileInputProps = {
    handleFileChange: (e: React.ChangeEvent<HTMLInputElement>) => void
}

const FileInput = ({ handleFileChange }: FileInputProps) => {
    return (
        <label className="text-white flex items-center justify-center rounded px-4 py-2 bg-[#4E2ECFE6] hover:bg-[#4E2ECF] cursor-pointer">
            <span>Select CSV to upload</span>
            <input
                type='file'
                accept='.csv'
                name='file'
                className="hidden"
                onChange={e => handleFileChange(e)} />
        </label>
    )
}

export default FileInput