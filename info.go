package main

import (
    "bufio"
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
    "strconv"
)

type Group struct {
	Name    string    `json:"name"`
	Id      int     `json:"id"`
	Users   []string  `json:"users"`
}

type User struct {
	Name         string     `json:"name"`
	Uid          int        `json:"uid"`
	Gid          int        `json:"gid"`
	Info         string     `json:"info"`
	Home         string     `json:"home"`
	Shell        string     `json:"shell"`
}

type Info struct {
	 Hostname      string                `json:"hostname"`
	 Groups        []Group               `json:"groups"`
	 Users         []User                `json:"users"`
}

func getinfofile(filename string) string {
	var infofile = ""
	hostenv := os.Getenv("HOST_ETC")
	if( len(hostenv)>0 ) {
		if b,_:= exists( hostenv+"/"+filename ); b {
			infofile = hostenv+"/"+filename
		}
	}
	if(len(infofile)==0) {
		if b,_:=exists("/data/etc/"+filename);b {
			infofile = "/data/etc/"+filename
		}
	}
	if(len(infofile)==0) {
		if b,_:=exists("/etc/"+filename);b {
			infofile = "/etc/"+filename
		}
	}
	return infofile
}

func getgroups() []Group {
	var p []Group
	
	Infofile := getinfofile("group")
	file, err := os.Open(Infofile)
	defer file.Close()
	if( err == nil ) {
	    reader := bufio.NewReader(file)
	    for {
                var buffer bytes.Buffer
                var l []byte
                var isPrefix bool
                for {
                    l, isPrefix, err = reader.ReadLine()
                    buffer.Write(l)
                    // If we've reached the end of the line, stop reading.
                    if !isPrefix { break; }
                    // If we're just at the EOF, break
                    if err != nil { break; }
                }
                if err == io.EOF { break; }
                line := buffer.String()
		
		s := strings.Split(line,":")
		var t Group
                t.Name  = s[0]
		t.Id,_    = strconv.Atoi(s[2])
		t.Users = make( []string,0 )
		if( len(s[3])>0 ) {
		    var u []string
		    ul := strings.Split(s[3],",")
		    if( len(ul)>0 ) { for i:=0 ; i<len(ul) ; i++ { u = append( u, ul[i] ) } ; t.Users = u }
		}
		p = append(p,t)
            }
	}
	
	return p
}

func getusers() []User {
	var p []User
	
	Infofile := getinfofile("passwd")
	file, err := os.Open(Infofile)
	defer file.Close()
	if( err == nil ) {
	    reader := bufio.NewReader(file)
	    for {
                var buffer bytes.Buffer
                var l []byte
                var isPrefix bool
                for {
                    l, isPrefix, err = reader.ReadLine()
                    buffer.Write(l)
                    // If we've reached the end of the line, stop reading.
                    if !isPrefix { break; }
                    // If we're just at the EOF, break
                    if err != nil { break; }
                }
                if err == io.EOF { break; }
                line := buffer.String()
		s := strings.Split(line,":")
		var t User
                t.Name  = s[0]
		t.Uid,_    = strconv.Atoi(s[2])
		t.Gid,_    = strconv.Atoi(s[3])
		t.Info  = s[4]
		t.Home  = s[5]
		t.Shell = s[6]
		p = append(p,t)
            }
	}
	
	return p
}

func CalcInfo() ([]byte) {
    p := Info{}
    
    p.Hostname = gethostname()
    p.Groups = getgroups()
    p.Users = getusers()

    b, err := json.Marshal( p )
    if err != nil {
        fmt.Println("Error:", err)
    }

    return b
}

func SrvInfo(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Content-Type", "application/json; charset=UTF-8")
	h.Set("Expires", "0")
	w.Write([]byte("{\"info\":"+string(CalcInfo())+"}"))
	logger(r)
}
