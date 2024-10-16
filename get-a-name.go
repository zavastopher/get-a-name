package main

import (
	"bufio"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	nameDir := "name-files/counted"
	fileSystem := os.DirFS(nameDir)
	var nameFiles []string

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		nameFiles = append(nameFiles, path)

		return nil
	})

	random_file := nameFiles[rand.Intn(len(nameFiles))]

	file, err := os.Open(nameDir + "/" + random_file)
	check(err)

	name_scanner := bufio.NewScanner(file)
	name_scanner.Scan()
	count, err := strconv.Atoi(name_scanner.Text())
	check(err)
	i := rand.Intn(count)
	for name_scanner.Scan() && i > 0 {
		i--
	}
	print(name_scanner.Text())
}
