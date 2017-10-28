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

	data, err := gremlin.Query(
		`graph.addVertex(label, "person", "name", "marko2", "age", 92)`).Exec()
	if err != nil {
		log.Fatalf("Query Exec Error:%v\n", err)
	}

	fmt.Println(string(data))
}

/*
This:
`graph.addVertex(id, 111, label, "person", "name", "marko2", "age", 92)`).Exec()
returns:
[{"id":111,"label":"person","type":"vertex","properties":{"name":[{"id":13,"value":"marko2"}],"age":[{"id":14,"value":92}]}}]

If I omit the id like this:
`graph.addVertex(label, "person", "name", "marko2", "age", 92)`).Exec()
then one is genterate. This returns:
[{"id":17,"label":"person","type":"vertex","properties":{"name":[{"id":18,"value":"marko2"}],"age":[{"id":19,"value":92}]}}]

*/
