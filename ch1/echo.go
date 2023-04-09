package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(echo1(os.Args), " Elasped Time(s): ", time.Since(start).Seconds())
	start = time.Now()
	fmt.Println(echo2(os.Args), " Elasped Time(s): ", time.Since(start).Seconds())
	start = time.Now()
	fmt.Println(echo3(os.Args), " Elasped Time(s): ", time.Since(start).Seconds())
	start = time.Now()
	fmt.Println(echo(os.Args), " Elasped Time(s): ", time.Since(start).Seconds())
	start = time.Now()
	fmt.Println(echoExercise(os.Args), " Elasped Time(s): ", time.Since(start).Seconds())
	start = time.Now()
}

// Echo Implementation 1
func echo1(args []string) string {
	var res, sep string
	for i := 1; i < len(args); i++ {
		res += sep + args[i]
		sep = " "
	}
	return res
}

// Echo Implementation 2
func echo2(args []string) string {
	var res, sep string
	for _, arg := range args[1:] {
		res += sep + arg
		sep = " "
	}
	return res
}

// Echo Implementation 3
func echo3(args []string) string {
	return strings.Join(args[1:], " ")
}

// My Echo Implementation
func echo(args []string) string {
	res := strings.Builder{}
	sep := " "
	for i, arg := range args[1:] {
		res.WriteString(arg)
		if i < len(args)-1 {
			res.WriteString(sep)
		}
	}
	return res.String()
}

// Echo Exercise
func echoExercise(args []string) string {
	res := strings.Builder{}
	for i, arg := range args {
		res.WriteString(fmt.Sprintf("%d %s \n", i, arg))
	}
	return res.String()
}
