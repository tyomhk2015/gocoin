# GoCoin
Building blockchain and crypto currency with Go lang.

Why am I doing this?
To get used to Go lang and blockchain.

## Summary 🔎

<a href="#user-content-day1">Day 1</a>　2021/11/15
* Warming up: reviewed the fundamentals of Go.

<a href="#user-content-day2">Day 2</a>　2021/11/16
* Warming up: reviewed the fundamentals of Go, structs & pointer.

<a href="#user-content-day3">Day 3</a> ~ Day 6　2021/11/17 ~ 2021/11/20
* Taking <a href="https://github.com/tyomhk2015/gocoin/blob/main/terms/terms.md">notes</a> about concepts related to blockchain. 

<a href="#user-content-day7">Day 7</a>　2021/11/21
* My first blockchain. <a href="https://github.com/tyomhk2015/gocoin/tree/main/blockchain">Source Code</a>

<a href="#user-content-day8">Day 8</a>　2021/11/22
* 🌟 Created a http server with Go, and a blockchain web application. <a href="https://github.com/tyomhk2015/gocoin/tree/main/goserver">Source Code</a>

<a href="#user-content-day9">Day 9</a>　2021/11/23
* Started creating REST API of my blockchain.
* Marshal & Unmarshal
* Interface
* Struct field tags
* Encode & Decode JSON
* Added 'REST client' extension for speeding up REST tests.
* Mux & NewServeMux

<a href="#user-content-day10">Day 10</a>　2021/11/25
* Studies some theories about blockchain.
<a href="https://github.com/tyomhk2015/gocoin/blob/main/terms/terms.md">Notes 📝</a> 

<a href="#user-content-day11">Day 11</a>　2021/11/26
* Continued creating REST API of my blockchain.
* gorilla/mux
* Troubleshooting
* Middleware
* Adapter
* 🌟 Done with the REST API project. (<a href="https://github.com/tyomhk2015/gocoin/tree/main/restapi">Link</a>)

<a href="#user-content-day12">Day 12</a>　2021/11/27
* Started building a CLI application.
* Pros of CLI application: Automate some simple tasks, like downloading, moving files to other directories, deleting files etc, via command line.
* flag (standard library)
* 🌟 Done with the CLI application. (<a href="https://github.com/tyomhk2015/gocoin/tree/main/cli">Link</a>)
* Persistence, a.k.a database.

<a href="#user-content-day13">Day 13</a>　2021/11/28
* Started adding blocks with the bolt DB.

<a href="#user-content-day14">Day 14</a>　2021/11/29
* Added <a href="https://github.com/ShoshinNikita/boltBrowser">boltBrowser</a> to look in the the database created by <a href="https://github.com/boltdb/bolt">bolt</a>.

<a href="#user-content-day15">Day 15</a>　2021/11/30
* Implemented a feature of retrieving the a block from bolt DB using the block's hash.

<a href="#user-content-day16">Day 16</a>　2021/12/2
* Implemented a feature of retrieving the all blocks from bolt DB using the recent block's hash.
* 🌟 Creating a blockchain with DB is done.
* Started creating mining or proof of work, PoW.

<a href="#user-content-day17">Day 17</a>　2021/12/4
* Implemented dynamic difficulty for adding blocks to the blockchain.
* 🌟 Finished implementing PoW/Mining.
* Started implementing 'transactions'.

<a href="#user-content-day18">Day 18</a>　2021/12/5
* Continued implementing coinbase transactions.
* Troubleshooting

<a href="#user-content-day19">Day 19</a>　2021/12/6
* Implemented mempool and confirming transactions in the mempool

<a href="#user-content-day20">Day 20</a>　2021/12/7
* Implementing 'unspent' transactions.

<a href="#user-content-day21">Day 21</a>　2021/12/9
* Collaborating 'unspent' transactions and mempool.
* Label
* Deadlock
* 🌟 Finished implementing transaction feature, including 'unspent' transactions, mempool, and preventing deadlock.
* Started working on wallets.

<a href="#user-content-day22">Day 22</a>　2021/12/10
* Continued working on wallets.

<a href="#user-content-day23">Day 23</a>　2021/12/13
* Continued working on wallets.

<hr>

#### Resource 📖

❓ <a href="javascript:void(0);" target="_blank" rel="noopener">What is 'crypto currency' and why did it come to existence?</a>
<br>
❓ <a href="javascript:void(0);" target="_blank" rel="noopener">What is 'block chain'?</a>

## Notes 📝

