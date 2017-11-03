package house

const testVersion = 1

type pair struct {
	noun, verb string
}

var words = []pair{
	pair{
		noun: "the house that Jack built",
		verb: "",
	},
	pair{
		noun: "the malt",
		verb: "lay in",
	},
	pair{
		noun: "the rat",
		verb: "ate",
	},
	pair{
		noun: "the cat",
		verb: "killed",
	},
	pair{
		noun: "the dog",
		verb: "worried",
	},
	pair{
		noun: "the cow with the crumpled horn",
		verb: "tossed",
	},
	pair{
		noun: "the maiden all forlorn",
		verb: "milked",
	},
	pair{
		noun: "the man all tattered and torn",
		verb: "kissed",
	},
	pair{
		noun: "the priest all shaven and shorn",
		verb: "married",
	},
	pair{
		noun: "the rooster that crowed in the morn",
		verb: "woke",
	},
	pair{
		noun: "the farmer sowing his corn",
		verb: "kept",
	},
	pair{
		noun: "the horse and the hound and the horn",
		verb: "belonged to",
	},
}

func Verse(i int) string {
	return "This is " + recursiveVerse(i) + "."
}

func recursiveVerse(i int) string {
	var currentPair pair = words[i-1]
	if currentPair.verb == "" {
		return currentPair.noun
	}
	return currentPair.noun + "\nthat " + currentPair.verb + " " + recursiveVerse(i-1)
}

func Song() (s string) {
	s = Verse(1)
	for i := 2; i <= len(words); i++ {
		s += "\n\n"
		s += Verse(i)
	}
	return
}
