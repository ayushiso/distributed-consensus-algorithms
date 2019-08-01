package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	message  int
	willdrop bool
}

// returns true iff there are two unique non-zero values in the array
func unique(arr []int) bool {
	var start int
	var val int
	for i := range arr {
		if arr[i] != 0 {
			start = i
			val = arr[i]
			break
		}
	}

	for i := start + 1; i < len(arr); i++ {
		if arr[i] != val && arr[i] != 0 {
			return true
		}
	}
	return false
}

func Agreement(n, t int) [][]int {

	// recieve and send messages through this array
	messages := make([][]int, n)
	for i := range messages {
		messages[i] = make([]int, n)
	}

	studentPool := make(map[int]Student)
	decisionTracker := make([][]int, t+1)
	decisionTracker[0] = make([]int, n)

	for i := 0; i < n; i++ {
		var student Student
		student.message = rand.Intn(2*n) + 1
		student.willdrop = false
		studentPool[i] = student
		decisionTracker[0][i] = student.message
	}

	for i := 0; i < t; i++ {

		if len(studentPool) == 0 {
			break
		}

		// // erase messages of all students who dropped during last round
		for j := range messages {
			if messages[j][0] != 0 {
				_, exists := studentPool[j]
				if !exists {
					for k := range messages[j] {
						messages[j][k] = 0
					}
				}
			}
		}

		// choose how many students are going to drop
		numDrops := rand.Intn(min(t, len(studentPool)))
		//fmt.Println("numDrops=", numDrops)

		// choose which students are going to drop
		dropList := randNums(n, numDrops, studentPool)

		// writing step (GOROUTINE?)

		for j := range studentPool {

			if dropList[j] {
				student := studentPool[j]
				student.willdrop = true
				studentPool[j] = student
			}

			writeMessageS(j, studentPool, messages)
		}

		// choosing step (GOROUTINE?)

		for j := range studentPool {
			//fmt.Println(studentPool[j].drop)
			student := studentPool[j]
			student.message = chooseMessage(j, messages)
			studentPool[j] = student
		}

		//printArray(messages)
		//fmt.Println(i)
		decisionTracker[i+1] = make([]int, n)
		for j := range decisionTracker[i+1] {
			decisionTracker[i+1][j] = studentPool[j].message
		}
		fmt.Println("rounds done:", i+1)
		// printArray(messages)
		// fmt.Println("decisions...")
		// printArray(decisionTracker)
	}
	return decisionTracker
}

// randNums chooses num random numbers in range (0, n] and returns a dict of the chosen numbers
func randNums(length int, num int, studentPool map[int]Student) map[int]bool {
	randNums := make(map[int]bool)
	i := 0

	for i < num {
		newNum := rand.Intn(length)
		//fmt.Println(newNum)
		_, exists := studentPool[newNum]
		if randNums[newNum] || exists == false { //number has already been added to dropList, or this was dropped in previous rounds
			continue
		}
		randNums[newNum] = true
		i++
	}

	return randNums
}

func writeMessageS(j int, studentPool map[int]Student, messages [][]int) {

	// student will drop out in this round at some point
	if studentPool[j].willdrop {

		//at what time will this student drop?
		dropIndex := rand.Intn(len(messages))

		for i := range messages[j] {
			if i >= dropIndex {
				messages[j][i] = 0
				delete(studentPool, j)
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

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
