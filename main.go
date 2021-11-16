package main

import (
	"fmt"

	"github.com/tyomhk2015/gocoin/person"
)

func main() {
	vtuber := person.Person{}
	vtuber.Set("Kiryu", 999)
	fmt.Println("This talent is", vtuber)
}

func sing(shiny, smile, story string) {
	fmt.Println(shiny, smile, story)
}

type talent struct {
	name     string
	language string
}

func (t talent) stream(time int) {
	fmt.Printf("%s's stream starts at %dpm, %s.", t.name, time, t.language)
}
