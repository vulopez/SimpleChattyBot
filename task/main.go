package main

import (
	. "fmt"
	"strconv"
)

func greet(name, year string) {
	printPar("Hello! My name is "+name+".",
		"I was created in "+year+".")
}

func showName() {
	var name string
	Println("Please, remind me your name.")
	Scan(&name)
	Println("What a great name you have, " + name + "!")
}

func guessAge() {
	var rem3, rem5, rem7, age int

	printPar("Let me guess your age.",
		"Enter remainders of dividing your age by 3, 5 and 7.")

	_, err := Scan(&rem3, &rem5, &rem7)
	if err != nil {
		return
	}

	age = (rem3*70 + rem5*21 + rem7*15) % 105
	Println("Your age is " + strconv.Itoa(age) + "; that's a good time to start programming!")
}

func count() {
	var n int

	Println("Now I will prove to you that I can count to any number you want.")
	Scan(&n)
	for i := 0; i <= n; i++ {
		Printf("%d!\n", i)
	}
}

func startQuiz() {
	var answer int
	printPar("Let's test your programming knowledge.",
		"Why do we use methods?",
		"1. To repeat a statement multiple times.",
		"2. To decompose a program into several small subroutines.",
		"3. To determine the execution time of a program.",
		"4. To interrupt the execution of a program.",
	)
	for answer != 2 {
		_, _ = Scan(&answer)
		if answer != 2 {
			Println("Please, try again.")
		}
	}
}

func sayGoodbye() {
	Println("Congratulations, have a nice day!")
}

func main() {
	greet("Aid", "2023") // change it as you need
	showName()
	guessAge()
	count()
	startQuiz()
	sayGoodbye()
}

func printPar(a ...string) { // Print each argument in a line
	for i := range a {
		Println(a[i])
	}
}
