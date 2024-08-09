package main

import (
  "fmt"
  "sort"
)

type ConsistentHash struct {
  nodeCount   int 
  nodes       [] string
  totalSlots  int
}

func (cHash *ConsistentHash) nextHash (nodeHash string) (int, string) {
  low := 0
  high := cHash.nodeCount - 1

  for low < high {
    mid := (low + high) / 2

    if cHash.nodes[mid] < nodeHash {
      low = mid + 1
    } else {
      high = mid - 1
    }
  }

  return low, cHash.nodes[low]
}

func (cHash *ConsistentHash) addNode(node string) {
  newHash := sha256Hash(node)
  nextNodeIdx, _ := cHash.nextHash(newHash)
  newLength := cHash.nodeCount + 1

  // [TODO]: iterate over the keys in this nextNodeHash and migrate the keys to the newHash node server accordingly

  newNodes := make([]string, newLength)

  if nextNodeIdx == 0 {
    newNodes[0] = newHash 
    copy(newNodes[1:], cHash.nodes[:])
  } else if nextNodeIdx == cHash.nodeCount {
    copy(newNodes[:], cHash.nodes[:])
    newNodes[cHash.nodeCount] = newHash
  } else {
    copy(newNodes[:nextNodeIdx], cHash.nodes[:nextNodeIdx])
    newNodes[nextNodeIdx] = newHash
    copy(newNodes[nextNodeIdx+1:], cHash.nodes[nextNodeIdx:])
  }

  cHash.nodes = newNodes
  cHash.nodeCount = len(cHash.nodes)

} 

func (cHash *ConsistentHash) removeNode(node string) {
  if cHash.nodeCount <= 0 {
    fmt.Println("[WARN]: [removeNode] no new nodes to remove")
    return
  }

  nodeHash := sha256Hash(node)

  // [TODO]: migrate all keys present in this `node` and move it to the  next node 
  newLength := cHash.nodeCount - 1

  newNodes := make([]string, newLength)

  var found bool = false
  newNodesIdx := 0

  for i := 0; i < cHash.nodeCount; i++ {
    if cHash.nodes[i] == nodeHash {
      // skip this
      found = true
    } else {
      newNodes[newNodesIdx] = cHash.nodes[i]
      newNodesIdx = newNodesIdx + 1
    }
  }

  if !found {
    fmt.Println("[WARN]: [removeNode]: Node is not present in the ring")
    return
  }

  cHash.nodes = newNodes
  cHash.nodeCount = len(cHash.nodes)
}

func (cHash *ConsistentHash) findNodeForKey(key string) string {
  keyHash := sha256Hash(key)
  _, nodeHash := cHash.nextHash(keyHash)

  return nodeHash
}

func (cHash *ConsistentHash) listNodes() {
  fmt.Println("*** [begin] node server list ***")
  fmt.Println("TOTAL NODES: ", cHash.nodeCount) 
  for i := 0; i < cHash.nodeCount; i++ {
    fmt.Println(cHash.nodes[i])
  }
  fmt.Println("*** [end] node server list ***")
}

func newConsistentHash(nodeCount int, servers []string) *ConsistentHash {
  chash := &ConsistentHash{}

  chash.nodeCount = nodeCount

  for i := 0; i < len(servers); i++ {
    // hash the node
    h := sha256Hash(servers[i])
    chash.nodes = append(chash.nodes, h)
  }

  // arrange the node hash in increasing order
  sort.Strings(chash.nodes[:])

  return chash
}
