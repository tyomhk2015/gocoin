# gocoin
Building crypto currency with Go lang.

<a href="#user-content-day1">Day 1</a>　2021/11/15
* Warming up: reviewed the fundamentals of Go.

<a href="#user-content-day2">Day 2</a>　2021/11/16
* Warming up: reviewed the fundamentals of Go, structs & pointer.

<a href="#user-content-day3">Day 3</a> ~ Day 6　2021/11/17 ~ 2021/11/20
* Taking <a href="https://github.com/tyomhk2015/gocoin/blob/main/terms/terms.md">notes</a> about concepts related to blockchain. 

<a href="#user-content-day7">Day 7</a>　2021/11/21
* My first blockchain. <a href="https://github.com/tyomhk2015/gocoin/tree/main/blockchain">Source Code</a>

<a href="#user-content-day8">Day 8</a>　2021/11/22
* Created a http server with Go, using Go's standard library. 

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

* The `lowercase` variables also cannot be referred by parsing files, such as `html/template`.<br>
In other words, referring a title of the root page from Go with lowercase variable will give you nothing in return, because the lowercase cannot be accessed from outside of the file where the variable has been declared.

* Used html/template, standard library, for inserting Go data into a HTML template.<br>`Must()` is used for checking errors during the parsing of files w/ the template library.

* <a href="https://andybrewer.github.io/mvp/">MVPCSS</a>, an external CSS for focusing more on the logic of the project.<br>No CSS class names are required.

* When variable is conveyed from Go file to template, the first element that wants the variables must be written as `{{.Variable}}`. If you want to send the received variable to the children elements, it should be written as `{{.}}`. Commenting out is also possible.

* When taking an action or accessing specific route, the one request can have different method.

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
