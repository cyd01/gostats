package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/process"
)

type Process struct {
    Background     bool                     `json:"background"`
    Cmdline        string                   `json:"command"`
    CPUPercent     float64                  `json:"cpupercent"`
    CreateTime     int64                    `json:"createtime"`
    Cwd            string                   `json:"cwd"`
    Exe            string                   `json:"exe"`
    Foreground     bool                     `json:"foreground"`
    IsRunning      bool                     `json:"isrunning"`
    MemoryInfo*    process.MemoryInfoStat  `json:memoryinfo`
    MemoryPercent  float32                  `json:"memorypercent"`
    Name           string                   `json:"name"`
    Nice           int32                    `json:"nice"`
    NumFDs         int32                    `json:"numfds"`
    NumThreads     int32                    `json:"numthreads"`
    Pid            int32                    `json:"ppid"`
    Ppid           int32                    `json:"ppid"`
    Status         string                   `json:"status"`
    Terminal       string                   `json:"terminal"`
    Times*         cpu.TimesStat           `json:"times"`
    Username       string                   `json:"username"`
}

func CalcProcess() ([]byte) {
    
    ProcessList,_ := process.Processes()
    p := make( []Process, len(ProcessList) )
    
    if( len(ProcessList)>0 ) { 
        for i:=0 ; i<len(ProcessList) ; i++ {
	    p[i].Background,_ = ProcessList[i].Background()
	    p[i].Cmdline,_ = ProcessList[i].Cmdline()
	    p[i].CPUPercent,_ = ProcessList[i].CPUPercent()
	    p[i].Cwd,_ = ProcessList[i].Cwd()
	    p[i].Exe,_ = ProcessList[i].Exe()
	    p[i].Foreground,_ = ProcessList[i].Foreground()
	    p[i].IsRunning,_ = ProcessList[i].IsRunning()
	    p[i].MemoryInfo,_ = ProcessList[i].MemoryInfo()
            p[i].MemoryPercent,_ = ProcessList[i].MemoryPercent()
            p[i].Name,_ = ProcessList[i].Name()
	    p[i].Nice,_ = ProcessList[i].Nice()
	    p[i].NumFDs,_ = ProcessList[i].NumFDs()
	    p[i].NumThreads,_ = ProcessList[i].NumThreads()
            p[i].Pid = ProcessList[i].Pid
	    p[i].Ppid,_ = ProcessList[i].Ppid()
	    p[i].Status,_ = ProcessList[i].Status()
	    p[i].Terminal,_ = ProcessList[i].Terminal()
	    p[i].Times,_ = ProcessList[i].Times()
	    p[i].Username,_ = ProcessList[i].Username()
        } 
    }
    
    b, err := json.Marshal( p )
    if err != nil {
        fmt.Println("Error:", err)
    }

    return b
}

func SrvProcess(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Content-Type", "application/json; charset=UTF-8")
	h.Set("Expires", "0")
	w.Write([]byte("{\"processes\":"+string(CalcProcess())+"}"))
	logger(r)
}
