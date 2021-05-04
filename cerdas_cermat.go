package main

import (
	"fmt"
	"os"
	"time"
)

func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool, quit <-chan bool) {
	for {
		select {
		case <-ch:
			fmt.Println("Time is up! You have failed!")
			os.Exit(0)
		case <-quit:
			return
		}
	}
}

func main() {
	var ch = make(chan bool)
	var quit = make(chan bool)
	var timeout int = 5
	var points int = 0
	var readyAns string
	var answer string

	fmt.Println("Welcome to Cerdas Cermat bersama Bioviton!")
	time.Sleep(2 * time.Second)
	fmt.Println("Answer each question correctly and win 5 Million Rupiah!")
	time.Sleep(2 * time.Second)
	fmt.Println("You have 15 seconds to answer each question, or else you lose!")
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

	go timer(timeout, ch)
	go watcher(timeout, ch, quit)
	fmt.Println("Question 1 of 3")
	fmt.Printf("Point : %d\n", points)
	fmt.Println("Who is the first president of Indonesia? ")
	fmt.Scan(&answer)

	if answer == "Ir. Soekarno" || answer == "Ir. Sukarno" || answer == "Soekarno" {
		quit <- true
		fmt.Println("Correct Answer!")
		points += 1
	} else {
		fmt.Println("Booo wrong answer!!")
	}

	fmt.Println("-------------------------------------------------")
	time.Sleep(2 * time.Second)

	go timer(timeout, ch)
	go watcher(timeout, ch, quit)
	fmt.Println("Question 2 of 3")
	fmt.Printf("Point : %d\n", points)
	fmt.Println("Who is the inventor of Go? ")
	fmt.Scan(&answer)

	if answer == "China" {
		quit <- true
		fmt.Println("Correct Answer!")
		points += 1
	} else {
		fmt.Println("Booo wrong answer!!")
	}

	fmt.Println("-------------------------------------------------")
	time.Sleep(2 * time.Second)

	go timer(timeout, ch)
	go watcher(timeout, ch, quit)
	fmt.Println("Question 1 of 3")
	fmt.Printf("Point : %d\n", points)
	fmt.Println("Who is the perpetrator of John F. Kennedy assasination? ")
	fmt.Scan(&answer)

	if answer != "" {
		quit <- true
		fmt.Println("Correct Answer!")
		points += 1
	} else {
		fmt.Println("Booo wrong answer!!")
	}

	fmt.Println("-------------------------------------------------")
	time.Sleep(2 * time.Second)

	fmt.Println("Your point is ", points)
	if points == 3 {
		fmt.Println("Congratulations! You win the price! You really are a winner!")
	} else {
		fmt.Println("We are sorry that you have failed. Try again in the next Cerdas Cermat with Bioviton!")
	}
}
