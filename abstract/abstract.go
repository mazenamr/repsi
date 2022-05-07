package abstract

import (
	"encoding/json"
	"fmt"
	"os"
)

type AbstractMachine struct {
	StartingState string
	States        map[string]AbstractState
}

type AbstractState struct {
	IsTerminatingState bool
	Moves              map[string][]string
}

func (a AbstractState) AddMove(token string, state string) {
	a.Moves[token] = append(a.Moves[token], state)
}

func (a *AbstractMachine) Json() string {
	j, _ := json.MarshalIndent(a, "", "  ")
	return string(j)
}

func (a *AbstractMachine) Write(filename string) {
	f, _ := os.Create(fmt.Sprintf("%s.json", filename))
	defer f.Close()
	f.WriteString(a.Json())
}

func (a *AbstractMachine) Out(filename string) {
	a.Write(filename)
	a.Draw(filename)
}
