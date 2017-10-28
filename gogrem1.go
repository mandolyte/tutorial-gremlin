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

	data, err := gremlin.Query(`g.V().values('name')`).Exec()
	if err != nil {
		log.Fatalf("Query Exec Error:%v\n", err)
	}

	fmt.Println(string(data))
}

/*
$ go run gogrem1.go
["marko","vadas","lop","josh","ripple","peter"]

*/
