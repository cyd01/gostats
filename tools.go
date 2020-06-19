package main

import (
   "os"
)
   
func existsdir(path string) (bool, error) {
    fi, err := os.Stat(path)
    if err == nil {
        if( fi.IsDir() ) {
	    return true, nil
	} else {
	    return true, nil
	}
    }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}
