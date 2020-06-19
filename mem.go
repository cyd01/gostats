package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/shirou/gopsutil/mem"
)

type Mem struct {
    Swap*       mem.SwapMemoryStat       `json:"swap"`
    Virtual*    mem.VirtualMemoryStat    `json:"virtual"`
    VirtualEx*  mem.VirtualMemoryExStat  `json:"virtualmemex"`
}

func CalcMem() ([]byte) {
    p := Mem{}
    p.Swap,_ = mem.SwapMemory()
    p.Virtual,_ = mem.VirtualMemory()
    p.VirtualEx,_ = mem.VirtualMemoryEx()
   
    b, err := json.Marshal( p )
    if err != nil {
        fmt.Println("Error:", err)
    }

    return b
}

func SrvMem(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Content-Type", "application/json; charset=UTF-8")
	h.Set("Expires", "0")
	w.Write([]byte("{\"mem\":"+string(CalcMem())+"}"))
	logger(r)
}
