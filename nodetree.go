package nmn

// Node represents a node in the node tree
type NodeTree struct {
	RootId string
	Nodes  map[string]*Node
}

// NewNodeTree creates a new NodeTree
func NewNodeTree() *NodeTree {
	return &NodeTree{
		Nodes: make(map[string]*Node),
	}
}

// AddNode adds a node to the NodeTree
// The node is added to the Nodes map with the node's id as the key
func (nt *NodeTree) AddNode(n *Node) {
	nt.Nodes[n.Id] = n
}

// GetNode returns a node from the NodeTree
// The node is retrieved from the Nodes map using the node's id as the key
func (nt *NodeTree) GetNode(id string) *Node {
	return nt.Nodes[id]
}

// GetRoot returns the root node of the NodeTree
func (nt *NodeTree) GetRoot() *Node {
	return nt.Nodes[nt.RootId]
}
