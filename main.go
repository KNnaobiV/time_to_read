package main

import (
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	filepath := "C:/Users/Ekenair/Documents/Code_Files/Go/example.txt"
	text := textReader(filepath)
	wordCount, timeToRead := calculateTimeToRead(text, 200)
	log.Printf(
		"Estimated time to read %s with word count %v: %v", 
		filepath, wordCount, &timeToRead,
	)
}

func calculateTimeToRead(text string, averageWordCountPerMinute int) (int, time.Duration) {
	punctuationMarks := []string{"!", ",", ".", ";", ":", "?", "-", "_"}
	wordsInString := strings.Fields(text)
	wordCount := 0
	for _, word := range wordsInString {
		for _, mark := range punctuationMarks {
			if word == mark {
				continue
			}
		}
		wordCount += 1
	}
	durationAsInt := wordCount / averageWordCountPerMinute
	duration := time.Duration(durationAsInt) * time.Minute
	return wordCount, duration
}

func textReader(fileLocation string) string {
	_, err := os.Stat(fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	text, err := os.ReadFile(fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	return string(text)
}
