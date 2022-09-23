import React, { useEffect, useState } from 'react';
import './App.css';
import axios from 'axios'

type SavedFile = {
  name: string,
  key: string
}

const http = axios.create({
  baseURL: '//localhost:8080',
  headers: {
    "Content-Type": "multipart/form-data",
  }
});

function App() {
  const [savedFiles, setSavedFiles] = useState<SavedFile[]>()
  const [fileToUpload, setFileToUpload] = useState<File | null>(null)
  const [selectedFile, setSelectedFile] = useState<string>("")

  useEffect(() => {
    http.get('/files')
      .then((response) => {
        setSavedFiles(response.data)
      })
      .catch((error) => {
        console.log(error);
      })
  })

  const handleUpload = (e: React.FormEvent) => {
    e.preventDefault()

    if (fileToUpload === null) {
      alert("No file selected for upload!")
    } else {
      var formData = new FormData()
      formData.append('file', fileToUpload)

      http.post('/upload', formData)
        .then((response) => {
          console.log(`File ${response.data} was uploaded successfully!`)
        })
        .catch((error) => {
          console.log(error);
        })
    }
  }

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target && e.target.files) {
      setFileToUpload(e.target.files[0])
    }
  }

  return (
    <div className="flex w-screen h-screen">
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
            return <div className="text-grey flex justify-between hover:text-white rounded px-4 py-2 bg-[#0182FF33] hover:bg-[#0182FFB3] cursor-pointer" key={file.key}>
              <div className="flex truncate gap-2">
                <span className="material-symbols-outlined">folder</span>
                <span className="truncate">{file.name}</span>
              </div>
              <span className="material-symbols-outlined">more_vert</span>
            </div>
          }) :
            <div className="text-white rounded px-4 py-2 bg-[#00499033]">
              <span>No files were uploaded yet...</span>
            </div>
          }
        </div>
        <form id={"form"} className="flex flex-col gap-2 w-1/2" onSubmit={e => handleUpload(e)}>
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
