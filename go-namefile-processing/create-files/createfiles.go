package main

import (
	"bufio"
	"fmt"
	//"io"
	"os"
	"log"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func buildFileName(name string) string{
	newname := "../name-files/"
	for _, ch := range name {
		if ch != ',' {
			if ch == ' '{
				newname += string("-")
			} else {
				newname += string(ch)
			}
		}
	}
	newname += ".txt"
	return newname
}

func main() {
	file, err := os.Open("../name-files/name-categories.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		fmt.Println(str)
		err := os.WriteFile(buildFileName(str), []byte(""), 0644)
		check(err)
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
