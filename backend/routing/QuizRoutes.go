package routing

import (
	"encoding/json"
	"net/http"
)

type apiResponse struct {
	Data interface{} `json:"data"`
}

const (
	multipleCoiceKey = "multiple-coice"
	freeTextKey      = "free-text"
)

type question struct {
	ID string `json:"id"`
	// multiple-choice | freetext
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// single answer -> radio button
type multipleCoiceData struct {
	AllowMultiple bool     `json:"allowMultiple"`
	Question      string   `json:"question"`
	Answers       []answer `json:"answers"`
}

// full test answers -> text input
type freeTextData struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type answer struct {
	Text      string `json:"text"`
	IsCorrect bool   `json:"isCorrect"`
}

func exampleQuizzes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(
		&apiResponse{
			// example data
			Data: []question{
				{
					ID:   "1",
					Type: multipleCoiceKey,
					Data: multipleCoiceData{
						Question: "What is love?",
						Answers: []answer{
							{
								Text:      "Baby don't hurt me",
								IsCorrect: true,
							},
							{
								Text:      "b",
								IsCorrect: false,
							},
							{
								Text:      "c",
								IsCorrect: false,
							},
						},
					},
				},
				{
					ID:   "2",
					Type: multipleCoiceKey,
					Data: multipleCoiceData{
						Question:      "Who's the boss?",
						AllowMultiple: true,
						Answers: []answer{
							{
								Text:      "Huba",
								IsCorrect: true,
							},
							{
								Text:      "Huba",
								IsCorrect: true,
							},
							{
								Text:      "Huba",
								IsCorrect: true,
							},
						},
					},
				},
				{
					ID:   "2",
					Type: freeTextKey,
					Data: freeTextData{
						Question: "How are you doing",
						Answer:   "mighty fine",
					},
				},
			},
		},
	)
}
