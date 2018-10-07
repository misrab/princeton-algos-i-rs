package graph

import (
	"testing"

	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	ADJACENCY_PATH = "../../data/kargerMinCut.txt"
)


func TestNewGraphFromEdgeList(t *testing.T) {
	g := NewGraph()

	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)
	g.AddEdge(1, 3, 3)
	g.AddEdge(1, 4, 4)
	g.AddEdge(4, 3, 5)

	mst := g.FindMST()
	log.Printf("%+v\n", mst)
}



func TestNewGraph(t *testing.T) {
	return

	graph := NewGraph()

	// insert nodes
	graph.insertNodeAdjacency(1, []uint64{2, 3, 4})
	graph.insertNodeAdjacency(2, []uint64{1, 3})
	graph.insertNodeAdjacency(3, []uint64{2, 1, 4})
	graph.insertNodeAdjacency(4, []uint64{1, 3})

	//log.Printf("%v\n", graph)
	//edge := graph.GetEdges()[0]
	//log.Printf("contracting %v\n", edge)
	//graph.ContractEdge(edge)
	//log.Printf("%v\n", graph)

	log.Println("contracting entire graph...")
	cuts := graph.ContractionAlgorithm()
	log.Printf("Got %d cuts for graph:\n %v\n", cuts, graph)
}

// read an adjacency list into a graph
func TestReadAdjacencyList(t *testing.T) {
	return

	t.Skip("skipping from file")
	var min uint64 = 9999999999
	num_trials := 1

	for i := 0; i < num_trials; i++ {
		graph, err := graphFromFile()
		if err != nil {
			log.Fatal(err)
		}
		//log.Println("contracting entire graph...")
		cuts := graph.ContractionAlgorithm()
		//log.Printf("Got %d cuts for graph:\n %v\n", cuts, graph)
		if cuts < min {
			min = cuts
		}

		log.Printf("Found mincut of %d in %d trials\n%v", min, num_trials, graph.GetEdges())
	}

}

func graphFromFile() (Graph, error) {
	graph := NewGraph()

	file, err := os.Open(ADJACENCY_PATH)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ints, err := stringToUint64Array(scanner.Text())
		if err != nil {
			return nil, err
		}
		//log.Println(ints)

		// add to graph
		graph.insertNodeAdjacency(ints[0], ints[1:])

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return graph, err
}

// A is the string
func stringToUint64Array(A string) ([]uint64, error) {
	a := strings.Split(A, "\t")
	b := make([]uint64, 0)
	for _, v := range a {
		if v == "" {
			continue
		}

		theint, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		b = append(b, uint64(theint))
	}

	return b, nil
}
