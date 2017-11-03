package protein

const testVersion = 1

var codonCases = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromCodon(input string) string {
	return codonCases[input]
}

func FromRNA(input string) (res []string) {
	size := len(input)
	for i := 0; i < size-2; i += 3 {
		protein := FromCodon(input[i : i+3])
		if protein == "STOP" {
			return
		}
		res = append(res, protein)
	}
	return
}
