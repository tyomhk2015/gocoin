# gocoin
Building blockchain and crypto currency with Go lang.

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
* Created a http server with Go, and a blockchain web application. <a href="https://github.com/tyomhk2015/gocoin/tree/main/goserver">Source Code</a>

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

* Finished creating a web application w/ blockchain feature. <a href="https://github.com/tyomhk2015/gocoin/tree/main/goserver">Source Code</a>

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