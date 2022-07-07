package main

import(
	"net/http"
	"log"
	"encoding/json"
)

type NFTJSON struct {
	Name string `json:"name"`
	PublicKey string `json:"pk"`
	Id int `json:"id"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1 style='color:steelblue'>Hello</h1>"))
}

func GetNFTJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenOne := NFTJSON {
		Name: "eldm",
		PublicKey: "0x",
		Id: 0,
	}
	json.NewEncoder(w).Encode(tokenOne)
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1 style='color:steelblue'>Hello</h1>"))
	w.Write([]byte("<h2 style='color:mint'>Test</h2>"))
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/hello", Hello) 
	http.HandleFunc("/NFTJSON", GetNFTJSON)
	log.Fatal(http.ListenAndServe(":5100", nil))
}