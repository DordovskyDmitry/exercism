package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[byte]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

var invalidDataError = errors.New("invalid nucleotide")

// Count counts number of occurrences of given nucleotide in given DNA.
//
// This is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Count method has a receiver of type DNA named d.
func (d DNA) Count(nucleotide byte) (int, error) {
	if isInvalid(nucleotide) {
		return 0, invalidDataError
	}
	s := 0
	ok := reduce(d, func(r byte) {
		if r == nucleotide {
			s++
		}
	})
	if !ok {
		return 0, invalidDataError
	}
	return s, nil
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{'A': 0, 'C': 0, 'T': 0, 'G': 0}
	ok := reduce(d, func(r byte) { h[r]++ })
	if !ok {
		return h, invalidDataError
	}
	return h, nil
}

func isInvalid(r byte) bool {
	return !(r == 'A' || r == 'C' || r == 'T' || r == 'G')
}

func reduce(d DNA, f func(r byte)) bool {
	for i := 0; i < len(d); i++ {
		r := d[i]
		if isInvalid(r) {
			return false
		}
		f(r)
	}
	return true
}
