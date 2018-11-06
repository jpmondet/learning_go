package pov

import (
	"fmt"
)

type Arc struct {
	fr string
	to string
}

type Node struct {
	name string
}

type Graph struct {
	nodes []Node
	arcs  []Arc
	root  string
}

func New() *Graph {
	g := Graph{
		nodes: nil,
		arcs:  nil,
	}
	return &g
}

func (g *Graph) AddNode(l string) {
	if g.nodes == nil {
		g.nodes = make([]Node, 1)
		g.nodes[0] = Node{name: l}
		g.root = l
	} else {
		n := Node{name: l}
		g.nodes = append(g.nodes, n)
	}
}

func (g *Graph) AddArc(from, to string) {
	if g.arcs == nil {
		g.arcs = make([]Arc, 1)
		g.arcs[0] = Arc{fr: from, to: to}
	} else {
		a := Arc{fr: from, to: to}
		g.arcs = append(g.arcs, a)
	}
}

func (g *Graph) ArcList() []string {
	var s []string
	for _, arc := range g.arcs {
		arcFmt := fmt.Sprintf("%s -> %s", arc.fr, arc.to)
		s = append(s, arcFmt)
	}
	return s
}

func (g *Graph) NodeArcs(node string) (arcNode []Arc) {
	for _, arc := range g.arcs {
		if arc.fr == node || arc.to == node {
			arcNode = append(arcNode, arc)
		}
	}
	return arcNode
}

func (g *Graph) FindNode(node string) Node {
	for _, n := range g.nodes {
		fmt.Println(n)
		if n.name == node {
			return n
		}
	}
	return Node{name: ""}
}

func (g *Graph) Childs(node string) (childs []Node) {
	arcs := g.NodeArcs(node)
	for _, arc := range arcs {
		if arc.to != node {
			child := g.FindNode(arc.to)
			childs = append(childs, child)
			fmt.Println(child, childs)
		}
	}
	return childs
}

func Rearrange(node string, g *Graph, g2 *Graph, parent string) (*Graph, *Graph) {
	var neigh string
	n := g2.FindNode(node)
	if n.name == "" {
		g2.AddNode(node)
		arcs := g.NodeArcs(node)
		for _, arc := range arcs {
			if parent == "" { // We are NewRoot
				if arc.to == node {
					arc.to = arc.fr
					arc.fr = node
				}
				g2.AddArc(arc.fr, arc.to)
				g, g2 = Rearrange(arc.to, g, g2, node)
			} else { // We are a child
				if arc.to != node {
					neigh = arc.to
				} else {
					neigh = arc.fr
				}
				neighNode := g2.FindNode(neigh)
				if neighNode.name == "" {
					if arc.to == parent {
						arc.to = arc.fr
						arc.fr = parent
					}
					g2.AddArc(arc.fr, arc.to)
				}
			}
		}
		childs := g.Childs(node)
		for _, child := range childs {
			g, g2 = Rearrange(child.name, g, g2, node)
		}
	}
	return g, g2
}

func (g *Graph) ChangeRoot(root, newRoot string) *Graph {
	//TODO
	g2 := New()
	g, g2 = Rearrange(newRoot, g, g2, "")
	return g2
}

func main() {
}
