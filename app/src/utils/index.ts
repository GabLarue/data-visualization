const checkFileLines = (lines: string[]): string[] => {
    var spliceToIndex = -1;

    for (var i = 1; i < lines.length; i++) {
        if (lines[i].split(',').length !== lines[i - 1].split(',').length) {
            spliceToIndex = i
        }
    }

    if (spliceToIndex !== -1) {
        lines.splice(0, spliceToIndex)
    }

    return lines
}

export const checkFileContent = (file: File): Promise<File> => {
    const reader = new FileReader()

    return new Promise((resolve) => {
        reader.readAsText(file)

        reader.onload = () => {
            if (typeof reader.result === 'string') {
                var lines = reader.result.split('\n')

                lines = checkFileLines(lines)

                var cleanedFile = new File(lines, file.name, { type: file.type, lastModified: Date.now() })

                resolve(cleanedFile)
            }
        }
    })
}