package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
)

type match struct {
	home   string
	away   string
	result string
}

type team struct {
	name string
	MP   int
	W    int
	D    int
	L    int
	P    int
}

// Tally takes in a multiline string where each line is in the format
// home_team_name;away_team_name;result
// and outputs the standings of the teams in a formatted table.
func Tally(input io.Reader, output io.Writer) error {
	matchStrings := make([]string, 0)
	matches := make([]match, 0)
	teams := make(map[string]team)
	raw, err := ioutil.ReadAll(input)
	if err != nil {
		return err
	}
	buffer := make([]byte, 0)
	comment := false
	// slice up the data by newlines
	for _, datum := range raw {
		// ignore new lines and comments, else fill buffer
		if string(datum) == "\n" && len(buffer) == 0 && comment == false {
			continue
		} else if string(datum) == "#" && len(buffer) == 0 {
			comment = true
			continue
		} else if string(datum) == "\n" && comment == true {
			comment = false
			continue
		} else if comment == true {
			continue
		}

		if string(datum) == "\n" {
			matchStrings = append(matchStrings, string(buffer[:]))
			buffer = make([]byte, 0)
		} else {
			buffer = append(buffer, datum)
		}
	}
	if len(matchStrings) == 0 {
		return errors.New("input not in proper format - no new lines encountered")
	}
	// split each line and convert it to a match
	for index, line := range matchStrings {
		pieces := strings.Split(line, ";")
		if len(pieces) != 3 {
			return errors.New("line " + strconv.Itoa(index+1) + " does not match the 3 piece format")
		}
		thisMatch := match{
			home:   pieces[0],
			away:   pieces[1],
			result: pieces[2],
		}
		matches = append(matches, thisMatch)
	}
	// read each match and enter into team stats
	for index, game := range matches {
		// check if the teams have stats yet
		if _, ok := teams[game.home]; !ok {
			teams[game.home] = team{name: game.home}
		}
		if _, ok := teams[game.away]; !ok {
			teams[game.away] = team{name: game.away}
		}
		home := teams[game.home]
		away := teams[game.away]
		// now adjust stats depending on result
		switch game.result {
		case "win":
			home.W++
			home.MP++
			home.P += 3
			away.L++
			away.MP++
		case "draw":
			home.D++
			home.MP++
			home.P++
			away.D++
			away.MP++
			away.P++
		case "loss":
			away.W++
			away.MP++
			away.P += 3
			home.L++
			home.MP++
		default:
			return errors.New("not a valid result for match " + strconv.Itoa(index))
		}

		teams[game.home] = home
		teams[game.away] = away
	}
	// now sort teams by points, then by name
	results := make([]team, 0)
	for _, entry := range teams {
		results = append(results, entry)
	}
	sort.Slice(results, func(i, j int) bool {
		if results[i].P == results[j].P {
			return results[i].name < results[j].name
		}
		return results[i].P > results[j].P
	})
	// finally, write the output
	// table := "Team                           | MP |  W |  D |  L |  P" // 31
	writer := tabwriter.NewWriter(output, 4, 0, 0, ' ', tabwriter.Debug)
	fmt.Fprintln(writer, "Team                           \t MP\t  W\t  D\t  L\t  P")
	for _, result := range results {
		fmt.Fprintln(writer, result.name+"\t  "+strconv.Itoa(result.MP)+"\t  "+strconv.Itoa(result.W)+"\t  "+strconv.Itoa(result.D)+"\t  "+strconv.Itoa(result.L)+"\t  "+strconv.Itoa(result.P))
	}
	writer.Flush()

	return nil
}
