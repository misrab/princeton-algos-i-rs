package main

import (
	"bufio"
	//"log"
	"os"
	"strconv"
	"strings"

	"github.com/misrab/stanford-algos-rs/go-src/digraph"
)

const (
	PATH = "../data/scc_small.txt"
)

func main() {
	file, err := os.Open(PATH)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	c := make(chan string)
	done := make(chan struct{})
	go handleLine(c, done)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	close(c)
	<-done

}

func handleLine(c chan string, done chan struct{}) {
	graph := digraph.NewDiGraph()

	for line := range c {
		values := strings.Split(line, " ")
		strintfrom, _ := strconv.Atoi(values[0])
		strintto, _ := strconv.Atoi(values[1])

		from := uint64(strintfrom)
		to := uint64(strintto)

		graph.AddEdge(from, to)
	}

	// we're going to find strongly connected components
	reversed_graph := graph.Reverse()
	labels_old_to_new, labels_new_to_old := digraph.TopologicallyOrder(reversed_graph)

	println("done")
	close(done)
}
