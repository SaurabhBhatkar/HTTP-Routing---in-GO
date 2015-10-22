package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MyJson struct {
	Name string `json:"name"`
}

type MyJsonRes struct {
	Salutation string
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func PostHello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

	var myjson3 MyJson
	s3 := json.NewDecoder(req.Body)
	err := s3.Decode(&myjson3)
//	fmt.Fprintf(rw, myjson3.Name)
	var myJsonresp2 MyJsonRes
	myJsonresp2.Salutation = "PHello " + myjson3.Name + " !"
	b2, err := json.Marshal(myJsonresp2)
	if err != nil {
	}
	fmt.Fprintf(rw, string(b2))
}

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	mux.POST("/Hello", PostHello)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
