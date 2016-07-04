# tutorial-gremlin
Notes, etc. on learning TinkerPop &amp; Gremlin

## First steps
After downloading and following instructions on
http://tinkerpop.apache.org/docs/current/tutorials/getting-started/, I discovered that Java is required...

Found this article:
https://www.digitalocean.com/community/tutorials/how-to-install-java-with-apt-get-on-ubuntu-16-04, which shows the following commands for Oracle's JDK:
```
# First, add Oracle's PPA, then update your package repository.
sudo add-apt-repository ppa:webupd8team/java
sudo apt-get update
sudo apt-get install oracle-java8-installer

# where is it?
$ sudo update-alternatives --config java
There is only one alternative in link group java (providing /usr/bin/java): /usr/lib/jvm/java-8-oracle/jre/bin/java
Nothing to configure.

# edit .profile to set JAVA_HOME env var
# then source it and test with
$ echo $JAVA_HOME
/usr/lib/jvm/java-8-oracle
```

## First five minutes
The tutorial advise against getting started with something like Titan or the Gremlin Server databases. There is an in-memory graph DB built into the REPL console.

```
graph = TinkerFactory.createModern()
g = graph.traversal()
g.V(1).outE('knows').inV().values('name')
g.V(1).out('knows').values('name')
g.V(1).out('knows').has('age', gt(30)).values('name')
```

## Next 15 minutes
- Create a graph. Note that id and label are built-in reserved property names for TinkerPop. In the addVertex() command below, all the properties and values are listed in ordered pairs.
```
graph = TinkerGraph.open()
v1 = graph.addVertex(id, 1, label, "person", "name", "marko", "age", 29)
v2 = graph.addVertex(T.id, 3, T.label, "software", "name", "lop", "lang", "java")
v1.addEdge("created", v2, id, 9, "weight", 0.4)

```
- Note that the REPL permits identifier assignments, which is not true for most property graph databases, but, obviously, helpful in a REPL!
- Now let's traverse, bit by bit, askint "What software has Marko created?"
```
// get a reference to the graph first
g = graph.traversal()

g.V().has('name','marko')
g.V().has('name','marko').outE('created').inV()
// or simpler since we're not doing anything with the edge:
g.V().has('name','marko').out('created')
g.V().has('name','marko').out('created').values('name')

```
























# EoF
