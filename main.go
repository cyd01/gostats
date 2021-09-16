package main

import (
    "fmt"
    "flag"
    "log"
    "net/http"
    "os"
    "strconv"
)

func main() {
    flag.Parse()
    
    if( flag.NArg()>=2 ) {

    port,err := strconv.Atoi( flag.Arg(0) )
	if( err!=nil ) {
	    fmt.Println( flag.Arg(0),"is not a number" )
	    os.Exit(1)
	}
	
	dir := flag.Arg(1)
	if b,_ := existsdir(dir); !b {
	    fmt.Println( "Can not find directory", dir )
	    os.Exit(2)
	}

	http.HandleFunc("/favicon.ico", FavSrvHandler)
	http.HandleFunc(os.Getenv("GOSTATS_BASE")+"/cpu", SrvCpu )
	http.HandleFunc(os.Getenv("GOSTATS_BASE")+"/disk", SrvDisk )
	http.HandleFunc(os.Getenv("GOSTATS_BASE")+"/docker", SrvDocker )
	http.HandleFunc(os.Getenv("GOSTATS_BASE")+"/host", SrvHost )
	http.HandleFunc(os.Getenv("GOSTATS_BASE")+"/info", SrvInfo )
	http.HandleFunc(os.Getenv("GOSTATS_BASE")+"/load", SrvLoad )
	http.HandleFunc(os.Getenv("GOSTATS_BASE")+"/mem", SrvMem )
	http.HandleFunc(os.Getenv("GOSTATS_BASE")+"/net", SrvNet )
	http.HandleFunc(os.Getenv("GOSTATS_BASE")+"/process", SrvProcess )
	http.HandleFunc(os.Getenv("GOSTATS_BASE")+"/usage", func(w http.ResponseWriter, r *http.Request) {
	    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        fmt.Fprintf(w, "[\"/cpu\",\"/disk\",\"/docker\",\"/host\",\"info\",\"/load\",\"/mem\",\"/net\",\"/process\",\"/usage\"]")
	} )

	if len(os.Getenv("GOSTATS_BASE"))>0 {
		/*
		http.HandleFunc(os.Getenv("GOSTATS_BASE")+"/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, dir+"/"+r.URL.Path[1:])
		} )
		*/
		http.Handle(os.Getenv("GOSTATS_BASE")+"/", http.StripPrefix(os.Getenv("GOSTATS_BASE")+"/",http.FileServer(http.Dir(dir))))
	} else {
		http.Handle(os.Getenv("GOSTATS_BASE")+"/", http.FileServer(http.Dir(dir)))
	}

	log.Println("Starting webserver on port", port, "to directory", dir)
	if len(os.Getenv("GOSTATS_BASE"))>0 { log.Println("With prefix "+os.Getenv("GOSTATS_BASE")) }
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))

    } else if( flag.NArg()>=1 ) {
	if( flag.Arg(0)=="cpu" ) { fmt.Println( string(CalcCpu()) )
	} else if( flag.Arg(0)=="disk" ) { fmt.Println( string(CalcDisk()) )
	} else if( flag.Arg(0)=="docker" ) { fmt.Println( string(CalcDocker()) )
	} else if( flag.Arg(0)=="host" ) { fmt.Println( string(CalcHost()) )
	} else if( flag.Arg(0)=="info" ) { fmt.Println( string(CalcInfo()) )
	} else if( flag.Arg(0)=="load" ) { fmt.Println( string(CalcLoad()) )
	} else if( flag.Arg(0)=="mem" ) { fmt.Println( string(CalcMem()) )
	} else if( flag.Arg(0)=="net" ) { fmt.Println( string(CalcNet()) )
	} else if( flag.Arg(0)=="process" ) { fmt.Println( string(CalcProcess()) )
	} else {
	    fmt.Println( "Unknown function" )
	}
    } else {
	
	fmt.Println( "Usage:" )
	fmt.Println( "    prog func          ... func is on of [ cpu, disk, docker, host, load, mem, net, process ]" )
	fmt.Println( "    prog port dir" )
    }

    
}


/*
https://github.com/shirou/gopsutil/
http://godoc.org/github.com/shirou/gopsutil/

You can set an alternative location to /proc by setting the HOST_PROC environment variable.

You can set an alternative location to /sys by setting the HOST_SYS environment variable.

You can set an alternative location to /etc by setting the HOST_ETC environment variable.

You can set an alternative location to /var by setting the HOST_VAR environment variable.

You can set an alternative location to /run by setting the HOST_RUN environment variable.

You can set an alternative location to /dev by setting the HOST_DEV environment variable.
*/
