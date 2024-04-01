package main

import (
	"keyguru/internal/exercise"
	"keyguru/internal/typing"
	"keyguru/pkg/ui"
	"strings"
)

func main() {
	// Create a new UI
	userInterface := ui.NewUI()

	// Display a welcome message
	userInterface.DisplayMessage("Welcome to KeyGuru! Let's improve your typing skills.")

	// Create an exercise store and load exercises
	store := exercise.NewExerciseStore()
	exercises, err := exercise.LoadAllExercisesFromDirectory("ressources")
	if err != nil {
		userInterface.DisplayMessage("Error loading exercises: " + err.Error())
		return
	}
	for _, ex := range exercises {
		store.AddExercise(ex)
	}

	// Ask the user to choose a difficulty
	userInterface.DisplayMessage("Please choose a difficulty level (1-3):")
	difficulty, err := userInterface.ReadIntInput()
	if err != nil {
		userInterface.DisplayMessage("Invalid difficulty level. Please enter a number between 1 and 3.")
		return
	}

	// Start a new typing session
	session, err := typing.StartSession(store, difficulty)
	if err != nil {
		userInterface.DisplayMessage("Error starting session: " + err.Error())
		return
	}

	// Display the exercise to the user
	userInterface.DisplayMessage("Start typing the following exercise:")
	userInterface.DisplayMessage(session.Exercise.Content)

	// Split the exercise content into words
	exerciseWords := strings.Split(session.Exercise.Content, " ")

	// Create a slice to hold the words the user has typed
	typedWords := make([]string, 0, len(exerciseWords))

	// Initialize an error counter
	errorCount := 0

	// Read the user's input
	for {
		input, err := userInterface.ReadInput()
		if err != nil {
			userInterface.DisplayMessage("Error reading input: " + err.Error())
			return
		}

		// Add the words the user has typed to the slice
		typedWords = append(typedWords, strings.Split(input, " ")...)

		// Record the user's typing
		session.RecordTyping(input)

		// Check each typed word against the corresponding word in the exercise
		for i, word := range typedWords {
			if i < len(exerciseWords) && word != exerciseWords[i] {
				errorCount++
			}
		}

		// If the user has typed all the words in the exercise, end the session
		if len(typedWords) >= len(exerciseWords) {
			session.EndSession()
			break
		}
	}

	// Display the session results
	userInterface.DisplayResults(session.Results.Wpm, session.Results.Accuracy, errorCount, session.Results.TotalTime)
}
