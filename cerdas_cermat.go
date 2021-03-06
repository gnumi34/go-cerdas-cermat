package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func watcher(timeout int, quit <-chan bool) {
	select {
	case <-time.After(time.Duration(timeout) * time.Second):
		fmt.Println("Time is up! You have failed!")
		os.Exit(0)
	case <-quit:
		return
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var quitWatcher = make(chan bool)
	var timeout int = 3
	var points int = 0
	var readyAns string
	var answer string

	fmt.Println("Welcome to Cerdas Cermat bersama Bioviton!")
	time.Sleep(2 * time.Second)
	fmt.Println("Answer each question correctly and win 5 Million Rupiah!")
	time.Sleep(2 * time.Second)
	fmt.Println("You have 3 seconds to answer each question, or else you lose!")
	time.Sleep(2 * time.Second)
	fmt.Println("Are you ready? (type Y or N)")
	fmt.Scanln(&readyAns)

	for readyAns != "Y" && readyAns != "y" {
		fmt.Println("Now, are you ready? (type Y or N) ")
		fmt.Scan(&readyAns)
	}

	fmt.Println("Let's begin!")
	fmt.Println("-------------------------------------------------")
	time.Sleep(2 * time.Second)

	go watcher(timeout, quitWatcher)
	fmt.Println("Question 1 of 3")
	fmt.Printf("Point : %d\n", points)
	fmt.Println("Who is the first president of Indonesia? ")
	answer, _ = in.ReadString('\n')

	if answer == "Ir. Soekarno\n" || answer == "Ir. Sukarno\n" || answer == "Soekarno\n" {
		quitWatcher <- true
		fmt.Println("Correct Answer!")
		points += 1
	} else {
		quitWatcher <- true
		fmt.Println("Booo wrong answer!!")
	}

	fmt.Println("-------------------------------------------------")
	time.Sleep(2 * time.Second)

	go watcher(timeout, quitWatcher)
	fmt.Println("Question 2 of 3")
	fmt.Printf("Point : %d\n", points)
	fmt.Println("Who is the inventor of Go? ")
	fmt.Scanln(&answer)

	if answer == "China" {
		quitWatcher <- true
		fmt.Println("Correct Answer!")
		points += 1
	} else {
		quitWatcher <- true
		fmt.Println("Booo wrong answer!!")
	}

	fmt.Println("-------------------------------------------------")
	time.Sleep(2 * time.Second)

	go watcher(timeout, quitWatcher)
	fmt.Println("Question 3 of 3")
	fmt.Printf("Point : %d\n", points)
	fmt.Println("Who is the perpetrator of John F. Kennedy assasination? ")
	fmt.Scanln(&answer)

	if answer != "" {
		quitWatcher <- true
		fmt.Println("Correct Answer!")
		points += 1
	} else {
		quitWatcher <- true
		fmt.Println("Booo wrong answer!!")
	}

	fmt.Println("-------------------------------------------------")
	time.Sleep(2 * time.Second)

	fmt.Println("Your point is :", points)
	if points == 3 {
		fmt.Println("Congratulations! You win the price! You really are a winner!")
	} else {
		fmt.Println("We are sorry that you have failed. Try again in the next Cerdas Cermat with Bioviton!")
	}
}
