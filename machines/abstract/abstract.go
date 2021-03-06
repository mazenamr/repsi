package abstract

import (
	"encoding/json"
	"fmt"
	"os"
)

type Machine struct {
	StartingState string
	States        map[string]State
}

type State struct {
	IsTerminatingState bool
	Moves              map[string][]string
}

func (a State) AddMove(token string, state string) {
	a.Moves[token] = append(a.Moves[token], state)
}

func (a *Machine) Json() string {
	j, _ := json.MarshalIndent(a, "", "  ")
	return string(j)
}

func Load(filename string) *Machine {
	file, _ := os.Open(filename)
	defer file.Close()
	var a Machine
	json.NewDecoder(file).Decode(&a)
	return &a
}

func (a *Machine) Save(filename string) {
	f, _ := os.Create(fmt.Sprintf("%s.json", filename))
	defer f.Close()
	f.WriteString(a.Json())
}

func (a *Machine) Out(filename string) {
	a.Save(filename)
	a.Draw(filename)
}
