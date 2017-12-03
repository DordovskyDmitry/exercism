package beer

import (
	"fmt"
	"strings"
)

func Song() string {
	s, _ := Verses(99, 0)
	return s
}

func Verses(to, from int) (string, error) {
	if to < from || to > 99 || from < 0 {
		return "", fmt.Errorf("Error")
	}
	song := ""
	for i := to; i >= from; i-- {
		s, e := Verse(i)
		if e != nil {
			return "", e
		}
		song += s
		song += "\n"
	}
	return song, nil
}

func Verse(i int) (string, error) {
	if i > 99 || i < 0 {
		return "", fmt.Errorf("Error")
	}
	return firstLine(i) + secondLine(i), nil
}

func firstLine(i int) string {
	count := bottles(i)
	return countBeerOnWall(strings.ToUpper(string(count[0]))+count[1:]) + ", " + countBeer(count) + ".\n"
}

func secondLine(i int) string {
	if i == 0 {
		return fmt.Sprintf("Go to the store and buy some more, %s.\n", countBeerOnWall(bottles(99)))
	}
	count := bottles(i - 1)
	thing := "one"
	if i == 1 {
		thing = "it"
	}
	return fmt.Sprintf("Take %s down and pass it around, %s.\n", thing, countBeerOnWall(count))
}

func bottles(i int) string {
	if i == 0 {
		return "no more bottles"
	}
	if i == 1 {
		return "1 bottle"
	}
	return fmt.Sprintf("%d bottles", i)
}

func countBeerOnWall(s string) string {
	return countBeer(s) + " on the wall"
}

func countBeer(s string) string {
	return s + " of beer"
}
