package fsfire

type FileExtension string

func (file FileExtension) String() string {
	return string(file)
}

const (
	FileExtensionWithImageOfJPEG FileExtension = "jpeg"
	FileExtensionWithImageOfPNG  FileExtension = "png"
	FileExtensionWithImageOfGIF  FileExtension = "gif"
	FileExtensionWithImageOfBMP  FileExtension = "bmp"
)

const (
	FileExtensionWithOfficeOfExcel FileExtension = "xlsx"
	FileExtensionWithOfficeOfWord  FileExtension = "docx"
	FileExtensionWithOfficeOfCSV   FileExtension = "csv"
)

const (
	FileExtensionWithDocumentOfMarkdown FileExtension = "md"
	FileExtensionWithDocumentOfFileText FileExtension = "txt"
)
