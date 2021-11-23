package main

import (
	"github.com/tyomhk2015/gocoin/restapi/explorer"
	"github.com/tyomhk2015/gocoin/restapi/rest"
)

func main() {
	go explorer.Start(8889)
	rest.Start(8888)
}
