package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/tyomhk2015/gocoin/blockchain/blockchain"
)

func Start(portNum int) {
	port = fmt.Sprintf(":%d", portNum)
	loadTemplates()
	prepareHandlers()
	createServer()
}

const templateLocation string = "templates/"

var (
	port      string
	templates *template.Template
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func createServer() {
	// Make a server with go lang w/ standard library, net/http.
	// Detect any errors while server is running.
	log.Fatal(http.ListenAndServe(port, nil))
}

func prepareHandlers() {
	// Add routes
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/add", handleAdd)
}

func handleAdd(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		data := homeData{"Add your block!", nil}
		templates.ExecuteTemplate(rw, "add", data)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusMovedPermanently)
	}
}

func handleHome(rw http.ResponseWriter, r *http.Request) {
	// http.ResponseWriter: The writer of data to send to the users.
	data := homeData{"This is running on the Go server.", blockchain.GetBlockchain().ShowAllBlocks()}
	templates.ExecuteTemplate(rw, "home", data)
}

func loadTemplates() {
	// ** all folder symbol is not supported, yet.
	templates = template.Must(template.ParseGlob(templateLocation + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateLocation + "partials/*.gohtml"))
}
