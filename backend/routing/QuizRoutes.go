package routing

import (
	"encoding/json"
	"net/http"

	"github.com/fourtf/studyhub/utils"
)

type apiResponse struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

type quiz struct {
	ID string `json:"id"`
	// single | multi | freetext
	Type     string   `json:"type"`
	Question string   `json:"question"`
	Answer   *string  `json:"answer,omitifempty"`
	Answers  []answer `json:"answers,omitifempty"`
}

type answer struct {
	Text      string `json:"text"`
	IsCorrect bool   `json:"isCorrect"`
}

func exampleQuizzes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(
		&apiResponse{
			Error: "",
			// example data
			Data: []quiz{
				{
					ID:       "1",
					Type:     "single",
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
				{
					ID:       "2",
					Type:     "multi",
					Question: "Who's the boss?",
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
				{
					ID:       "2",
					Type:     "freetext",
					Question: "How are you doing",
					Answer:   utils.MakeStrPtr("mighty fine"),
				},
			},
		},
	)
}
