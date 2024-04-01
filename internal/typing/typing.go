package typing

import (
	"keyguru/internal/exercise"
	"time"
)

type Session struct {
	Exercise  *exercise.Exercise // Exercice en cours
	StartTime time.Time          // L'heure de démarrage de la session
	EndTime   time.Time          // L'heure de fin de la session
	Results   *Results           // Les résultats de la session comme la vitesse, nombre d'erreurs, etc.
}

type Results struct {
	Wpm         float64 // Mots par minute
	Accuracy    float64 // Précision en pourcentage
	Errors      int     // Nombre d'erreurs
	TotalTime   float64 // Temps total en secondes
	TotalTyping int     // Nombre total de frappes
}

func StartSession(store *exercise.ExerciseStore, difficulty int) (*Session, error) {
	ex, err := store.GetExerciseByDifficulty(difficulty)
	if err != nil {
		return nil, err
	}

	return &Session{
		Exercise:  ex,
		StartTime: time.Now(),
		Results:   &Results{},
	}, nil
}

func (s *Session) RecordTyping(userInput string) {
	totalTyping := len(userInput)
	errors := calculateErrors(userInput, s.Exercise.Content)

	s.Results.TotalTyping += totalTyping
	s.Results.Errors += errors
}

func calculateErrors(userInput, exerciseContent string) int {
	errors := 0
	for i := 0; i < len(userInput) && i < len(exerciseContent); i++ {
		if userInput[i] != exerciseContent[i] {
			errors++
		}
	}

	return errors
}

func (s *Session) EndSession() {
	s.EndTime = time.Now() // End the session timer

	// Calculate total time
	s.Results.TotalTime = float64(s.EndTime.Sub(s.StartTime)) / float64(time.Second)

	// Calculate Words Per Minute (WPM)
	s.Results.Wpm = float64(s.Results.TotalTyping) / 5 / (s.Results.TotalTime / 60)

	// Calculate Accuracy
	totalPossibleErrors := len(s.Exercise.Content)
	if totalPossibleErrors > 0 {
		s.Results.Accuracy = 1 - float64(s.Results.Errors)/float64(totalPossibleErrors)
	}
}
