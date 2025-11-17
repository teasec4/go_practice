package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

func main(){

	if len(os.Args) < 2 {
		printHelp()
		return
	}

	if len(os.Args) < 4 {
		fmt.Println("Usage go run . <add|sub|sum|div> x y")
		return
	}

	cmd := os.Args[1]

	switch cmd{
		case "help":
			printHelp()
		case "add", "sub", "mul", "div", "mod":
			executeMath(cmd)
		default:
			fmt.Println("Unknown command:", cmd)
			printHelp()
	}
}

// math func
func executeMath(cmd string){
	if len(os.Args) < 4 {
		fmt.Println("Not enough arguments!")
		printHelp()
		return
	}

	x, err1 := strconv.Atoi(os.Args[2])
	y, err2 := strconv.Atoi(os.Args[3])

	
	if err1 != nil || err2 != nil{
		fmt.Println("Error:", err1, err2)
		return
	}

	result, err := calc(cmd, x, y)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
	
}

// calc
func calc(cmd string, x, y int) (int, error){

	switch cmd{
		case "add":
			return x + y, nil
		case "sub":
			return x - y, nil
		case "mul":
			return x * y, nil
		case "div":
			return x / y, nil
		case "mod":
			if y == 0 {
				return 0, errors.New("modulo by zero")
			}
			return x % y, nil
		default:
			return 0, errors.New("Unknown command")
	}
	
}


// help func 
func printHelp() {
	fmt.Println(`
		Usage:
		go run . <command> <x> <y>

		Commands:
		add  x y    → x + y
		sub  x y    → x - y
		mul  x y    → x * y
		div  x y    → x / y
		mod  x y    → остаток от деления
		help        → показать это сообщение

		Examples:
		go run . add 5 10
		go run . div 20 4
		go run . mod 10 3
	`)
}