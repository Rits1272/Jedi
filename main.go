package main

import (
  "fmt"
  "jedi/node"
)

func main() {
  fmt.Println("Welcome to Jedi - A distributed, scalable and consistent cache store")
  servers := [...] string {"1", "2", "3", "4", "5"}
  chash := newConsistentHash(5, servers[:]) 
  chash.listNodes()

  node.Start()
}

/*
JEDI
- scalable (consistent hashing)
- supports replication
- heartbeat mechanism and failover support (zookeeper?)
- Jedi client to get/set keys to/from cache store
*/