### **<a href="javascript:void(0);" id="day1">Day 1</a>** ☀️
2021/11/15

#### 💡 **Basics**
* Data types like int, int8, int16, int32, the default bit will be 32bit unless they are explicitly specified.

* Function's arguments have same data type, they can be shorthanded.
<pre>
func sing(shiny string, smile string, story string) {}
▼
func sing(shiny, smile, story string) {}
</pre>

* <a href="https://pkg.go.dev/fmt">fmt</a>: Format a string/letter in either digit, binary, or as string. Useful for change the data into a different form.<br>I think this may become handy when serialization is needed. 🤔
<pre>
randomNum := -127
randomLetter := 'a'
fmt.Printf("Left: %b, Right: %b", randomNum, randomLetter)
fmt.Println(
  fmt.Sprintf("%x", randomNum),
  ":",
  reflect.TypeOf(fmt.Sprintf("%x", randomNum))
)

// Returns
// Left: -1111111, Right: 1100001
// -7f : string
</pre>

* Methods / Receiver function : Adding a functon to struct, similar to getter/setter in Java.<br>
(Condition) The method or func must have struct written right after the keyword 'func'.
<pre>
type talent struct {
  name     string
  language string
}

func (t talent) stream(time int) {
  fmt.Printf("%s's stream starts at %dpm, %s.", t.name, time, t.language)
}

func main() {
  sheep := talent{"Watame", "JP"}
  sheep.stream(11)
}

// Return: Watame's stream starts at 11 pm, JP.
</pre>

<hr>

### **<a href="javascript:void(0);" id="day2">Day 2</a>** ☀️
2021/11/16

#### 💡 **Structs & Pointer**

* Uppercase = Exportable
* Lowercase = Not exportable, private.

🖉 The GOPATH setting that I wrote during the 'go_basics' can be ignored. go.mod file handles the enviorment path of go. 

<hr>


### **<a href="javascript:void(0);" id="day3">Day 3</a>** ☀️
2021/11/17

* Blockchain Notes

- Smart Contract
- Main chain & side chain
- NFT (Non-fungible Token)
- Decentralization
- Consensus mechanism
- Proof of work
- Proof of stake
- ...

All the notes about terms will be stored in my <a href="https://github.com/tyomhk2015/gocoin/blob/main/terms/terms.md">terms folder</a>. 📂

### **<a href="javascript:void(0);" id="day7">Day 7</a>** ☀️
2021/11/21

* Created my first blockchain. <a href="https://github.com/tyomhk2015/gocoin/tree/main/blockchain">Source Code</a>

### **<a href="javascript:void(0);" id="day8">Day 8</a>** ☀️
2021/11/22

* Created http server-side with Go, using Go's standard library.

💡 The `lowercase` variables also cannot be referred by parsing files, such as `html/template`.<br>
In other words, referring a title of the root page from Go with lowercase variable will give you nothing in return, because the lowercase cannot be accessed from outside of the file where the variable has been declared.

* Used html/template, standard library, for inserting Go data into a HTML template.<br>`Must()` is used for checking errors during the parsing of files w/ the template library.

💡 <a href="https://andybrewer.github.io/mvp/">MVPCSS</a>, an external CSS for focusing more on the logic of the project.<br>No CSS class names are required.

💡 When variable is conveyed from Go file to template, the first element that wants the variables must be written as `{{.Variable}}`. If you want to send the received variable to the children elements, it should be written as `{{.}}`. Commenting out is also possible.

💡 When taking an action or accessing specific route, the one request can have different method.

<pre>
http.HandleFunc("/add", handleAdd)

func handleAdd(rw http.ResponseWriter, r *http.Request) {
  switch r.Method {
  case "GET":
    ...
  case "POST":
    ...
  }
}
</pre>

🌟Finished creating a web application w/ blockchain feature. <a href="https://github.com/tyomhk2015/gocoin/tree/main/goserver">Source Code</a>

