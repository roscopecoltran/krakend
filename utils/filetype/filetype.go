package filetype

import (
	"fmt"
	"gopkg.in/h2non/filetype.v1"
	"io/ioutil"
)

func GetType(filePath string) (string, string) {
	buf, _ := ioutil.ReadFile(filePath)
	kind, unknown := filetype.Match(buf)
	if unknown != nil {
		fmt.Printf("Unknown: %s", unknown)
		return unknown
	}
	fmt.Printf("File path: %s, File type: %s, MIME: %s\n", filePath, kind.Extension, kind.MIME.Value)
	return kind.Extension, kind.MIME.Value
}

func GetTypeByClass(filePath string, expectedClass string) bool {
	buf, _ := ioutil.ReadFile(filePath)
	fmt.Printf("File path: %s\n", filePath)
	fmt.Printf("Expected file class: %s", expectedClass)

	if filetype.IsImage(buf) {
		fmt.Println("File IS an image")
		return true
	}
	fmt.Printf("Not an image, filepath: %s", filePath)
	return false
}

func IsSupportedExtension(fileExtension string) bool {
	fmt.Printf("File extension: %s: %s\n", fileExtension)
	// Check if file is supported by extension
	if filetype.IsSupported(fileExtension) {
		fmt.Printf("Extension IS supported, Extension: %s\n", fileExtension)
		return true
	}
	fmt.Printf("Extension NOT supported, Extension: %s\n", fileExtension)
	return false
}

func IsSupportedMimeType(mimeType string) bool {
	// Check if file is supported by extension
	if filetype.IsMIMESupported(mimeType) {
		fmt.Printf("MIME type IS supported, Mime Type: %s\n", mimeType)
		return true
	}
	fmt.Printf("MIME type NOT supported, Mime Type: %s\n", mimeType)
	return false
}

func IsFileImage(filePath string) bool {
	// Read a file
	buf, _ := ioutil.ReadFile(filePath)
	// We only have to pass the file header = first 261 bytes
	head := buf[:261]
	if filetype.IsImage(head) {
		fmt.Println("File IS an image")
		return true
	}
	fmt.Println("Not an image")
	return false
}
