package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	example := User{
		Name: "Ujjawal",
		Age:  23,
	}

	switch r.Method {
	case "GET":
		err := json.NewEncoder(w).Encode(example)
		if err != nil {
			fmt.Println("error in encoding", err.Error())
		}

	case "POST":
		userInfo := User{}
		err := json.NewDecoder(r.Body).Decode(&userInfo)
		if err != nil {
			fmt.Println("error in decoding", err.Error())
		}
		fmt.Fprint(w, userInfo)

	default:
		fmt.Fprint(w, "This method is not supported.")
	}
}

func main() {
	fmt.Println("Staring Server")

	http.HandleFunc("/xyz", handler)
	http.ListenAndServe("localhost:8080", nil)
}
