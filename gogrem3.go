package main

import (
	"fmt"
	"log"

	"github.com/go-gremlin/gremlin"
)

func main() {

	if err := gremlin.NewCluster("ws://localhost:8182"); err != nil {
		log.Fatalf("NewCluster Error:%v\n", err)
	}

	data, err := gremlin.Query(`g.V().has('name',x)valueMap()`).Bindings(gremlin.Bind{"x": "marko"}).Exec()
	if err != nil {
		log.Fatalf("Query Exec Error:%v\n", err)
	}

	fmt.Println(string(data))
}

/*
$ go run gogrem3.go
[{"name":["marko"],"age":[29]}]
$
*/
