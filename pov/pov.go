package pov

import (
	"fmt"
)

type Graph struct {
	nodes map[string]int
	names []string
	arcs  [][]int
}

func New() *Graph {
	g := Graph{
		nodes: nil,
		arcs:  nil,
	}
	return &g
}

func (g *Graph) nodeExist(l string) bool {
	if _, ok := g.nodes[l]; !ok {
		return false
	}
	return true
}

func (g *Graph) getParents(l string) (parents []int) {
	nodeIndex := g.nodes[l]
	for frIndex, tto := range g.arcs {
		for _, to := range tto {
			if to == nodeIndex {
				parents = append(parents, frIndex)
			}
		}
	}
	return parents
}

func (g *Graph) getChilds(l string) (childs []int) {
	nodeIndex := g.nodes[l]
	childs = make([]int, len(g.names))
	for _, to := range g.arcs[nodeIndex] {
		childs = append(childs, to)
	}
	return childs
}

func (g *Graph) AddNode(l string) {
	if g.nodes == nil {
		g.names = append(g.names, l)
		m := make(map[string]int)
		m[l] = 0
		g.nodes = m
	} else {
		g.nodes[l] = len(g.names)
		g.names = append(g.names, l)
	}
}

func (g *Graph) AddArc(from, to string) {
	var frIndex, toIndex int
	var ok bool
	if g.arcs == nil {
		g.arcs = make([][]int, 100)
	}
	if frIndex, ok = g.nodes[from]; !ok {
		g.AddNode(from)
	}
	if toIndex, ok = g.nodes[to]; !ok {
		panic("Wrong test case...")
	}
	frIndex = g.nodes[from]
	toIndex = g.nodes[to]
	g.arcs[frIndex] = append(g.arcs[frIndex], toIndex)
}

func (g *Graph) ArcList() []string {
	var s []string
	for frIndex, tto := range g.arcs {
		for _, to := range tto {
			arcFmt := fmt.Sprintf("%s -> %s", g.names[frIndex], g.names[to])
			s = append(s, arcFmt)
		}
	}
	return s
}

/*
func (g *Graph) ChangeRoot(root, newRoot string) *Graph {
	var swapped bool
	g2 := New()
	for k, _ := range g.nodes {
		g2.AddNode(k)
	}
	for frIndex, tto := range g.arcs {
		fmt.Println("ChangeRoot_g2: ", g2)
		if frIndex != g.nodes[newRoot] {
			fmt.Println("ChangeRoot_NotNewRoot: ", g2)
			swapped = false
			childs := g.getChilds(g.names[frIndex])
			for child := range childs {
				if child == g.nodes[newRoot] {
					swapped = true
				}
			}
			if swapped {
				g2.AddArc(newRoot, g.names[frIndex])
				for child := range childs {
					if child != g.nodes[newRoot] {
						g2.AddArc(newRoot, g.names[child])
					}
				}
			}
		} else {
			fmt.Println("ChangeRoot_NewRoot: ", g2)
			for _, to := range tto {
				g2.AddArc(newRoot, g.names[to])
			}
		}
	}

	return g2
}*/

func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	nr := g.nodes[newRoot]
	var f func(int) bool
	f = func(n int) (found bool) {
		if n == nr {
			return true
		}
		a := g.arcs[n]
		for i, to := range a {
			if f(to) {
				last := len(a) - 1
				a[i] = a[last]
				g.arcs[n] = a[:last]
				g.arcs[to] = append(g.arcs[to], n)
				return true
			}
		}
		return false
	}
	f(g.nodes[oldRoot])
	return g
}

func main() {
}
