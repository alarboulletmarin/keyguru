package exercise

import (
	"os"
	"path/filepath"
	"strings"
)

// Exercise represents a typing exercise with its content, difficulty, and name.
type Exercise struct {
	Content    string // The text of the typing exercise
	Difficulty int    // A way to categorize the difficulty of the exercise (optional)
	Name       string // The title or name of the exercise, for reference
}

func LoadExerciseFromFile(filename string) (*Exercise, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	normContent := normalizeContent(string(content))
	difficulty := defineDifficulty(normContent)
	exercise := &Exercise{
		Content:    normContent,
		Difficulty: difficulty,
		Name:       filepath.Base(filename), // use base filename
	}
	return exercise, nil
}

// normalizeContent replaces newline characters with spaces in a string
func normalizeContent(content string) string {
	return strings.ReplaceAll(content, "\n", " ")
}

func LoadAllExercisesFromDirectory(directory string) ([]*Exercise, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var exercises []*Exercise
	for _, file := range files {
		if !file.IsDir() {
			exercise, err := LoadExerciseFromFile(directory + "/" + file.Name())
			if err != nil {
				return nil, err
			}
			exercises = append(exercises, exercise)
		}
	}

	return exercises, nil
}

// Define the difficulty based on the length of the content
func defineDifficulty(content string) int {
	length := len(content)
	switch {
	case length <= 100:
		return 1
	case length <= 200:
		return 2
	default:
		return 3
	}
}
