package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"learningGo/Gophercises/cyoa"
)

// go build builds a binary and name it with the name of the directory
// Write a working solution
// Refactor (create functions, new structs to improve readability)
func main() {
	filePath := "/Users/ronald/Documents/Online-Courses/LearningGo/Gophercises/cyoa/gopher.json"
	fileName := flag.String("file", filePath, "the JSON file with CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s. \n", *fileName)
	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	story, err := cyoa.JsonStory(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", story)
}
