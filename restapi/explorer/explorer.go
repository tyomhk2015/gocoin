package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/tyomhk2015/gocoin/blockchain"
)

func Start(portNum int) {
	port = fmt.Sprintf(":%d", portNum)
	handler = http.NewServeMux()
	loadTemplates()
	prepareHandlers()
	createServer()
}

const templateLocation string = "templates/"

var (
	port      string
	templates *template.Template
	handler   *http.ServeMux
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func createServer() {
	fmt.Printf("Listening to localhost%s", port)
	// Make a server with go lang w/ standard library, net/http.
	// Detect any errors while server is running.
	log.Fatal(http.ListenAndServe(port, handler))
}

func prepareHandlers() {
	// Add routes
	handler.HandleFunc("/", handleHome)
	handler.HandleFunc("/add", handleAdd)
}

func handleAdd(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		data := homeData{"Add your block!", nil}
		templates.ExecuteTemplate(rw, "add", data)
	case "POST":
		blockchain.Blockchain().AddBlock()
		http.Redirect(rw, r, "/", http.StatusMovedPermanently)
	}
}

func handleHome(rw http.ResponseWriter, r *http.Request) {
	// http.ResponseWriter: The writer of data to send to the users.
	data := homeData{"This is running on the Go server.", blockchain.Blocks(blockchain.Blockchain())}
	templates.ExecuteTemplate(rw, "home", data)
}

func loadTemplates() {
	// ** all folder symbol is not supported, yet.
	templates = template.Must(template.ParseGlob(templateLocation + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateLocation + "partials/*.gohtml"))
}
