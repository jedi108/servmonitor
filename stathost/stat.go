package stathost

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type HostStat struct {
	CPUUsedPercent int
	MemUsedPercent int
	HARDfree       int
	meminit        *mem.VirtualMemoryStat
	iOCountersStat *map[string]disk.IOCountersStat
}

func NewHostStat() *HostStat {
	hs := &HostStat{}
	mem, _ := mem.VirtualMemory()
	hs.meminit = mem

	// ret, err := disk.IOCounters()
	// if err != nil {
	// 	panic(err)
	// }
	// hs.iOCountersStat = &ret

	return hs
}

func (hs *HostStat) Cp() {
	cpuPer, _ := cpu.Percent(0, false)
	hs.CPUUsedPercent = int(cpuPer[0])
}

func (hs *HostStat) MemUsed() {
	hs.MemUsedPercent = int(hs.meminit.UsedPercent)
}

/*
todo: disks
*/
func (hs *HostStat) DiskStat() {
	for part, io := range *hs.iOCountersStat {
		fmt.Println(part, io)
	}
}

func (hs *HostStat) DiskUsage() {
	hd, _ := disk.Usage("C:")
	// mb
	hs.HARDfree = int(hd.Free) / 1024 / 1024
}
