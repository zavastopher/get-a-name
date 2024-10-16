package main

import (
	"bufio"
	"io/fs"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	nameDir := "name-files"
	fileSystem := os.DirFS(nameDir)
	//var nameFiles []string

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if path != "name-categories.txt" && path != "." && path != "counted" && path[:7] != "counted" {
			//nameFiles = append(nameFiles, path)
			print(path)
			srcFile := nameDir + "/" + path
			srcTxt, err := os.ReadFile(srcFile)
			check(err)
			os.Stdout.Write(srcTxt)
			file, err := os.Open(srcFile)
			check(err)
			scanner := bufio.NewScanner(file)
			count := 0
			for scanner.Scan() {
				count++
			}
			err = file.Close()
			check(err)

			countedFile := nameDir + "/counted/" + path
			err = os.WriteFile(countedFile, []byte(strconv.Itoa(count)+"\n"), 0666)
			check(err)
			f, err := os.OpenFile(countedFile, os.O_APPEND, 0666)
			check(err)
			_, err = f.Write(srcTxt)
			check(err)
			// os.Stdout.Write(srcTxt)

		}

		return nil
	})

	// random_file := nameFiles[rand.Intn(len(nameFiles))]

	// names = os.ReadFile(random_file(nameDir + "/" + random_file))
	// i = rand.Int()
	// len = 0
	// while()
}
