package main

import (
    "log"
    "net/http"
    
    //"github.com/resfulGOJI"

)

func main() {
	//NewRouter()
    router := NewRouter()

    //log.Fatal(http.ListenAndServe(":8080", router))

    err := http.ListenAndServe(":9090", router) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
