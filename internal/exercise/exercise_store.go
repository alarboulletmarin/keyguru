package exercise

import (
	"errors"
	"math/rand"
	"time"
)

const ErrNoExercises = "no exercises found with the specified difficulty"

type ExerciseStore struct {
	Exercises []*Exercise
}

func NewExerciseStore() *ExerciseStore {
	return &ExerciseStore{
		Exercises: make([]*Exercise, 0),
	}
}

func (es *ExerciseStore) AddExercise(ex *Exercise) {
	es.Exercises = append(es.Exercises, ex)
}

func (es *ExerciseStore) GetExerciseByDifficulty(difficulty int) (*Exercise, error) {
	var exerciseOptions []*Exercise
	for _, ex := range es.Exercises {
		if ex.Difficulty == difficulty {
			exerciseOptions = append(exerciseOptions, ex)
		}
	}
	if len(exerciseOptions) == 0 {
		return nil, errors.New(ErrNoExercises)
	}
	rand.Seed(time.Now().UnixNano())
	return exerciseOptions[rand.Intn(len(exerciseOptions))], nil
}
