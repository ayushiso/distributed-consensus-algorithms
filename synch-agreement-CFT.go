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
	willdrop bool
}

func Agreement(n, t int) {

	// recieve and send messages through this array
	messages := make([][]int, n)
	for i := range messages {
		messages[i] = make([]int, n)
	}

	studentPool := make(map[int]Student)

	for i := 0; i < n; i++ {
		var student Student
		student.message = rand.Intn(100) + 1
		student.willdrop = false
		studentPool[i] = student
	}

	for i := range studentPool {
		//fmt.Println("HI")
		fmt.Println(studentPool[i])
	}

	for i := 0; i <= t; i++ {

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
		fmt.Println("numDrops=", numDrops)

		// choose which students are going to drop

		dropList := randNums(n, numDrops, studentPool)

		// writing step (GOROUTINE?)

		for j := range studentPool {

			if len(studentPool) == 0 {
				break
			}

			if dropList[j] {
				student := studentPool[j]
				student.willdrop = true
				studentPool[j] = student
			}

			writeMessage(j, studentPool, messages)
		}

		// choosing step (GOROUTINE?)

		for j := range studentPool {
			//fmt.Println(studentPool[j].drop)
			student := studentPool[j]
			student.message = chooseMessage(j, messages)
			studentPool[j] = student
		}

		printArray(messages)
		fmt.Println("rounds done:", i+1)
	}
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

func writeMessage(j int, studentPool map[int]Student, messages [][]int) {

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

func printArray(arr [][]int) {
	for elems := range arr {
		fmt.Println(arr[elems])
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
