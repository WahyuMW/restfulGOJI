package main

import (
    "net/http"

    //"github.com/gorilla/mux"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route
/*
func NewRouter() *mux.Router {

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
    }

    return router
}
*/
var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "NeedInput",
        "GET",
        "/v1/account/",
        wrong,
    },
    Route{
        "Verify",
        "GET",
        "/v1/account/{transport}/{code}/{number}",
        verif,
    },

    Route{
        "VerifyCode",
        "GET",
        "/v1/accounts/code/{verification_code}",
        verifcode,
    },

    Route{
        "UserAuth",
        "POST",
        "/v1/accounts/code/{verification_code}",
        inputauth,
    },
}