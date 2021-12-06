package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/tyomhk2015/gocoin/blockchain"
	"github.com/tyomhk2015/gocoin/utils"
)

func Start(portNum int) {
	port = fmt.Sprintf(":%d", portNum)
	// muxRouter = http.NewServeMux()
	muxRouter = mux.NewRouter()
	// middleware
	muxRouter.Use(jsonContentTypeMiddleware)

	prepareHandlers()
	createServer()
}

var (
	port      string
	muxRouter *mux.Router
)

type url string

type balanceResponse struct {
	Address string `json:"address"`
	Balance int    `json:"balance"`
}

type addingTxPayload struct {
	Receiver string `json:"receiver"`
	Amount   int    `json:"amount"`
}

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

func createServer() {
	fmt.Printf("Listening to localhost%s", port)
	log.Fatal(http.ListenAndServe(port, muxRouter))
}

func prepareHandlers() {
	muxRouter.HandleFunc("/", documentation).Methods("GET")                // Mux's feature; pre-defining acceptable methods.
	muxRouter.HandleFunc("/blocks", blocks).Methods("GET", "POST")         // If requests' methods are not specified by mux's Method(),
	muxRouter.HandleFunc("/blocks/{hash:[a-f0-9]+}", block).Methods("GET") // Hash will be the URL parameter, the hexdecimal.
	muxRouter.HandleFunc("/status", status).Methods("GET")
	muxRouter.HandleFunc("/balance/{address}", balance).Methods("GET")  // Show all transaction outputs of a specific address.
	muxRouter.HandleFunc("/mempool", mempool).Methods("GET")            // Show all unconfirmed transactions, mempool.
	muxRouter.HandleFunc("/transactions", transactions).Methods("POST") // Add a transaction to the mempool.
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See documentation.",
		},
		{
			URL:         url("/status"),
			Method:      "GET",
			Description: "See the status of the blockchain.",
		},
		{
			URL:         url("/blocks"),
			Method:      "GET",
			Description: "Show all blocks.",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add a block",
			Payload:     "data: string",
		},
		{
			URL:         url("/blocks/{hash}"),
			Method:      "GET",
			Description: "Show only one block.",
		},
		{
			URL:         url("/balance/{address}"),
			Method:      "GET",
			Description: "Show transaction outputs for an address.",
		},
		{
			URL:         url("/mempool"),
			Method:      "GET",
			Description: "Show all unconfirmed transactions.",
		},
		{
			URL:         url("/transactions"),
			Method:      "POST",
			Description: "Add a transaction to the mempool.",
		},
	}

	// Change the returned data as JSON, not text/plain.
	// Some precautions for writing http response writer;
	// https://ryanc118.medium.com/dont-make-this-mistake-with-go-http-servers-bd313baee41
	// rw.Header().Add("Content-Type", "application/json")

	// Marshal Way
	// jsonData := convertJSON(data)

	// NewEncoder Way
	err := json.NewEncoder(rw).Encode(data)
	utils.HandleErr(err)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(rw).Encode(blockchain.Blockchain().Blocks())
	case "POST":
		blockchain.Blockchain().AddBlock()
		rw.WriteHeader(http.StatusCreated)
	}
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

func block(rw http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	hash := param["hash"]
	block, err := blockchain.FindBlock(hash)
	jsonEncoder := json.NewEncoder(rw)
	if err == nil {
		// Convert go struct to JSON.
		jsonEncoder.Encode(block)
	} else {
		jsonEncoder.Encode(errorResponse{fmt.Sprint(err)})
	}
}

func convertJSON(data []URLDescription) []byte {
	// Marshal: Converting a struct into JSON.
	// Unmarshal: Converting JSON to a data, that is comprehensible to desired programming language, Go's struct.

	// Marshal way
	b, err := json.Marshal(data)
	utils.HandleErr(err)
	return b
}

// Middleware:
// A function that will be called first.
// A function called before the final destination.
// A function that is in the middle of ordinary flow.
// A function that intercepts and modify the data before sending them to the next destination.
func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	// HandlerFunc is type, the arguments are for constructing the HandlerFunc type.

	// Apdapter:
	// Recieves arguments and create a function that the Handler(return interface) requires.
	// Saves time for not structing and creating type from the very bottom.
	// Just put some ingredients you need to the adapter(type), and the adapter will
	// created the necessary func()s for you.
	// Inside the HandlerFunc, the necessary items will be created to satisfy conditions for returning 'http.Handler'.
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}
func status(rw http.ResponseWriter, r *http.Request) {
	// Show the status of the blockchain.
	json.NewEncoder(rw).Encode(blockchain.Blockchain())
}

// Show all transaction outputs of a specific address.
func balance(rw http.ResponseWriter, r *http.Request) {
	address := mux.Vars(r)["address"]
	totalFlag := r.URL.Query().Get("total")
	switch totalFlag {
	case "true":
		balance := blockchain.Blockchain().BalanceByAddress(address)
		utils.HandleErr(json.NewEncoder(rw).Encode(balanceResponse{address, balance}))
	default:
		utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.Blockchain().TxOutsByAddress(address)))
	}
}

func mempool(rw http.ResponseWriter, r *http.Request) {
	// Show transactions in the mempool.
	utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.Mempool.Txs))
}

func transactions(rw http.ResponseWriter, r *http.Request) {
	// Add a transaction in the mempool.
	var payload addingTxPayload
	utils.HandleErr(json.NewDecoder(r.Body).Decode(&payload))
	err := blockchain.Mempool.AddTx(payload.Receiver, payload.Amount)
	if err != nil {
		json.NewEncoder(rw).Encode(errorResponse{"Not enough coins."})
	}
	rw.WriteHeader(http.StatusCreated)
}
