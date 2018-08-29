package main

import (
	"flag"
	"strconv"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
	"encoding/json"
	"github.com/bingoohuang/go-utils"
)

var (
	port string
)

func init() {
	portArg := flag.Int("port", 8884, "Port to serve.")

	flag.Parse()

	port = strconv.Itoa(*portArg)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/token/{tokenId}", getToken)
	http.Handle("/", r)

	fmt.Println("start to listen at ", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func getToken(w http.ResponseWriter, r *http.Request) {
	go_utils.HeadContentTypeJson(w)
	vars := mux.Vars(r)
	tokenId := vars["tokenId"]

	json.NewEncoder(w).Encode(tokenId)
}

type ExpireableToken struct {
	Value   string `json:"value"`
	Expired uint64 `json:"expiredMillis"`
}

type TokenRefreshable interface {
	refresh(tokenId string) ExpireableToken
}
