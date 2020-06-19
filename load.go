package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/shirou/gopsutil/load"
)

type Load struct {
    Avg*   load.AvgStat   `json:"avg"`
    Misc*  load.MiscStat  `json:"misc"`
}

func CalcLoad() ([]byte) {
    p := Load{}
    
    p.Avg,_ = load.Avg()
    p.Misc,_ = load.Misc()
    
    b, err := json.Marshal( p )
    if err != nil {
        fmt.Println("Error:", err)
    }

    return b
}

func SrvLoad(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Content-Type", "application/json; charset=UTF-8")
	h.Set("Expires", "0")
	w.Write([]byte("{\"load\":"+string(CalcLoad())+"}"))
	logger(r)
}
