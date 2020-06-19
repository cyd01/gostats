package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/shirou/gopsutil/cpu"
)

type Cpu struct {
    Count          int            `json:"count"`
    Info[]         cpu.InfoStat   `json:"info"`
    Times          cpu.TimesStat  `json:"times"`
    TimesPerCpu[]  cpu.TimesStat  `json:"timespercpu"`
}

func CalcCpu() ([]byte) {
    p := Cpu{}
    p.Count,_ = cpu.Counts(true)
    p.Info,_ = cpu.Info()
    t,_ := cpu.Times(false)
    p.Times = t[0]
    p.TimesPerCpu,_ = cpu.Times(true)
   
    b, err := json.Marshal( p )
    if err != nil {
        fmt.Println("Error:", err)
    }

    return b
}

func SrvCpu(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Content-Type", "application/json; charset=UTF-8")
	h.Set("Expires", "0")
	w.Write([]byte("{\"cpu\":"+string(CalcCpu())+"}"))
	logger(r)
}
