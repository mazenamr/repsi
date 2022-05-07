package abstract

import (
	"fmt"
	"log"
	"repsi/consts"
	"sort"
	"strings"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

func (a *AbstractMachine) Draw(filename string) {
	g := graphviz.New()

	graph, err := g.Graph(graphviz.Directed)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := g.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	graph.SetRankDir(cgraph.LRRank)

	nodes := make(map[string]*cgraph.Node)

	for name, s := range a.States {
		n, err := graph.CreateNode(name)
		n.SetPenWidth(consts.NodePenWidth)
		n.SetStyle(cgraph.FilledNodeStyle)
		nodes[name] = n
		if err != nil {
			log.Fatal(err)
		}
		if s.IsTerminatingState {
			n.SetShape(cgraph.DoubleCircleShape)
			n.SetFillColor(consts.NodeTerminatingColor)
		} else if name == a.StartingState {
			n.SetShape(cgraph.CircleShape)
			n.SetFillColor(consts.NodeStartingColor)
		} else {
			n.SetShape(cgraph.CircleShape)
			n.SetFillColor(consts.NodeColor)
		}
	}

	n, err := graph.CreateNode("")
	if err != nil {
		log.Fatal(err)
	}
	n.SetShape(cgraph.NoneShape)
	e, err := graph.CreateEdge("", n, nodes[a.StartingState])
	if err != nil {
		log.Fatal(err)
	}
	e.SetLen(consts.EdgeLen)
	e.SetPenWidth(consts.EdgePenWidth)
	e.SetArrowSize(consts.ArrowSize)

	edges := make(map[string]*cgraph.Edge)
	moves := make(map[string][]string)

	for from, s := range a.States {
		for t, to := range s.Moves {
			for _, to := range to {
				edgeName := fmt.Sprintf("M-%s-%s", from, to)
				if _, ok := edges[edgeName]; !ok {
					e, err := graph.CreateEdge(edgeName, nodes[from], nodes[to])
					if err != nil {
						log.Fatal(err)
					}
					edges[edgeName] = e
					moves[edgeName] = make([]string, 0)
					e.SetLen(consts.EdgeLen)
					e.SetPenWidth(consts.EdgePenWidth)
					e.SetArrowSize(consts.ArrowSize)
				}

				switch t {
				case string(consts.WildcardToken):
					t = "<any>"
				case "":
					t = "<empty>"
				case " ":
					t = "<space>"
				case "\t":
					t = "<tab>"
				case "\n":
					t = "<newline>"
				case ",":
					t = "<comma>"
				case ".":
					t = "<dot>"
				}

				if len(t) > 1 {
					if t[0] == '[' {
						if t[1] != '^' {
							t = fmt.Sprintf("<any of> [%s]", t[2:len(t)-1])
						} else {
							t = fmt.Sprintf("<any except> [%s]", t[3:len(t)-1])
						}
					}

				}

				moves[edgeName] = append(moves[edgeName], t)
				sort.SliceStable(moves[edgeName], func(i, j int) bool {
					return moves[edgeName][i] < moves[edgeName][j]
				})

				lable := strings.Join(moves[edgeName], ", ")
				edges[edgeName].SetLabel(lable)
			}
		}
	}

	if err := g.RenderFilename(graph, graphviz.SVG, fmt.Sprintf("%s.svg", filename)); err != nil {
		log.Fatal(err)
	}
}
