package main

import (
    "io/ioutil"
    "os"
)

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}

// existsdir returns whether the given path exists and is a directory
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

// Read all file an return string
func readfile(path string) string {
    file, err := os.Open(path)
    if err != nil {
        return ""
    }
    defer file.Close()
    b, err := ioutil.ReadAll(file)
    return string(b)
}
