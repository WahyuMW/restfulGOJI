package main

import (
    //"net/http"
    "goji.io"
    "goji.io/pat"
    //"github.com/gorilla/mux"
)

func NewRouter() *goji.Mux {


    mux := goji.NewMux()
    mux.HandleFunc(pat.Get("/"), Index)
    mux.HandleFunc(pat.Get("/v1/account/"), wrong)
    mux.HandleFunc(pat.Get("/v1/account/:transport/:code/:number"), verif)
    mux.HandleFunc(pat.Get("/v1/accounts/code/:verification_code"), verifcode)
    mux.HandleFunc(pat.Post("/v1/accounts/code/:verification_code"), inputauth)

/*func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        handler = route.HandlerFunc
        handler = Logger(handler, route.Name)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)

    }*/
    return mux
}