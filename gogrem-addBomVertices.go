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
	for n := range adds {
		err := execute(adds[n])
		if err != nil {
			log.Fatalf("Query Exec Error:%v\n", err)
		}
	}
}

func execute(info string) error {
	data, err := gremlin.Query(info).Exec()
	fmt.Println(string(data))

	return err
}

/*
graph.addVertex(label, "part", "name", "P100")
*/
var adds []string

func init() {
	adds = append(adds, `graph.addVertex(label, "part", "name", "P100")`)
	adds = append(adds, `graph.addVertex(label, "part", "name", "P200")`)
	adds = append(adds, `graph.addVertex(label, "part", "name", "P210")`)
	adds = append(adds, `graph.addVertex(label, "part", "name", "P220")`)
	adds = append(adds, `graph.addVertex(label, "part", "name", "P300")`)
	adds = append(adds, `graph.addVertex(label, "part", "name", "P310")`)
	adds = append(adds, `graph.addVertex(label, "part", "name", "P320")`)
	adds = append(adds, `graph.addVertex(label, "part", "name", "P330")`)
	adds = append(adds, `graph.addVertex(label, "part", "name", "P340")`)
	adds = append(adds, `graph.addVertex(label, "part", "name", "P350")`)
	adds = append(adds, `graph.addVertex(label, "part", "name", "P500")`)
	adds = append(adds, `graph.addVertex(label, "part", "name", "P600")`)
	adds = append(adds, `graph.addVertex(label, "part", "name", "P700")`)
}
