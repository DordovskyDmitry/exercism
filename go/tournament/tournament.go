package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type Result struct {
	mp, w, d, l, p int
}

func Tally(r io.Reader, buffer io.Writer) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	strs := matchStrings(string(b))
	teamsToResult, err := results(strs)
	if err != nil {
		return err
	}
	teams := allTeams(teamsToResult)
	teams = sortTeams(teams, teamsToResult)
	writeResults(buffer, teams, teamsToResult)
	return nil
}

func matchStrings(b string) (res []string) {
	strs := strings.Split(b, "\n")
	for _, str := range strs {
		if len(str) > 0 && str[0] != '#' {
			res = append(res, str)
		}
	}
	return
}

func results(strs []string) (map[string]*Result, error) {
	results := make(map[string]*Result, 0)
	for _, str := range strs {
		t1, t2, result, ok := parseLine(str)
		if !ok || (result != "win" && result != "draw" && result != "loss") {
			return nil, errors.New("Invalid line")
		}
		if _, ok := results[t1]; !ok {
			results[t1] = &Result{}
		}
		if _, ok := results[t2]; !ok {
			results[t2] = &Result{}
		}
		results[t1].mp += 1
		results[t2].mp += 1
		switch result {
		case "win":
			results[t1].w += 1
			results[t1].p += 3
			results[t2].l += 1
		case "draw":
			results[t1].d += 1
			results[t1].p += 1
			results[t2].d += 1
			results[t2].p += 1
		case "loss":
			results[t1].l += 1
			results[t2].w += 1
			results[t2].p += 3
		}

	}
	return results, nil
}

func parseLine(str string) (string, string, string, bool) {
	s := strings.Split(str, ";")
	if len(s) != 3 {
		return "", "", "", false
	}
	return s[0], s[1], s[2], true
}

func allTeams(teamsToResult map[string]*Result) []string {
	res := make([]string, 0, len(teamsToResult))
	for k := range teamsToResult {
		res = append(res, k)
	}
	return res
}

func sortTeams(teams []string, teamsToResult map[string]*Result) []string {
	sort.SliceStable(teams, func(i, j int) bool {
		return (teamsToResult[teams[i]].p > teamsToResult[teams[j]].p) ||
			(teamsToResult[teams[i]].p == teamsToResult[teams[j]].p && teams[i] < teams[j])
	})
	return teams
}

func writeResults(buffer io.Writer, teams []string, teamsToResult map[string]*Result) {
	io.WriteString(buffer, "Team                           | MP |  W |  D |  L |  P\n")
	for _, team := range teams {
		writeResultLine(buffer, team, teamsToResult[team])
	}
}

func writeResultLine(buffer io.Writer, team string, r *Result) {
	io.WriteString(buffer, fmt.Sprintf("%-30s |%3d |%3d |%3d |%3d |%3d\n", team, r.mp, r.w, r.d, r.l, r.p))
}
