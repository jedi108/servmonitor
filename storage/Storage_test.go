package storage

import "testing"

func TestInsert(t *testing.T) {
	stor, err := NewStorage("test_file.csv")
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
