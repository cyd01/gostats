package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/shirou/gopsutil/disk"
)

type Disk struct {
    Device      string           `json:"device"`
    Mountpoint  string           `json:"mountpoint"`
    Fstype      string           `json:"fstype"`
    Opts        string           `json:"opts"`
    Usage*      disk.UsageStat  `json:"usage"`
}

func CalcDisk() ([]byte) {
    
    PartitionsList,_ := disk.Partitions(false)
    p := make( []Disk, len(PartitionsList) )
    
    if( len(PartitionsList)>0 ) { 
        for i:=0 ; i<len(PartitionsList) ; i++ {
	    p[i].Device = PartitionsList[i].Device
	    p[i].Mountpoint = PartitionsList[i].Mountpoint
	    p[i].Fstype = PartitionsList[i].Fstype
	    p[i].Opts = PartitionsList[i].Opts
	    p[i].Usage,_ = disk.Usage( p[i].Mountpoint )
	}
    }

    b, err := json.Marshal( p )
    if err != nil {
        fmt.Println("Error:", err)
    }

    return b
}

func SrvDisk(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Content-Type", "application/json; charset=UTF-8")
	h.Set("Expires", "0")
	w.Write([]byte("{\"disks\":"+string(CalcDisk())+"}"))
	logger(r)
}
