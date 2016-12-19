package main

import (
	"fmt"
  	//"github.com/gorilla/mux"
    "goji.io/pat"
	"net/http"
    "html/template"
	"io/ioutil"
	"io"
	"encoding/json"
)

var verif_code string
var username string
var password string

func verifcode (w http.ResponseWriter, r *http.Request) {

	_, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }

    //vars := mux.Vars(r)
	//verif_code = vars["verification_code"]
    verif_code = pat.Param(r, "verification_code")

	fmt.Println(verif_code)
	//fmt.Fprintf(w, "verification code    : ")
	//fmt.Fprintf(w, verifcode)

    //fmt.Println("method:", r.Method) //get request method

    t, _ := template.ParseFiles("auth.html")
    t.Execute(w, nil)




}

func inputauth (w http.ResponseWriter, r *http.Request) {
	    r.ParseForm()
        // logic part of log in
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])

        username = r.FormValue("username")
		password = r.FormValue("password")

RepoCreateInput(Input{Code: code, Number: numstring, Transport: transport, Token: verif_code, Username: username, Password: password})

w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    //w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(inputs); err != nil {
        panic(err)
    }

}
	

