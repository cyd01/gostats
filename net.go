package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/shirou/gopsutil/net"
)

type Net struct {
    Conntracks[]              net.ConntrackStat   `json:"conntracks"`
    Filters[]                 net.FilterStat      `json:"filters"`
    Interfaces[]              net.InterfaceStat   `json:"interfaces"`
    IOCounters                net.IOCountersStat  `json:"iocounters"`
    IOCountersPerInterface[]  net.IOCountersStat  `json:"iocountersperinterface"`
}

func CalcNet() ([]byte) {
    p := Net{}
    
    p.Conntracks,_ = net.ConntrackStats(true)
    p.Interfaces,_ = net.Interfaces()
    p.Filters,_ = net.FilterCounters()
    t,_ := net.IOCounters(false)
    p.IOCounters = t[0]
    p.IOCountersPerInterface,_ = net.IOCounters(true)
   
    b, err := json.Marshal( p )
    if err != nil {
        fmt.Println("Error:", err)
    }

    return b
}

func SrvNet(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Content-Type", "application/json; charset=UTF-8")
	h.Set("Expires", "0")
	w.Write([]byte("{\"net\":"+string(CalcNet())+"}"))
	logger(r)
}
