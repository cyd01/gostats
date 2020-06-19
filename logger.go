package main

import (
    "html"
    "log"
    "net/http"
)

func logger( r *http.Request, s ...string ) {
	ra := r.RemoteAddr
	forward := r.Header.Get("X-Forwarded-For")
	if( len(forward)>0 ) { ra=ra+","+forward }
	if( len(s)>0 ) { log.Println( r.Method, ra, html.EscapeString(r.URL.Path), s ) 
	} else { log.Println( r.Method, ra, html.EscapeString(r.URL.Path), s ) 
	}
}

