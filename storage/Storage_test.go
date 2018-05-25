package storage

import (
	"os"
	"testing"
)

func TestInsert(t *testing.T) {
	fileTest := "test_file.csv"
	stor, err := NewStorage(fileTest)
	if err != nil {
		t.Fatal(err)
	}
	var data = []string{"Line3", "Hello Readers of", "111"}

	err = stor.InsertData(data)
	if err != nil {
		t.Fatal(err)
	}

	stor.startDay = stor.startDay - 1

	err = stor.InsertData(data)
	if err != nil {
		t.Fatal(err)
	}

	stor.Close()
	stor.fappend.Close()
	stor.csvWriter.Flush()
	err = os.Remove(fileTest)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Remove(GetFileNameFromDate())
	if err != nil {
		t.Fatal(err)
	}
}

// TEST WITH POPLINE
//
// func TestInsert(t *testing.T) {
// 	stor, err := NewStorage("test_file.csv")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	var data = [][]string{{"Line3", "Hello Readers of"}, {"Line2", "golangcode.com"}}

// 	err = stor.InsertData(data)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	err = stor.PopLine()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	stor.Close()
// }
