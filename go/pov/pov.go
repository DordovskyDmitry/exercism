package pov

type Graph struct {
	rel     map[string][]string
	parents map[string]string
}

func New() *Graph {
	return &Graph{rel: make(map[string][]string), parents: make(map[string]string)}
}
func (g *Graph) AddNode(nodeLabel string) {
	g.rel[nodeLabel] = []string{}
}
func (g *Graph) AddArc(from, to string) {
	g.rel[from] = append(g.rel[from], to)
	g.parents[to] = from
}
func (g *Graph) ArcList() (res []string) {
	for k, v := range g.rel {
		for _, el := range v {
			s := k + " -> " + el
			res = append(res, s)
		}
	}
	return
}
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	if oldRoot == newRoot {
		return g
	}

	// Change relations
	current := newRoot
	path := []string{current}
	for current != oldRoot {
		parent := g.parents[current]
		g.rel[parent] = remove(g.rel[parent], current)
		g.rel[current] = append(g.rel[current], parent)
		current = parent
		path = append(path, current)
	}

	// Change parents
	delete(g.parents, newRoot)
	for i := 0; i < len(path)-1; i++ {
		oldChild := path[i]
		oldParent := path[i+1]
		g.parents[oldParent] = oldChild
	}

	return g
}

func remove(l []string, item string) []string {
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}
