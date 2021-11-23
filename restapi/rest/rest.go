package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tyomhk2015/gocoin/blockchain/blockchain"
	"github.com/tyomhk2015/gocoin/utils"
)

func Start(portNum int) {
	port = fmt.Sprintf(":%d", portNum)
	handler = http.NewServeMux()
	prepareHandlers()
	createServer()
}

var (
	port    string
	handler *http.ServeMux
)

type url string

// Interface implementation
func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type URLDescription struct {
	// Strcut field tags → Changes the way of showing the name of the keys.
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

// A struct for recieving and converting JSON from a request.
type BlockBody struct {
	// Must match the JSON key name with struct field name.
	// Otherwise an empty string will be returned.
	Message string
	// Signal  string
}

func createServer() {
	fmt.Printf("Listening to localhost%s", port)
	log.Fatal(http.ListenAndServe(port, handler))
}

func prepareHandlers() {
	handler.HandleFunc("/", documentation)
	handler.HandleFunc("/blocks", blocks)
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See documentation.",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "See documentation.",
			Payload:     "data: string",
		},
		{
			URL:         url("/blocks/{id}"),
			Method:      "GET",
			Description: "See documentation.",
			Payload:     "data: string",
		},
	}

	// Change the returned data as JSON, not text/plain.
	// Some precautions for writing http response writer;
	// https://ryanc118.medium.com/dont-make-this-mistake-with-go-http-servers-bd313baee41
	rw.Header().Add("Content-Type", "application/json")

	// Marshal Way
	// jsonData := convertJSON(data)

	// NewEncoder Way
	err := json.NewEncoder(rw).Encode(data)
	utils.HandleErr(err)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Convert struct to JSON
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().ShowAllBlocks())
	case "POST":
		// Convert JSON to struct
		var addedBlockBody BlockBody
		err := json.NewDecoder(r.Body).Decode(&addedBlockBody)
		utils.HandleErr(err)
		blockchain.GetBlockchain().AddBlock(addedBlockBody.Message)
		fmt.Println(addedBlockBody)
		rw.WriteHeader(http.StatusCreated)
	}
}

func convertJSON(data []URLDescription) []byte {
	// Marshal: Converting a struct into JSON.
	// Unmarchal: Converting JSON to a data,
	// that is comprehensible to desired programming language, Go's struct.

	// Marshal way
	b, err := json.Marshal(data)
	utils.HandleErr(err)
	return b
}
