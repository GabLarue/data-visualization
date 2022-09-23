import React, { useEffect, useState } from 'react';
import './App.css';
import { Dialog } from '@headlessui/react';
import Api from './api/api';
import FileUploadForm from './components/file-upload-form.tsx';
import FileViewerModal from './components/file-viewer-modal';
import Banner from './components/banner';
import { SavedFile } from './api/types';
import FilesList from './components/files-list';

function App() {
  const [savedFiles, setSavedFiles] = useState<SavedFile[]>([])
  const [selectedFile, setSelectedFile] = useState<string>("")
  const [isFileOpen, setIsFileOpen] = useState<boolean>(false)
  const api = new Api

  useEffect(() => {
    (async () => {
      const files = await api.getAllFiles()
      setSavedFiles(files)
    })()
  })

  const openFile = async (id: string) => {
    const file = await api.getFileById(id)
    setSelectedFile(file)
    setIsFileOpen(true)
  }

  return (
    <div className="flex w-screen h-screen">
      <FileViewerModal isFileOpen={isFileOpen} setIsFileOpen={setIsFileOpen} file={selectedFile} />
      <div className="flex flex-col justify-center items-center text-3xl gap-4 p-4 text-white w-1/2 bg-gradient-to-b from-[#1D1D42] to-[#4E2ECF]">
        <Banner />
      </div>
      <div className="flex flex-col gap-4 w-1/2 bg-white justify-center items-center">
        <div className="w-1/2 flex flex-col gap-2">
          <FilesList openFile={openFile} files={savedFiles} />
        </div>
        <FileUploadForm />
      </div>
    </div>
  );
}

export default App;