![gocoin](https://user-images.githubusercontent.com/35278730/142892010-5b8b20a4-629f-4969-a683-52648a1a29d2.gif)

### **<a href="javascript:void(0);" id="day9">Day 9</a>** ☀️
2021/11/23

<a href="https://pkg.go.dev/github.com/tendermint/tendermint/libs/json">Resource</a> 📖

💡 Marshal & Unmarshal
<br>`Marshal`: Convert Go struct to JSON.
<br>`Unmarshal`: Convert JSON to Go struct.
<pre>
b, err := json.Marshal(data)
</pre>

💡 Struct field tags
<br>For changing the name of the `key`s of struct when converted to JSON.
<br>The `json:"xxxx"` part is the struct field tags.
<pre>
type URLDescription struct {
  URL         string `json:"url"`
  Method      string `json:"method"`
  Description string `json:"description"`
  Payload     string `json:"payload,omitempty"`
}
</pre>
<pre>
Usage example: 
Useful for changing the uppercase to lowercase,
because, most of time, JSON use lowercases.
</pre>

💡 Implementation of other interface, similar feature of `implements` in Java.
<br>More then one interfaces can be implemented.
<br>There is no need to explicitly write that you are implementing an interface.
<br>Go does this for you, implicitly.

<pre>
type Stringer interface {
  String() string // ①
}
</pre>
<pre>
By appending '①' part, after the receiver function,
you can change how the designated type/struct looks
when called by, for example, fmt.Println().
However, the signatures must be kept as it is.
Otherwise the implementation won't be in effect.
</pre>
<pre>
func (c customType) String() string {
  return "This is implemented w/ Stringer()"
}

fmt.Println(customType{})
// Return: This is implemented w/ Stringer()
</pre>
<pre>
Usage example:
During the conversion from struct to JSON, you can intercept the struct fields,
and customize them in the way you want to show.
E.g: Instead of using 'URL' field name, changing it to 'domain' and show this on JSON format.
</pre>

⚠️ Troubleshooting

Problem:
<br>
Got the following message after I was testing REST POST request with 'REST client' extension.
<pre>
invalid character '}' looking for beginning of object key string
</pre>

Solution: ✔️
<br>
Fixed the issue by validating JSON format. There was a comma at the end of the value.
<pre>
// REST client; restapi.http
{
  "message": "TOKINO SORA",  <<< This one comma was the cause of the trouble.
}

↓↓

{
  "message": "TOKINO SORA"
}
</pre>

💡 Mux & NewServeMux 

* Mux: Multiplexer
<br>A part in server side that takes care of `clients' requests` by looking at URLs and initiating correspondant handlers.
<br>The detail of the 2nd argument of http.ListenAndServe() was the multiplexer, `mux`. If this is not specified, the `http` server will use default mux.

<img src="https://miro.medium.com/max/1400/0*XSB0-ot-IX4dh9nF" alt="https://medium.com/golang-with-azure/getting-started-making-it-a-golang-web-application-e2579636f50a" />

⚠️ Troubleshooting

Problem:
<br>
Got the following message when I created NewServeMux for explorer.go and rest.go
<pre>
runtime error: invalid memory address or nil pointer dereference
</pre>

Solution: ✔️
<br>
Fixed the issue by creating NewServeMux prior to adding http.HandleFunc()s.
<br>
The cause of the issue was that I was registering HandlerFunc()s earlier than creating NewServMux. This resulted that HandleFunc()s were referring to `nil`.
<pre>
nil.HandleFunc() // X
</pre>

### **<a href="javascript:void(0);" id="day10">Day 10</a>** ☀️
2021/11/25

* Studied some blockchain related theories.
<a href="https://github.com/tyomhk2015/gocoin/blob/main/terms/terms.md">Notes </a>

### **<a href="javascript:void(0);" id="day11">Day 11</a>** ☀️
2021/11/26

* Chose <a href="https://github.com/gorilla/mux">Gorilla/mux</a> for managing routes, taking some parameter from ULRs, of this REST API project.
<br>
It uses standard library of Go, and I believe it is stable, which means there is likely less chance of maintaining the project in the future. This may save some time.

⚠️ Troubleshooting

Problem:
<br>
<a href="https://github.com/gorilla/mux">Gorilla/Mux</a> was properly installed, but VSCode did not import it automatically.

Solution: ✔️
<br>
Fixed the issue by changing the value of `GO111MODULE`, using the command, `go env -w GO111MODULE=on`.
<br>

* While debugging, I had a chance to read <a href="https://maelvls.dev/go111module-everywhere/">an article</a> about 'GO111MODULE', did not read it all yet. However, I believe this may become handy when having trouble with directories.

<br>

💡 Middleware
* A function that will be called first.
* A function called before the final destination.
* A function that is in the middle of ordinary flow.
* A function that intercepts and modify the data before sending them to the next destination.

<pre>
// Before proceeding to adding/getting the blocks, w/ the middleware, json will be added to the http-header.
func jsonContentTypeMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
    rw.Header().Add("Content-Type", "application/json")
    next.ServeHTTP(rw, r)
  })
}
</pre>

