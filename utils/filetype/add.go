package filetype

import (
	"fmt"
	"gopkg.in/h2non/filetype.v1"
)

func newMatcher(buf []byte) bool {
	return len(buf) > 1 && buf[0] == 0x01 && buf[1] == 0x02
}

func Add(extFile string, mimeType string) (bool, error) {
	newType := filetype.NewType(extFile, mimeType)
	// Register the new matcher and its type
	filetype.AddMatcher(newType, newMatcher)
	if ok := IsSupportedCustomExtType(extFile); !ok {
		return false, err
	}
	if ok := IsSupportedCustomMimeType(mimeType); !ok {
		return false, err
	}
	return true, nil
}

func IsSupportedCustomExtType(extFile string) bool {
	// Check if the new type is supported by extension
	if filetype.IsSupported(extFile) {
		fmt.Println("New supported type: foo")
		return true
	}
	return false
}

func IsSupportedCustomMimeType(mimeType string) bool {
	// Check if the new type is supported by mime type
	if filetype.IsSupported(mimeType) {
		fmt.Printf("New supported type: %s \n", mimeType)
		return true
	}
	return false
}

func IsDetectedCustomFile(filePath string, expectFileType string) bool {
	buf, _ := ioutil.ReadFile(filePath)
	fmt.Printf("File path: %s\n", filePath)
	fmt.Printf("Expected file class: %s", expectFileType)
	// Try to match the file
	kind, _ := filetype.Match(buf)
	if kind == filetype.Unknown {
		fmt.Println("Unknown file type")
		return false
	}
	if expectFileType == kind.Extension {
		fmt.Printf("File path: %s, type IS matched: %s\n", filePath, expectFileType, kind.Extension)
		return true
	}
	fmt.Printf("File path: %s, type NOT matched: %s\n", filePath, expectFileType, kind.Extension)
	return false
}
