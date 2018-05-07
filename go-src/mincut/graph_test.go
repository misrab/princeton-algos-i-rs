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

func TestNewGraph(t *testing.T) {
	graph := NewGraph()

	// insert nodes
	graph.insertNodeAdjacency(1, []uint64{2, 3, 4, 5})

	log.Printf("%v\n", graph)
	edge := graph.GetEdges()[0]
	log.Printf("contracting %v\n", edge)
	graph.ContractEdge(edge)
	log.Printf("%v\n", graph)

}

// read an adjacency list into a graph
func TestReadAdjacencyList(t *testing.T) {
	t.Skip("skipping from file")

	_, err := graphFromFile()
	if err != nil {
		log.Fatal(err)
	}

	//for _, node := range graph.GetNodes() {
	//	log.Printf("%v\n", node)
	//}

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
