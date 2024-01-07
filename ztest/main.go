package main

import (
	"fmt"
	"log"
	"time"
)

type obj struct {
	threadRuning   bool
	dataChannel    chan int
	stopperChannel chan string
	endProgram     chan int
}

func main() {

	myobj := &obj{}
	myobj.dataChannel = make(chan int, 1)
	myobj.stopperChannel = make(chan string, 1)
	myobj.endProgram = make(chan int, 1)

	go func() {
		count := 0
		myobj.threadRuning = true
		localStop := false
		log.Printf("Listner-Thread : Start")

		defer func() {
			myobj.endProgram <- 1
		}()

		for myobj.threadRuning && !localStop {
			//Do some work and continue the next steps.
			count = count + 1
			time.Sleep(time.Duration(1000) * time.Millisecond)
			if myobj.threadRuning && (count == 10) {
				//Post the result back to master thread.
				log.Println("Listner :  will stop since count is 10")
				myobj.dataChannel <- count
				myobj.stopperChannel <- "stop"
				localStop = true

			} else if myobj.threadRuning {
				myobj.dataChannel <- count
			} else {
				log.Println("Listner :  will stop since master has already ended.")
				localStop = true
			}
		}
		log.Println("Listner - End")
	}()

	go func() {
		stop := 1

		for stop != 0 {
			fmt.Println("Enter digit 0 to trigger stopper thread")
			fmt.Scanf("%d", &stop)
		}

		if myobj.threadRuning {
			log.Println("Stopper :  stop the master thread")
			myobj.stopperChannel <- "stop"
		} else {
			log.Println("Stopper :  Thread was alreay halted")
		}

	}()

	for {
		select {
		case <-myobj.stopperChannel:
			log.Printf("Master - thread : stop")
			myobj.threadRuning = false
			close(myobj.stopperChannel)
			close(myobj.dataChannel)
		case data := <-myobj.dataChannel:
			log.Println("Master - thread : data from listener (", data, ")")
			continue
		}
		if !myobj.threadRuning {
			break
		}
	}

	<-myobj.endProgram
	log.Println("Will end program now")
}