💡 Adapter
* Recieves arguments and create a function that the Handler(return interface) requires.
* Saves time for not structing and creating type from the very bottom.
* Just put some ingredients you need to the adapter(type), and the adapter will created the necessary func()s for you.
<pre>
func jsonContentTypeMiddleware(next http.Handler) http.Handler {
  // HandlerFunc is type, the arguments are for constructing the HandlerFunc type.
  // Inside the HandlerFunc, the necessary items will be created to satisfy conditions for returning 'http.Handler'.
  return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
    rw.Header().Add("Content-Type", "application/json")
    next.ServeHTTP(rw, r)
  })
}
</pre>

🌟Done with the REST API project. (<a href="https://github.com/tyomhk2015/gocoin/tree/main/restapi">Link</a>)
![restapi](https://user-images.githubusercontent.com/35278730/143592552-96ebcf3c-d3c3-4857-a4ba-3df29c85864c.gif)


### **<a href="javascript:void(0);" id="day12">Day 12</a>** ☀️
2021/11/27

* Started building a CLI application.

💡 Flag
* A standard go package that enables to develop CLI application.
* For building a large CLI application, using <a href="https://github.com/spf13/cobra">Cobra</a> is a good option. <a href="https://github.com/spf13/cobra/blob/master/projects_using_cobra.md">Who uses cobra?</a> 🐍

* I guess this is how homebrew, npm, yarn, or rbenv etc works.

* Specifing array range.
<pre>
var arr = [5]int{1, 2, 3, 4, 5}
fmt.Println(arr[2:])

// Result: [3 4 5]
// Syntax: arr[beginning index : ending index]
// If the ending index is unspecified, the ending index wiill be the last index of the array.
</pre>

* flag.NewFlagSet() : Useful for commands that have a lot of flags.
* usage: A message that is displayed if some conditions were not met for the command line.
* When assigning a value to a flag, a whitespace or an equal sign works same.
<pre>
go run main.go rest -port 9999
go run main.go rest -port=9999
</pre>

🌟Done with the CLI application. (<a href="https://github.com/tyomhk2015/gocoin/tree/main/cli">Link</a>)
![cli](https://user-images.githubusercontent.com/35278730/143676543-b667df15-5e73-4bcd-a369-f51b1ee5d0a0.gif)


💡 Persistence / Database: Bolt

* Used <a href="https://github.com/boltdb/bolt">bolt</a> as the database for the blockchain project.
<br><br>
Reasons for choosing bolt
<br>
① Very easy to use, low learning curve.
<br>
② The api is fixed, no furthure changes, and it is stable.
<br>
③ The project with using bolt will less likely have issues with database logic after a year or 
so.
<br>
④ Many people are using the bolt.
<br>
⑤ Serves my purpose and meets my necessary needs.

### **<a href="javascript:void(0);" id="day13">Day 13</a>** ☀️
2021/11/28

* Started adding blocks with the bolt DB.

### **<a href="javascript:void(0);" id="day14">Day 14</a>** ☀️
2021/11/29

* Added <a href="https://github.com/ShoshinNikita/boltBrowser">boltBrowser</a> to look in the the database created by <a href="https://github.com/boltdb/bolt">bolt</a>.

### **<a href="javascript:void(0);" id="day15">Day 15</a>** ☀️
2021/11/30

* Implemented a feature of retrieving the a block from bolt DB using the block's hash.

### **<a href="javascript:void(0);" id="day16">Day 16</a>** ☀️
2021/12/2

* Implemented a feature of retrieving the all blocks from bolt DB using the recent block's hash.
* 🌟 Creating a blockchain with DB is done.
![dbchain](https://user-images.githubusercontent.com/35278730/144434854-6b8c5bbb-2ccb-49ed-8aeb-ab8bc7bcc5a3.gif)
* Started creating mining or proof of work, PoW.
<br>PoW: Something that makes the computers (nodes)to do, but has easy verification.

### **<a href="javascript:void(0);" id="day17">Day 17</a>** ☀️
2021/12/4

* Implemented dynamic difficulty for adding blocks to the blockchain.
<pre>
The difficulty of Bitcoin changes every once when 2016 blocks are added to the blockchain.

Source: https://www.coindesk.com/markets/2021/07/19/bitcoin-network-sees-fourth-straight-downward-difficulty-adjustment/
</pre>
* 🌟 Finished implementing PoW/Mining.
![difficulty4](https://user-images.githubusercontent.com/35278730/144707485-3878ead3-b5c6-4308-8e7c-cb793b32056e.gif)

* Started implementing 'transactions'.

### **<a href="javascript:void(0);" id="day18">Day 18</a>** ☀️
2021/12/5

* Continued implementing coinbase transactions.
* Updated more terms related to blockchain.

⚠️ Troubleshooting

Problem:
<br>
In the `rest.go`, postpended query, `?total=true`, to `/balance/{address}` handlerFunc(). However, could not desired result and got 404 responses.
<pre>
// rest.go
  muxRouter.HandleFunc("/balance/{address}?total=true", balance).Methods("GET")
</pre>

Solution: ✔️
<br>
Fixed the issue by erasing the postpended query. It looks like `http.Request` handles the queries.
<pre>
// rest.go
...
  muxRouter.HandleFunc("/balance/{address}, balance).Methods("GET")
...
  totalFlag := r.URL.Query().Get("total")
...
</pre>

### **<a href="javascript:void(0);" id="day19">Day 19</a>** ☀️
2021/12/6

* Implemented mempool and confirming transactions in the mempool

### **<a href="javascript:void(0);" id="day20">Day 20</a>** ☀️
2021/12/7

* There are two types of transaction outputs, `spent` transaction outputs and `unspent` transaction outputs.
<br> The `unspent` transaction outputs are the ones referred mostly at the future transaction inputs.
<pre>
REMEMBER
- Transaction INPUTs are some transaction OUTPUTs that were created in the past.
- Transaction INPUTs will refer to the past transaction OUTPUTs.
- Once the transaction OUTPUTs are referred at transaction INPUTs, this means they are spent or consumed. 
</pre>

### **<a href="javascript:void(0);" id="day21">Day 21</a>** ☀️
2021/12/9

* Collaborating 'unspent' transactions and mempool.
<br> Making sure than no more transactions to be excecuted when the unconfirmed transactions' input amount exceeds the owner's balance.

💡 Label
<br>
For stopping the nested loops at the desired spot.
<pre>
func isOnMempool(uTxOut *UTxOut) bool {
  exists := false
Outer: // <<< Label
  for _, tx := range Mempool.Txs {
    for _, txIn := range tx.TxIns {
      if txIn.TxID == uTxOut.TxID && txIn.Index == uTxOut.Index {
        exists = true
        break Outer
      }
    }
  }
  return exists
}
</pre>

💡 A tip for deciding function and method.
<br>
If your function `MUTATES` the struct, it should be a method. Otherwise, shouldn't be a method.
`Mutate` means the inner of the struct changes. E.g) The value of field changes, object changes etc.
<pre>
// An example of a method.
func (b *blockchain) txOuts() []*TxOut {
  var txOuts []*TxOut
  blocks := Blocks(b)
  for _, block := range blocks {
    for _, tx := range block.Transactions {
      txOuts = append(txOuts, tx.TxOuts...)
    }
  }  
  return txOuts
}

// An example of function
func BalanceByAddress(address string, b *blockchain) int {
  var totalBalance int
  txOuts := UTxOutsByAddress(address, b)
  for _, txOut := range txOuts {
    totalBalance += txOut.Balance
  }
  return totalBalance
}
</pre>

💡 DeadLock
* Status when the program cannot proceeed.
<br> When the do func() inside the Do(), does not finish and Do() is being called again, it will cause a deadlock. 
<pre>
Do(func{
  ...
  Do(func{
    // Being called, but returning nothing, yet.
  })
  ...
})
</pre>

* 🌟 Finished implementing transaction feature, including 'unspent' transactions, mempool, and preventing deadlock.

* Started working on wallets.
<br>Signatures(private keys and public keys) , persistence of wallets, and signatures and verifications applied to transactions.

### **<a href="javascript:void(0);" id="day22">Day 22</a>** ☀️
2021/12/10

* Continued on Wallets
<br> 1. Hash a message or a block.
<br> 2. Generate a pair of keys, a private Key & a public key.
<br> 3. Sign the hash.
<pre>
  Hashed message + Private Key => Signature
</pre>
<br> Private Key: For signing.
<br> Saved as a file. (An example of a Cold Wallet)
<br> Public Key: Verifiying signatures.
<br> Verification process: 
<pre>
  Hashed message + Signature + Public Key => True / False (Verified or not)
</pre>

### **<a href="javascript:void(0);" id="day23">Day 23</a>** ☀️
2021/12/13

* Continued on Wallets
* Q. What does 'r' and 's' means in ECDSA?
<br>Maybe this must be the start point.
<br>https://crypto.stackexchange.com/questions/50716/what-is-the-relation-between-x-y-and-r-s-in-an-ecdsa-signature