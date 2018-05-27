package main

import (
	"bufio"
	//"log"
	"fmt"
	"os"

	"strconv"
	"strings"

	"sort"

	"github.com/misrab/stanford-algos-rs/go-src/digraph"
)

const (
	PATH = "../data/scc.txt"
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

// template
//func handleLine(c chan string, done chan struct{}) {
//	close(done)
//}

// course 2 - programming assignment 1
/*
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
	sccs := digraph.FindSCCs(graph)
	scc_lens := make([]int, 0)
	for _, list := range sccs {
		scc_lens = append(scc_lens, len(list))
	}
	// decreasing order
	sort.Slice(scc_lens, func(i int, j int) bool { return scc_lens[i] > scc_lens[j] })

	// for programming assignment
	fmt.Printf("%v\n", scc_lens[:6])

	println("done")
	close(done)
}*/
