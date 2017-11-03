package foodchain

import "fmt"

const testVersion = 3

type triples struct {
	animal, secondLine, addition string
}

var words = []triples{
	triples{
		animal:     "fly",
		secondLine: "",
		addition:   "",
	},
	triples{
		animal:     "spider",
		secondLine: "It wriggled and jiggled and tickled inside her.",
		addition:   " that wriggled and jiggled and tickled inside her",
	},
	triples{
		animal:     "bird",
		secondLine: "How absurd to swallow a bird!",
		addition:   "",
	},
	triples{
		animal:     "cat",
		secondLine: "Imagine that, to swallow a cat!",
		addition:   "",
	},
	triples{
		animal:     "dog",
		secondLine: "What a hog, to swallow a dog!",
		addition:   "",
	},
	triples{
		animal:     "goat",
		secondLine: "Just opened her throat and swallowed a goat!",
		addition:   "",
	},
	triples{
		animal:     "cow",
		secondLine: "I don't know how she swallowed a cow!",
		addition:   "",
	},
	triples{
		animal:     "horse",
		secondLine: "She's dead, of course!",
		addition:   "",
	},
}

func Song() string {
	return Verses(1, len(words))
}

func Verses(from, to int) (s string) {
	for i := from; i < to; i++ {
		s += Verse(i)
		s += "\n\n"
	}
	s += Verse(to)
	return
}

func Verse(n int) (s string) {
	start := words[n-1]
	s = fmt.Sprintf("I know an old lady who swallowed a %s.\n%s", start.animal, start.secondLine)
	if n != 1 && n != len(words) {
		s += "\n"
	}
	if n != len(words) {
		for i := n - 1; i > 0; i-- {
			s += fmt.Sprintf("She swallowed the %s to catch the %s%s.\n", words[i].animal, words[i-1].animal, words[i-1].addition)
		}
		s += fmt.Sprintf("I don't know why she swallowed the %s. Perhaps she'll die.", words[0].animal)
	}
	return
}
