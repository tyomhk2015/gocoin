package main

import (
	"fmt"
	"reflect"
)

func main() {
	// sing("A", "B", "C")
	randomNum := -127
	// randomLetter := 'a'
	// fmt.Printf("Left: %b, Right: %b", randomNum, randomLetter)
	fmt.Println(fmt.Sprintf("%x", randomNum), ":", reflect.TypeOf(fmt.Sprintf("%x", randomNum)))

	sheep := talent{"Watame", "JP"}
	sheep.stream(11)
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
