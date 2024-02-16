package main

import (
	"os"
)

func main() {
	// if we want to design a package, a great way to begin is by pretending it already exists, and
	//writing code that uses it to solve our problem.
	// os.Stdout,  a file handle representing the standard output.
	//  determines where stuff ends up getting printed.
	// it's okay but Small inconveniences like this to supply generic arguments like os.stdout
	// add up, over a big enough API it becomes a hindernce
	// we can make os.stdout default if argument passed is nil but this makes it worse
	// as it does not tell users why nil is passed.
	hello.PrintTo(os.Stdout)
	// a much better method
	hello.Main()
}
