package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Response struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

var response Response

// Display error if link too deep
func MissingLengthCall(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Response{Code: "1", Message: "Please use append the number of digits you want returned (e.g. <url>/api/3"})
}

func FibonacciToDigits(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	strDigits := params["digits"]
	digits, err := strconv.Atoi(strDigits)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(FibonacciLoop(digits))
	}
}

func FibonacciLoop(n int) []int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f
}

// main function to boot up everything
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/fibonacci", MissingLengthCall).Methods("GET")
	router.HandleFunc("/fibonacci/{digits}", FibonacciToDigits).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
