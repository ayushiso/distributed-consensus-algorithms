package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	Agreement(10, 5)
}

type Student struct {
	message  int
	drop     bool
	willdrop bool
}

func Agreement(n, t int) {

	// recieve and send messages through this array
	messages := make([][]int, n)
	for i := range messages {
		messages[i] = make([]int, n)
	}

	studentPool := make([]Student, n)

	//initialize student initial messages
	for i := range studentPool {
		studentPool[i].message = rand.Intn(100) + 1
		studentPool[i].drop = false
		studentPool[i].willdrop = false
		//fmt.Println(i, studentPool[i])
	}

	// for i := range studentPool {
	// 	fmt.Println(studentPool[i])
	// }

	for i := 0; i <= t; i++ {

		numDropped := 0

		// erase messages of all students who dropped during last round
		for i := range studentPool {
			if messages[i][0] != 0 && studentPool[i].drop {
				for j := range messages[i] {
					messages[i][j] = 0
				}
			}

			if studentPool[i].drop {
				numDropped++
			}
		}

		// choose how many students are going to drop
		numDrops := rand.Intn(t - numDropped + 1)
		fmt.Println(numDrops)
		// choose which students are going to drop
		dropList := randNums(n, numDrops, studentPool)
		fmt.Println(dropList)

		for j := range studentPool {
			//fmt.Println(student.message)

			if !studentPool[j].drop {
				if dropList[j] {
					studentPool[j].willdrop = true
				}

				writeMessage(j, studentPool, messages)
			}
		}

		for j := range studentPool {
			//fmt.Println(studentPool[j].drop)
			if !studentPool[j].drop {
				studentPool[j].message = chooseMessage(j, messages)
			}
		}

		printArray(messages)
		fmt.Println("rounds done:", i+1)
	}
}

// randNums chooses num random numbers in range (0, n] and returns a dict of the chosen numbers
func randNums(length int, num int, studentPool []Student) map[int]bool {
	randNums := make(map[int]bool)
	i := 0

	for i < num {
		newNum := rand.Intn(length)
		//fmt.Println(newNum)
		if randNums[newNum] || studentPool[i].drop == true { //number has already been added to dropList, or this was dropped in previous rounds
			continue
		}
		randNums[newNum] = true
		i++
	}

	return randNums
}

func writeMessage(j int, studentPool []Student, messages [][]int) {

	// student will drop out in this round at some point
	if studentPool[j].willdrop {
		studentPool[j].drop = true

		//at what time will this student drop?
		dropIndex := rand.Intn(len(messages))

		for i := range messages[j] {
			if i >= dropIndex {
				messages[j][i] = 0
			} else {
				messages[j][i] = studentPool[j].message
			}
		}

		// clean round by this student
	} else {
		for i := range messages[j] {
			messages[j][i] = studentPool[j].message
			//fmt.Println(student.message)
		}
	}
}

// chooseMessage returns the minimum non-zero element in the j'th column of messages
func chooseMessage(j int, messages [][]int) int {
	var min int
	var minIndex int

	// find first non-zero element
	// there has to be atleast one non-zero element (because this student didn't drop, so it atleast has its own message)
	for i := range messages {
		if messages[i][j] != 0 {
			min = messages[i][j]
			minIndex = i
			break
		}
	}

	// range from first non-zero element to end, find minimum non-zero message
	for i := minIndex; i < len(messages); i++ {
		if messages[i][j] < min && messages[i][j] != 0 {
			min = messages[i][j]
		}
	}

	return min
}

func printArray(arr [][]int) {
	for elems := range arr {
		fmt.Println(arr[elems])
	}
}
