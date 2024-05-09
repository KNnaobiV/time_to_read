package main

import "github.com/knnaobiv/time_to_read/cmd"


func main() {
	cmd.Execute()
}

// func main() {
// 	filepath := "C:/Users/Ekenair/Documents/Code_Files/Go/example.txt"
// 	text := textReader(filepath)
// 	wordCount, timeToRead := calculateTimeToRead(text, 200)
// 	log.Printf(
// 		"Estimated time to read %s with word count %v: %v", 
// 		filepath, wordCount, &timeToRead,
// 	)
// }
