package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/shirou/gopsutil/host"
)

type Host struct {
    Info*          host.InfoStat          `json:"info"`
    Temperature[]  host.TemperatureStat  `json:"temperature"`
    Users[]        host.UserStat         `json:"users"`
}

func CalcHost() ([]byte) {
    p := Host{}
    
    p.Info,_ = host.Info()
    p.Temperature,_ = host.SensorsTemperatures()
    p.Users,_ = host.Users()
    
    b, err := json.Marshal( p )
    if err != nil {
        fmt.Println("Error:", err)
    }

    return b
}

func SrvHost(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Content-Type", "application/json; charset=UTF-8")
	h.Set("Expires", "0")
	w.Write([]byte("{\"host\":"+string(CalcHost())+"}"))
	logger(r)
}
