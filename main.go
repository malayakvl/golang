package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"strconv"
	"math"
	"encoding/json"
)

func strToIntArr(s string) []int {
	strs := strings.Split(s, ",")
	res := make([]int, len(strs))
	for i := range res {
		_, err := strconv.Atoi(strs[i])
	    if err != nil {
	        panic(fmt.Sprintf("the given input is invalid. Element on index(%i) is not a number", i))
	        break 
	    } else {
	    	res[i], _ = strconv.Atoi(strs[i])
	    }
	}
	return res
}


func IsPrime(value int) bool {
    for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
        if value%i == 0 {
            return false
        }
    }
    return value > 1
}



func checkArray(w http.ResponseWriter, r *http.Request) {
	var result = []bool{}
	if r.URL.Path != "/home" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	errForm := r.ParseForm()
	if errForm != nil {
		log.Fatal(errForm)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	
	switch r.Method {
		case "GET":		
			 fmt.Fprintf(w, "Sorry, only POST methods are supported.")
		case "POST":
			values := r.URL.Query()
			var _strData = values["dataCheck"]
			justString := strings.Join(_strData,",")
			var dataArrayCheck = strToIntArr(justString)

		    for i := 0; i < len(dataArrayCheck); i++ {
		        if IsPrime(dataArrayCheck[i]) {
		        	result = append(result, true)
		        } else {
		        	result = append(result, false)
		        }
		    }
			jsonResp, err := json.Marshal(result)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			w.Write(jsonResp)
			return
		default:
			fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/home", checkArray)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":3002", nil); err != nil {
		log.Fatal(err)
	}
}