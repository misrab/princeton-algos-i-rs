package main

import (
	"bufio"
	//"log"
	"fmt"
	"os"

	"strconv"
	//"strings"
	//"sort"
	"github.com/misrab/stanford-algos-rs/go-src/medianheap"
)

const (
	PATH = "../data/median.txt"
)

func main() {
	fmt.Println("running main.go...")

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

func str_to_uint64(s string) uint64 {
	as_int, _ := strconv.Atoi(s)
	return uint64(as_int)
}

// template
//func handleLine(c chan string, done chan struct{}) {
//	close(done)
//}

// course 2 - pa 3 (heaps)
func handleLine(c chan string, done chan struct{}) {
	mh := medianheap.NewMedianHeap()

	var median uint64
	var sum uint64 = 0
	medians := make([]uint64, 0)
	for line := range c {
		integer := str_to_uint64(line)

		mh.Insert(integer)
		median = mh.ExtractMedian()
		medians = append(medians, median)

		sum += median

		mh.Insert(median)
	}

	fmt.Printf("%+v\n", medians)
	fmt.Printf("sum md 10k: %d\n", sum%10000)

	close(done)
}

// course 2 - pa 2 (dijkstra)
/*func handleLine(c chan string, done chan struct{}) {
	graph := digraph.NewDiGraph()

	for line := range c {

		values := strings.Split(line, "\t")

		if len(values) < 2 {
			continue
		}

		vertex_id := str_to_uint64(values[0])

		for _, edge_info := range values[1:] {

			id_and_weight := strings.Split(edge_info, ",")

			if len(id_and_weight) < 2 {
				continue
			}

			next_id := str_to_uint64(id_and_weight[0])
			weight := str_to_uint64(id_and_weight[1])

			graph.AddEdge(vertex_id, next_id, weight)
		}
	}

	lenghths_from_one := digraph.GetDistancesFromVertex(graph, 1)

	to_ids := []uint64{7, 37, 59, 82, 99, 115, 133, 165, 188, 197}
	result := ""
	for _, id := range to_ids {
		result += fmt.Sprintf("%d,", lenghths_from_one[id])
	}
	println(result)

	close(done)
}*/

// course 2 - programming assignment 1 (scc's)
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
