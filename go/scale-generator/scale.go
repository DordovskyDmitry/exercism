package scale

import "strings"

var sharpScale = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var flatScale = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
var useSharps = map[string]bool{
	"C":  true,
	"G":  true,
	"D":  true,
	"A":  true,
	"a":  true,
	"E":  true,
	"B":  true,
	"F#": true,
	"e":  true,
	"b":  true,
	"f#": true,
	"c#": true,
	"g#": true,
	"d#": true,
}
var indices = make(map[string]int)

func init() {
	for i, v := range sharpScale {
		indices[v] = i
		indices[strings.ToLower(v)] = i
	}
	for i, v := range flatScale {
		if _, ok := indices[v]; !ok {
			indices[v] = i
			indices[strings.ToLower(v)] = i
		}
	}
}

func Scale(tonic, interval string) []string {
	var scale []string
	if useSharps[tonic] {
		scale = sharpScale
	} else {
		scale = flatScale
	}
	return list(scale, indices[tonic], interval)
}

func list(scale []string, index int, interval string) []string {
	if interval == "" {
		return append(scale[index:], scale[:index]...)
	}
	result := []string{}
	for _, v := range interval {
		result = append(result, scale[index%12])
		if v == 'M' {
			index += 2
		} else if v == 'm' {
			index++
		} else {
			index += 3
		}
	}
	return result
}
