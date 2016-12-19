package main

import (
	"fmt"
  	"regexp"
  	"strconv"
  	//"github.com/gorilla/mux"
  	//"goji.io"
    "goji.io/pat"
	"net/http"
	"io/ioutil"
	"io"
)
var code string
var numstring string
var transport string

func regex(numstring string) int {
	 match, _ := regexp.MatchString("^8[0-9]{9,11}$", numstring )
    fmt.Println("regexp: ", match)
    if match == true {
    	respons := 200
    	return respons
    } else {
    	respons := 400
    	return respons
    	//w.WriteHeader(400)
    }
}

func verif (w http.ResponseWriter, r *http.Request) {

	_, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
       //w.WriteHeader(413)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
        //w.WriteHeader(413)
    }

	code = pat.Param(r, "code")
	num := pat.Param(r, "number")
	transport = pat.Param(r, "transport")
    //vars := mux.Vars(r)
	//code = vars["code"]
	//num := vars["number"]
	//transport = vars["transport"]

	fmt.Fprintf(w, "Code            : %s \n", code)
	fmt.Fprintf(w, "Original number : %s \n", num)
	fmt.Fprintf(w, "Transport       : %s \n", transport)

	var numint int
	if _, err := fmt.Sscanf(num, "%14d", &numint); err != nil {
			fmt.Println("err")
	}
	numstring = strconv.Itoa(numint)
	finalnum :=code+numstring

	fmt.Println(finalnum)
	fmt.Fprintf(w, "Phone Number    : ")
	fmt.Fprintf(w, finalnum)


	if (transport == "sms") || (transport == "voice") {
		returncode := strconv.Itoa(regex(numstring))
		fmt.Fprintf(w, "\nReturn ")
		fmt.Fprintf(w, returncode)	
	} else {
		var response = "415"
		fmt.Fprintf(w, "\nReturn:")
		fmt.Fprintf(w, response)
		//w.WriteHeader(415)
	}
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

}

func wrong (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "need input transport type, code, phone number")
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}
