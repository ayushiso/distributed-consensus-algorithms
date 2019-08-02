package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Process struct {
	pref  int
	round int
}

type Message struct {
	round  int
	pref   int
	ratify bool
}

func SimulateBenOr(n, t int) {
	processes := initProcesses(n)
	messages := make(chan Message)
	ratifier := make(chan Message)
	var wg sync.WaitGroup

	wg.Add(n)
	for i := range processes {
		go simulateOneProc(n, t, processes[i], messages, ratifier, &wg)
	}
	wg.Wait()
}

func simulateOneProc(n, t int, process Process, messages chan Message, ratifier chan Message, wg *sync.WaitGroup) {
	for true {
		fmt.Println("starting round", process.round)
		var message1 Message
		message1.pref = process.pref
		message1.round = process.round

		// send this process's preference and round
		messages <- message1

		incoming := make([]Message, 0)

		// wait for n-t messages
		for i := 1; i <= n-t; i++ {
			newMessage := <-messages

			// discard messages from older rounds
			if newMessage.round >= process.round {
				incoming = append(incoming, newMessage)
			}
		}

		// check for n/2 majority of a particular message
		value, count := majority2(incoming)

		var message2 Message
		message2.round = process.round

		if count > n/2 {
			message2.ratify = true
			message2.pref = value
		} else {
			message2.ratify = false
		}
		ratifier <- message2

		// wait for n-t messages
		check := 0
		for i := 1; i <= n-t; i++ {
			newMessage := <-messages

			// discard messages from older rounds
			if newMessage.round >= process.round {
				if newMessage.ratify {
					process.pref = newMessage.pref
					check++
				}
			}
		}
		if check > t {
			break
		}
		process.round++
	}
	wg.Done()
}

func initProcesses(n int) []Process {
	processes := make([]Process, n)
	for i := range processes {
		processes[i].pref = rand.Intn(2) + 1
		processes[i].round = 1
	}
	return processes
}
