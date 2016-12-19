package main


type Input struct {
    Id        		int       	`json:"id"`
    Transport		string    	`json:"transport"`
    Code 			string      `json:"code"`
    Number       	string 		`json:"number"`
    Token			string		`json:"verifcode"`
    Username		string		`json:"username"`
    Password		string		`json:"Password"`
}

type Inputs []Input
