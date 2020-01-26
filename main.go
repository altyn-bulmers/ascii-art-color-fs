package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func main() {
	arg := os.Args
	arguments := []string(arg[1:2])
	colorArg := arg[2:]
	str := ""
	lenArg := 0
	for range arguments {
		lenArg++
	}
	for i, input := range arguments {
		if i == lenArg-1 {
			str += input
			break
		}
		str += input + " "
	}

	fileName := "standard.txt"

	var colorRunes []rune
	Paint := White
	for _, arg := range colorArg {
		length := 0
		for j := range arg {
			length = j + 1
		}
		if length > 8 && arg[:8] == "--color=" {
			runes_temp := []rune(arg[8:])

			for i := 0; i < length-8; i++ {
				colorRunes = append(colorRunes, runes_temp[i])
			}
			if string(colorRunes) == "white" {
				Paint = White
			} else if string(colorRunes) == "teal" {
				Paint = Teal
			} else if string(colorRunes) == "magenta" {
				Paint = Magenta
			} else if string(colorRunes) == "purple" {
				Paint = Purple
			} else if string(colorRunes) == "yellow" {
				Paint = Yellow
			} else if string(colorRunes) == "green" {
				Paint = Green
			} else if string(colorRunes) == "red" {
				Paint = Red
			} else if string(colorRunes) == "black" {
				Paint = Black
			}

		

		} else if arg == "standard" {
			fileName = "standard.txt"
		} else if arg == "shadow" {
			fileName = "shadow.txt"
		} else if arg == "thinkertoy" {
			fileName = "thinkertoy.txt"
		}

	}

	size := 0
	for i, b := range []byte(str) {
		if b == 92 && i < len([]byte(str))-1 && []byte(str)[i+1] == 110 {
			size++
		}
	}
	splittwo := string(byte(92)) + string(byte(110))
	words := strings.Split(str, splittwo)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer file.Close()

	rawBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(rawBytes), "\n")

	for _, word := range words {
		for h := 0; h < 9; h++ {
			for _, l := range []byte(word) {
				for i, line := range lines {
					if i == (int(l)-32)*9+h {
						fmt.Print(Paint(line))
					}
				}
			}
			fmt.Println()
		}
	}

}
