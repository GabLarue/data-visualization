import React, { useEffect, useState } from 'react';
import './App.css';
import { Dialog } from '@headlessui/react';
import Api from './api/api';

type SavedFile = {
  name: string,
  key: string
}

function App() {
  const [savedFiles, setSavedFiles] = useState<SavedFile[]>()
  const [fileToUpload, setFileToUpload] = useState<File | null>(null)
  const [selectedFile, setSelectedFile] = useState<string>("")
  const [isFileOpen, setIsFileOpen] = useState<boolean>(false)
  const api = new Api

  useEffect(() => {
    (async() => {
      const files = await api.getAllFiles()
      setSavedFiles(files)
    })()
  })

  const handleUpload = async() => {
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

  const openFile = async(id: string) => {
    const file = await api.getFileById(id)
    setSelectedFile(file)
    setIsFileOpen(true)
  }

  return (
    <div className="flex w-screen h-screen">
      <Dialog open={isFileOpen} onClose={() => setIsFileOpen(false)} className="absolute w-full h-full inset-0">
        <Dialog.Panel className="w-full h-full p-10 overflow-y-auto">
          <div className="h-full relative flex flex-col p-10 overflow-y-auto bg-white rounded-lg shadow-lg">
            <div className="break-all">
              <span>{selectedFile}</span>
            </div>
            <button className="text-white w-[fit-content] flex items-center justify-center rounded px-4 py-2 bg-[#4E2ECFE6] hover:bg-[#4E2ECF] cursor-pointer sticky bottom-10 right-10" onClick={() => setIsFileOpen(false)}>Close</button>
          </div>
        </Dialog.Panel>
      </Dialog>
      <div className="flex flex-col justify-center items-center text-3xl gap-4 p-4 text-white w-1/2 bg-gradient-to-b from-[#1D1D42] to-[#4E2ECF]">
        <span className="text-white">
          <span className="text-[#FF7425] font-semibold">Visualize </span>
          your
          <span className="text-[#FF7425] font-semibold"> data </span>
          easily
        </span>
        <span className="material-symbols-outlined text-6xl text-[#FF7425]">insights</span>
      </div>
      <div className="flex flex-col gap-4 w-1/2 bg-white justify-center items-center">
        <div className="w-1/2 flex flex-col gap-2">
          {savedFiles ? savedFiles.map(file => {
            return <div onClick={() => openFile(file.key)} className="text-grey flex justify-between hover:text-white rounded px-4 py-2 bg-[#0182FF33] hover:bg-[#0182FFB3] cursor-pointer" key={file.key}>
              <div className="flex truncate gap-2">
                <span className="material-symbols-outlined">folder</span>
                <span className="truncate">{file.name}</span>
              </div>
            </div>
          }) :
            <div className="text-white rounded px-4 py-2 bg-[#00499033]">
              <span>No files were uploaded yet...</span>
            </div>
          }
        </div>
        <form id={"form"} className="flex flex-col gap-2 w-1/2" onSubmit={handleUpload}>
          <label className="text-white flex items-center justify-center rounded px-4 py-2 bg-[#4E2ECFE6] hover:bg-[#4E2ECF] cursor-pointer">
            <span>Select CSV to upload</span>
            <input type='file' accept='.csv' name='file' className="hidden" onChange={e => handleFileChange(e)}></input>
          </label>
          <button type="submit" disabled={fileToUpload ? false : true} className={`${!fileToUpload && "opacity-20"} truncate text-white flex items-center justify-center rounded px-4 py-2 bg-[#0182FFE6] hover:bg-[#0182FF]`}>
            {fileToUpload ? `Upload ${fileToUpload.name}` : "Upload CSV"}
          </button>
        </form>
      </div>
    </div>
  );
}

export default App;
