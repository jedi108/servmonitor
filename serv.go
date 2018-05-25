package main

import (
	"fmt"
	"time"

	"github.com/jedi108/servmonitor/stathost"
	"github.com/jedi108/servmonitor/storage"
)

func main() {
	stor, err := storage.NewStorage(storage.GetFileNameFromDate())
	if err != nil {
		panic(err)
	}
	hs := stathost.NewHostStat()
	for {
		t := time.Now()
		hs.Cp()
		hs.MemUsed()
		stor.InsertData([]string{t.Format("20060102150405"), fmt.Sprintf("%v", hs.CPUUsedPercent),
			fmt.Sprintf("%v", hs.MemUsedPercent)})
		// fmt.Println(hs.CPUUsedPercent, hs.MemUsedPercent)
		time.Sleep(2 * time.Second)
	}
}
