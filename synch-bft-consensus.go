package main

import (
	"fmt"
	"math/rand"
)

// message = 1 (attack) or 2 (retreat), loyal = false if traitor
type General struct {
	message int
	loyal   bool
}

func Byz(l, t int) [][]int {

	var n int
	if (l+t)%2 == 0 {
		n = l + t + 1
	} else {
		n = l + t
	}

	// recieve and send messages through this array
	messages := make([][]int, n)
	for i := range messages {
		messages[i] = make([]int, n)
	}

	genPool := make(map[int]General)
	decisionTracker := make([][]int, t+2)
	decisionTracker[0] = make([]int, n)
	fmt.Println(len(decisionTracker))

	// initialize dictionary with first "k" generals as traitors
	traitors := 0
	for i := 0; i < n; i++ {
		var general General

		if traitors < t {
			general.loyal = false
			traitors++
		} else {
			general.loyal = true
		}

		// assign a message only if general is loyal
		if general.loyal {
			message := rand.Intn(2) + 1
			general.message = message
		}

		genPool[i] = general
		decisionTracker[0][i] = general.message
	}

	// t+1 total super-rounds
	for k := 0; k < t+1; k++ {

		// round 1
		for gen := range genPool {
			if genPool[gen].loyal {
				writeMessageB(messages, gen, genPool)
			} else { // traitor general, randomly sends messages
				for cell := range messages[gen] {
					messages[gen][cell] = rand.Intn(2) + 1
				}
			}
		}

		// round 2
		leader, _ := majority(messages, k)

		// reading step
		for gen := range genPool {
			if genPool[gen].loyal {
				general := genPool[gen]
				decision, count := majority(messages, gen)

				if count >= l {
					general.message = decision

				} else {
					if genPool[k].loyal {
						general.message = leader
					} else { // k'th general is traitor and sends different messages to the generals
						general.message = rand.Intn(2) + 1
					}
				}
				genPool[gen] = general
			}
		}

		// add decisions to array
		decisionTracker[k+1] = make([]int, n)
		for j := range decisionTracker[k+1] {
			decisionTracker[k+1][j] = genPool[j].message
		}
		fmt.Println("rounds done:", k+1)
		//printArray(messages)
	}

	return decisionTracker

}

func majority(array [][]int, k int) (int, int) {
	ones, twos := 0, 0
	for i := range array {
		if array[i][k] == 1 {
			ones++
		} else if array[i][k] == 2 {
			twos++
		}
	}
	if ones > twos {
		return 1, ones
	}
	return 2, twos
}

func writeMessageB(messages [][]int, gen int, genPool map[int]General) {
	for i := range messages[gen] {
		messages[gen][i] = genPool[gen].message
	}
}

func randBool() bool {
	return rand.Float32() < 0.5
}
