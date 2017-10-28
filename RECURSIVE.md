# Recursive Testing

## Steps

### Step 0
Start the server:
```
bin/gremlin-server.sh conf/gremlin-server-modern.yaml
```
It listens on port 8182

Start the concole (if needed):
```
bin/gremlin.sh
gremlin> :remote connect tinkerpop.server conf/remote.yaml
```
The above loads in the modern graph data sample.

To open the in memory graph db:
```
graph = TinkerGraph.open()
```

### Step 1 Load the vertices

Here is the data:
```
sqlite> .headers on
sqlite> .mode column
sqlite> select * from part;
id          name
----------  ----------
1           P100
2           P200
3           P210
4           P220
5           P300
6           P310
7           P320
8           P330
9           P340
10          P350
11          P500
12          P600
13          P700
```

graph = TinkerGraph.open()
P100 = graph.addVertex(T.id, 1, T.label, "part", "name", "P100")
P200 = graph.addVertex(T.id, 2, T.label, "part", "name", "P200")
P210 = graph.addVertex(T.id, 3, T.label, "part", "name", "P210")
P220 = graph.addVertex(T.id, 4, T.label, "part", "name", "P220")
P300 = graph.addVertex(T.id, 5, T.label, "part", "name", "P300")
P310 = graph.addVertex(T.id, 6, T.label, "part", "name", "P310")
P320 = graph.addVertex(T.id, 7, T.label, "part", "name", "P320")
P330 = graph.addVertex(T.id, 8, T.label, "part", "name", "P330")
P340 = graph.addVertex(T.id, 9, T.label, "part", "name", "P340")
P350 = graph.addVertex(T.id, 10, T.label, "part", "name", "P350")
P500 = graph.addVertex(T.id, 11, T.label, "part", "name", "P500")
P600 = graph.addVertex(T.id, 12, T.label, "part", "name", "P600")
P700 = graph.addVertex(T.id, 13, T.label, "part", "name", "P700")





### Step 2 Load the edges

Overview:
- P100 is connected P200, P210, and P220
- P200 is connected P300 and P310
- P210 is connected to P320 and P330
- P220 is connected to P340 and P350
- P500 is connected to P600 is connected to P700 is connected to P500



P100.addEdge("has_part", P200, id, 1001)
P100.addEdge("has_part", P210, id, 1002)
P100.addEdge("has_part", P220, id, 1003)
P200.addEdge("has_part", P300, id, 1004)
P200.addEdge("has_part", P310, id, 1005)
P210.addEdge("has_part", P320, id, 1006)
P210.addEdge("has_part", P330, id, 1007)
P220.addEdge("has_part", P340, id, 1008)
P220.addEdge("has_part", P350, id, 1009)
P500.addEdge("has_part", P600, id, 1010)
P600.addEdge("has_part", P700, id, 1011)
P700.addEdge("has_part", P500, id, 1012)



## Traversals
```
g = graph.traversal()
==>graphtraversalsource[tinkergraph[vertices:13 edges:9], standard]

gremlin> g.V(1).valueMap()
==>[name:[P100]]

gremlin> g.V(1).repeat(out()).emit().path().by('name')
==>[P100, P200]
==>[P100, P210]
==>[P100, P220]
==>[P100, P200, P300]
==>[P100, P200, P310]
==>[P100, P210, P320]
==>[P100, P210, P330]
==>[P100, P220, P340]
==>[P100, P220, P350]

gremlin> g.V().has('name','P500')
==>v[11]

g.V().has('name','P500').repeat(out()).emit().path().by('name')
==> goes into loop, but stops with a "..." as last line. Apparently it knows that it is in a loop?

-- use simplePath()
g.V().has('name','P500').repeat(out()).emit().simplePath().path().by('name')
==> goes into a loop and hangs; shows two rows:
==>[P500, P600]
==>[P500, P600, P700]

... and then hangs ...

-- add the "times()" modulator(?)
g.V().has('name','P500').repeat(out()).times(1).emit().simplePath().path().by('name')
--> returns:
==>[P500, P600]

g.V().has('name','P500').repeat(out()).times(2).emit().simplePath().path().by('name')
==>[P500, P600]
==>[P500, P600, P700]

g.V().has('name','P500').repeat(out()).times(3).emit().simplePath().path().by('name')
==>[P500, P600]
==>[P500, P600, P700]

-- this will order the rows properly
g.V(1).repeat(out()).emit().path().by('name').order().by {String.join("", it.objects())}
```
