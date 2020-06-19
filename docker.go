package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/shirou/gopsutil/docker"
)

type Docker struct {
    DockerList[]       string                   `json:"dockerlist"`
    Dockerstats[]      docker.CgroupDockerStat  `json:"dockerstats"`
    DockerMenStats[]*  docker.CgroupMemStat 
}

func CalcDocker() ([]byte) {
    p := Docker{}
    
    p.DockerList,_ = docker.GetDockerIDList()
    p.Dockerstats,_ = docker.GetDockerStat()
    
    if( len(p.Dockerstats)>0 ) {
        p.DockerMenStats=make( []*docker.CgroupMemStat,len(p.Dockerstats) )
	for i:=0 ; i<len(p.Dockerstats) ; i++ {
	    p.DockerMenStats[i],_ = docker.CgroupMemDocker( p.Dockerstats[i].ContainerID )
	}
    }
    
    b, err := json.Marshal( p )
    if err != nil {
        fmt.Println("Error:", err)
    }

    return b
}

func SrvDocker(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Content-Type", "application/json; charset=UTF-8")
	h.Set("Expires", "0")
	w.Write([]byte("{\"docker\":"+string(CalcDocker())+"}"))
	logger(r)
}
