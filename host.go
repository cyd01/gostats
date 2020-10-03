package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "github.com/shirou/gopsutil/host"
)

type Host struct {
    Hostname       string                `json:"hostname"`
    Info*          host.InfoStat         `json:"info"`
    Temperature[]  host.TemperatureStat  `json:"temperature"`
    Users[]        host.UserStat         `json:"users"`
}

func gethostfile() string {
	var hostfile = ""
	hostenv := os.Getenv("HOST_ETC")
	if( len(hostenv)>0 ) {
		if b,_:= exists( hostenv+"/hostname" ); b {
			hostfile = hostenv+"/hostname"
		}
	}
	if(len(hostfile)==0) {
		if b,_:=exists("/data/etc/hostname");b {
			hostfile = "/data/etc/hostname"
		}
	}
	return hostfile
}

func gethostname() string {
   hostfile := gethostfile()
    if( len(hostfile)>0 ) {
        return readfile(hostfile)
    } else {
	i,_ := host.Info()
	return i.Hostname
    }
}

func CalcHost() ([]byte) {
    p := Host{}
    
    p.Info,_ = host.Info()
    p.Hostname = gethostname()

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
