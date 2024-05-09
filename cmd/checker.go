package cmd

import (
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

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
		color.Red(err.Error())
	}
	text, err := os.ReadFile(fileLocation)
	if err != nil {
		color.Red(err.Error())
	}
	return string(text)
}

var readTimeCmd = &cobra.Command{
	Use: "read-time",
	Aliases: []string{"r"},
	Short: "Checks read time for a text",
	Long: `Checks the time it will take to read a text file based on the 
			reading speed passed to it. Will use 200wpm as reading speed 
			if nothing is passed.`,
	Run: func(cmd *cobra.Command, args[]string) {
		if len(args) == 0 {
			color.Red("You must include the file to be read.")
			return
		}
		filePath := args[0]
		text := textReader(filePath)
		wordCount, timeToRead := calculateTimeToRead(text, 200)
		color.Green(
			"Estimated time to read %s with word count %v: %v", 
			filePath, wordCount, &timeToRead,
		)
	},
}

func init() {
	rootCmd.AddCommand(readTimeCmd)
}