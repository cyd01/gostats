package main

import (
    b64 "encoding/base64"
    "net/http"
)

func FavSrvHandler(w http.ResponseWriter, r *http.Request) {
    h := w.Header()
    h.Set("Content-Type", "image/x-icon")
    h.Set("Cache-Control", "public, max-age=300")
    h.Set("X-Server", "Go HTTP Server")
    
    fav := "AAABAAEAEBAAAAEAGABoAwAAFgAAACgAAAAQAAAAIAAAAAEAGAAAAAAAAAAAAGAAAABgAAAAAAAAAAAAAACA5v+A5v+A5v+A5v+A5v+A5v+A5v+A5v9f041f041f041f041f041f041f041f042A5v+A5v+A5v+A5v+A5v+A5v+A5v+A5v9f041f041f041f041f041f041f041f042A5v+A5v+A5v+A5v+A5v+A5v+A5v+A5v9f041f041f041f041f041f041f041f042A5v+A5v+A5v+A5v+A5v+A5v+A5v+A5v9f041f041f041f041f041f041f041f042A5v+A5v+A5v+A5v+A5v+A5v+A5v+A5v9f041f041f041f041f041f041f041f042A5v+A5v+A5v+A5v+A5v+A5v+A5v+A5v9f041f041f041f041f041f041f041f042A5v+A5v+A5v+A5v+A5v+A5v+A5v+A5v9f041f041f041f041f041f041f041f042A5v+A5v+A5v+A5v+A5v+A5v+A5v+A5v9f041f041f041f041f041f041f041f043/s4D/s4D/s4D/s4D/s4D/s4D/s4D/s4CAgP+AgP+AgP+AgP+AgP+AgP+AgP+AgP//s4D/s4D/s4D/s4D/s4D/s4D/s4D/s4CAgP+AgP+AgP+AgP+AgP+AgP+AgP+AgP//s4D/s4D/s4D/s4D/s4D/s4D/s4D/s4CAgP+AgP+AgP+AgP+AgP+AgP+AgP+AgP//s4D/s4D/s4D/s4D/s4D/s4D/s4D/s4CAgP+AgP+AgP+AgP+AgP+AgP+AgP+AgP//s4D/s4D/s4D/s4D/s4D/s4D/s4D/s4CAgP+AgP+AgP+AgP+AgP+AgP+AgP+AgP//s4D/s4D/s4D/s4D/s4D/s4D/s4D/s4CAgP+AgP+AgP+AgP+AgP+AgP+AgP+AgP//s4D/s4D/s4D/s4D/s4D/s4D/s4D/s4CAgP+AgP+AgP+AgP+AgP+AgP+AgP+AgP//s4D/s4D/s4D/s4D/s4D/s4D/s4D/s4CAgP+AgP+AgP+AgP+AgP+AgP+AgP+AgP8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
    sDec, _ := b64.StdEncoding.DecodeString(fav)
    
    w.Write(sDec)
    logger(r)
}
