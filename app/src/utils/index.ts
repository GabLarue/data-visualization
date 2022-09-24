const checkFileHeader = (lines: string[]): string[] => {
    for (var i = 1; i < lines.length - 1; i++) {
        if(lines[i].length !== lines[i-1].length){
            lines.shift()
        }
    }
    
    return lines
}

export const checkFileContent = (file: File): Promise<File> => {
    const reader = new FileReader()

    return new Promise((resolve, reject) => {
        reader.readAsText(file)

        reader.onload = () => {
            if (typeof reader.result === 'string') {
                var lines = reader.result.split('\n')

                lines = checkFileHeader(lines)

                var cleanedFile = new File(lines, file.name, { type: file.type, lastModified: Date.now() })

                resolve(cleanedFile)
            }
        }
    })
}