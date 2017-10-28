# Go Gremlin testing

## Go-Gremlin install
go get github.com/go-gremlin/gremlin

In console, this remote execution:
```
gremlin> :remote connect tinkerpop.server conf/remote.yaml
==>Connected - localhost/127.0.0.1:8182
gremlin> :> g.V().values('name')
==>marko
==>vadas
==>lop
==>josh
==>ripple
==>peter
gremlin>
```

Some sample code executions:
```
$ go run gogrem1.go
["marko","vadas","lop","josh","ripple","peter"]
$ go run gogrem2.go
["lop"]
$ go run gogrem3.go
[{"name":["marko"],"age":[29]}]
$
```

Some nice console commands:
```
gremlin> :> g.V().has('name','marko2').valueMap()
==>{name=[marko2], age=[92]}
gremlin> :> g.V().has('label','part').valueMap()
gremlin> :> g.V().has(label,'part').valueMap()
==>{name=[P100]}
gremlin> :> g.V().has(label,'part').drop()
gremlin> :> g.V().has(label,'part').valueMap()
gremlin> :> g.V().has(label,'part').count()
==>13
gremlin> :> g.V().has(label,'part').count()
==>13
gremlin> :> g.V().has(label,'part').drop()
gremlin> :> g.V().has(label,'part').count()
==>0

```
Add vertices for BOM
```
$ go run gogrem-addBomVertices.go
[{"id":48,"label":"part","type":"vertex","properties":{"name":[{"id":49,"value":"P100"}]}}]
[{"id":50,"label":"part","type":"vertex","properties":{"name":[{"id":51,"value":"P200"}]}}]
[{"id":52,"label":"part","type":"vertex","properties":{"name":[{"id":53,"value":"P210"}]}}]
[{"id":54,"label":"part","type":"vertex","properties":{"name":[{"id":55,"value":"P220"}]}}]
[{"id":56,"label":"part","type":"vertex","properties":{"name":[{"id":57,"value":"P300"}]}}]
[{"id":58,"label":"part","type":"vertex","properties":{"name":[{"id":59,"value":"P310"}]}}]
[{"id":60,"label":"part","type":"vertex","properties":{"name":[{"id":61,"value":"P320"}]}}]
[{"id":62,"label":"part","type":"vertex","properties":{"name":[{"id":63,"value":"P330"}]}}]
[{"id":64,"label":"part","type":"vertex","properties":{"name":[{"id":65,"value":"P340"}]}}]
[{"id":66,"label":"part","type":"vertex","properties":{"name":[{"id":67,"value":"P350"}]}}]
[{"id":68,"label":"part","type":"vertex","properties":{"name":[{"id":69,"value":"P500"}]}}]
[{"id":70,"label":"part","type":"vertex","properties":{"name":[{"id":71,"value":"P600"}]}}]
[{"id":72,"label":"part","type":"vertex","properties":{"name":[{"id":73,"value":"P700"}]}}]
$
```
