package main

import (
	//"flag"
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main(){
	OminoTest()
}

func OminoTest(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Tests")
	tests, _ := reader.ReadString('\n')
	// flag.Parse()
	// tests1 := flag.Arg(0)
	number, _ := strconv.ParseInt(tests, 10, 8)
	
	i := 0

	fmt.Println(number)
	for i = 0; i < number; i++{
		test(i+1)
	}
}

func test(numberOfCases int){

	fmt.Println("debugging")
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Println("Enter X value:")
	xValue, _ := reader2.ReadString('\n')
	X, _ := strconv.ParseInt(xValue, 10, 8)

	reader3 := bufio.NewReader(os.Stdin)
	fmt.Println("Enter number of rows:")
	rValue, _ := reader3.ReadString('\n')
	R, _ := strconv.ParseInt(rValue, 10, 8)

	reader4 := bufio.NewReader(os.Stdin)
	fmt.Println("Enter number of columns:")
	cValue, _ := reader4.ReadString('\n')
	C, _ := strconv.ParseInt(cValue, 10, 8)

	OminoGame(X, R, C, numberOfCases)
}

func OminoGame(X int, R int, C int, numberOfCases int){
	firstIndicator := 7
	secondIndicator := ((R * C) % X)
	thirdIndicator := X
	var fourthIndicator int

	if X%2 == 0{
		fourthIndicator = X / 2
	}else{
		fourthIndicator = (X / 2) + 1
	}

	winner := "Richard"

	if X < firstIndicator {
		if secondIndicator == 0 {
			if R >= thirdIndicator || C >= thirdIndicator {
				if R >= thirdIndicator {
					if C >= fourthIndicator {
						winner = "Gabriel"
					}
				}
				if C >= thirdIndicator {
					if R >= fourthIndicator {
						winner = "Gabriel"
					}
				}
			}
		}
	}
	fmt.Println(winner)
}