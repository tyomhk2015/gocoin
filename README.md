# gocoin
Building crypto currency with Go lang.

<a href="#user-content-day1">Day 1</a>　2021/11/15
* Warming up: reviewed the fundamentals of Go.


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

* <a href="https://pkg.go.dev/fmt">fmt</a>: Format a string/letter in either digit, binary, or as string. Useful for change the data into a different form.
<br>
I think this may become handy when serialization is needed. 🤔
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

* Methods / Receiver function : Adding a functon to struct.
<br>
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