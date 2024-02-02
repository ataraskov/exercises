// Pacakge tournament provides a simple tournament score system.
package tournament

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

// Team represents a team in a tournament.
type Team struct {
	Name    string
	Matches int
	Wins    int
	Ties    int
	Losses  int
}

// NewTeamFromResult parses a result string and returns teams.
func NewTeamFromResult(s string) (Team, Team, error) {
	pieces := strings.Split(s, ";")
	if len(pieces) != 3 {
		return Team{}, Team{}, fmt.Errorf("bad formating, found %d fields", len(pieces))
	}

	t1 := Team{}
	t1.Name = pieces[0]
	t1.Matches = 1
	t2 := Team{}
	t2.Name = pieces[1]
	t2.Matches = 1
	switch pieces[2] {
	case "win":
		t1.Wins += 1
		t2.Losses += 1
	case "loss":
		t1.Losses += 1
		t2.Wins += 1
	case "draw":
		t1.Ties += 1
		t2.Ties += 1
	default:
		return Team{}, Team{}, fmt.Errorf("bad match outcome: %v", pieces[2])
	}
	return t1, t2, nil
}

// Points returns the team's point score.
func (t *Team) Point() int {
	return t.Wins*3 + t.Ties
}

// Results returns the team's results.
func parse(r io.Reader) (map[string]Team, error) {
	buf, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(buf), "\n")
	teams := make(map[string]Team)
	for _, l := range lines {
		if len(l) == 0 || strings.HasPrefix(l, "#") {
			continue
		}
		t1, t2, err := NewTeamFromResult(l)
		if err != nil {
			return nil, err
		}
		for _, t := range []Team{t1, t2} {
			if _, found := teams[t.Name]; !found {
				teams[t.Name] = t
				continue
			}
			team := teams[t.Name]
			teams[t.Name] = Team{
				Name:    t.Name,
				Matches: team.Matches + 1,
				Wins:    team.Wins + t.Wins,
				Losses:  team.Losses + t.Losses,
				Ties:    team.Ties + t.Ties,
			}
		}
	}
	return teams, nil
}

// results writes tournament results.
func results(w io.Writer, teams map[string]Team) error {
	buf := strings.Builder{}
	format := "%-30v | %2v | %2v | %2v | %2v | %2v\n"
	buf.WriteString(fmt.Sprintf(format, "Team", "MP", " W", " D", " L", " P"))

	names := []string{}
	for k := range teams {
		names = append(names, k)
	}

	sort.Slice(names, func(i, j int) bool {
		t1 := teams[names[i]]
		t2 := teams[names[j]]
		if t1.Point() == t2.Point() {
			return t1.Name < t2.Name
		}
		return t1.Point() > t2.Point()
	})
	for _, name := range names {
		team := teams[name]
		buf.WriteString(fmt.Sprintf(format,
			team.Name,
			team.Matches,
			team.Wins,
			team.Ties,
			team.Losses,
			team.Point(),
		))
	}
	_, err := w.Write([]byte(buf.String()))
	if err != nil {
		return err
	}
	return nil
}

// Tally reads matches and writes a tournament results.
func Tally(reader io.Reader, writer io.Writer) error {
	teams, err := parse(reader)
	if err != nil {
		return err
	}
	results(writer, teams)
	if err != nil {
		return err
	}
	return nil
}
