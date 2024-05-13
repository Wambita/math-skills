package mathskills

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestOpenFile(t *testing.T) {
	// test for valid file path
	inputFile := "data.txt"
	_, err := OpenFile(inputFile)
	if err != nil {
		t.Errorf("Failed to open existing file: %v", err)
	}
}

// test for non existent file
func TestNonExistentFile(t *testing.T) {
	inputFile := "non-existent.txt"
	_, err := OpenFile(inputFile)
	if err == nil {
		t.Errorf("Expected error for non existent file, but got none")
	}
}

// test for wrong extension
func TestWrongExtension(t *testing.T) {
	inputFilepath := "untitled.png"
	if filepath.Ext(inputFilepath) == ".txt" {
		t.Errorf("Error: file has .txt expected wrong filepath")
	}
}

// test for empty file
func TestEmptyfile(t *testing.T) {
	inputFilepath := "empty.txt"
	fileinfo, err := os.Stat(inputFilepath)
	if fileinfo.Size() != 0 {
		t.Errorf("Expected error for empty file, found none")
	}
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
