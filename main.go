package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	//l := 10000
	//simulateByz(l)
	decisions := Byz(10000, 3000)
	// fmt.Println(decisions)
	writeTable(decisions, "loyal-10000-traitors-3000.tsv")
}

func simulateByz(l int) {
	for t := 0; t <= l/3; t += l / 10 {
		fmt.Println("t=", t)
		decisions := Byz(l, t)
		if unique(decisions[1]) {
			fmt.Println("writing to file")
			filename := "loyal-" + strconv.Itoa(l) + "-traitors-" + strconv.Itoa(t) + ".tsv"
			writeTable(decisions, filename)
		}
	}
}

func simulateMovies(n int) {
	for t := n / 10; t < n; t += n / 10 {
		fmt.Println("t=", t)
		decisions := Agreement(n, t)
		if unique(decisions[1]) {
			fmt.Println("writing to file")
			filename := "total-" + strconv.Itoa(n) + "-drops-" + strconv.Itoa(t) + ".tsv"
			writeTable(decisions, filename)
		}
	}
}

func writeTable(decisions [][]int, filename string) {
	f, _ := os.Create(filename)
	w := bufio.NewWriter(f)

	for row := range decisions {
		for col := range decisions[row] {
			val := strconv.Itoa(decisions[row][col])
			//fmt.Println(val)
			fmt.Fprint(w, val+"\t")
		}
		fmt.Fprint(w, "\n")
	}

	w.Flush()
}

func printArray(arr [][]int) {
	for elems := range arr {
		fmt.Println(arr[elems])
	}
}
