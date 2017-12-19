package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

type Garden struct {
	childPlants map[string][]string
}

var plants = map[rune]string{
	'C': "clover",
	'G': "grass",
	'R': "radishes",
	'V': "violets",
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	plantations := strings.Split(diagram, "\n")
	if len(plantations) != 3 {
		return nil, errors.New("invalid input")
	}
	plantations = plantations[1:]

	length := len(children)
	kids := make([]string, length)
	copy(kids, children)
	sort.Strings(kids)
	childPlants := make(map[string][]string, 0)
	for _, plantation := range plantations {
		if len(plantation) != length*2 {
			return nil, errors.New("invalid input")
		}
		for i, v := range plantation {
			kid := kids[i/2]
			plant, ok := plants[v]
			if !ok {
				return nil, errors.New("invalid input")
			}
			childPlants[kid] = append(childPlants[kid], plant)
		}
	}
	if len(childPlants) != length {
		return nil, errors.New("invalid input")
	}
	return &Garden{childPlants: childPlants}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	v, ok := g.childPlants[child]
	return v, ok
}
