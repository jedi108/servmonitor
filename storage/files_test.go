package storage

import (
	"os"
	"testing"
)

func TestCreateAndRemoveFile(t *testing.T) {
	fileName := GetFileNameFromDate() + ".tmp"
	fA, err := os.Create(fileName)
	if err != nil {
		t.Fatal(err)
	}
	fA.Close()
	if IsFileExists(fileName) {
		os.Remove(fileName)
	}
	if IsFileExists(fileName) {
		t.Fatal("File exists")
	}
}
