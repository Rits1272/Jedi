package node

// [TODO]: Support for expiration

type Node struct {
  maxKeys int
  cache   map[string] string
}

func newNode() *Node {
  newNode := &Node{}
  newNode.maxKeys = MAX_KEYS_ALLOWED

  return newNode
}

var node *Node = newNode()

func (node *Node) get(key string) (bool, string) {
  value, exists := node.cache[key]
  return exists, value
}

func (node *Node) set(key string, value string) {
  node.cache[key] = value
  return
}

func Start() {
  StartServer()
}
