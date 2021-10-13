package main

import (
    "fmt"
    "io/ioutil"
)

// temporary directory
const tmpDir = "apps"

func main() {
    
    // get files from `/tmp` directory
    files, _ := ioutil.ReadDir( tmpDir )

    // list information of each file
    for _, file := range files {
        
        fmt.Printf( "Name: %v, Size: %v kb, Mode: %v, IsDir: %v\n",
            file.Name(),
            file.Size() / 1000,
            file.Mode(),
            file.IsDir(),
        )
    }
}



