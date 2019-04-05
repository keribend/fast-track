package models

import (
	"fmt"
	"math/rand"
	"time"
)

type Answer struct {
	ID      uint   `json:"id"`
	Text    string `json:"text"`
	correct bool
}

type Question struct {
	ID      uint   `json:"id"`
	Text    string `json:"text"`
	Answers []Answer
}

type Questionnaire struct {
	Questions []Question `json:"items"`
}

type QuestionnaireAnswers struct {
	SelectedAnswers map[uint]uint `json:"selectedAnswers"`
}

type QuestionnaireResult struct {
	AnsweredQuestions   int     `json:"answeredQuestions"`
	CorrectAnswers      int     `json:"correctAnswers"`
	CurrentSuccessRatio float32 `json:"currentSuccessRatio"`
	GlobalSuccessRatio  float32 `json:"globalSuccessRatio"`
}

// mock types are used to mock the data caming from a generic datasource
type mockQuestion Question
type mockStatistics struct {
	numOfExecutions     int
	totalCorrectAnswers int
	totalQuestions      int
	successRatio        float32
}

var mockData []mockQuestion
var mockStats mockStatistics
var rnd *rand.Rand

func GetQuestionnaire(nQuestions int) Questionnaire {
	result := Questionnaire{}

	questions := map[uint]Question{}
	for len(questions) < nQuestions {
		question := getQuestion(4)
		if _, exist := questions[question.ID]; !exist {
			questions[question.ID] = question
		}
	}

	for _, question := range questions {
		result.Questions = append(result.Questions, question)
	}

	return result
}

func AnswerQuestionnaire(input *QuestionnaireAnswers) (QuestionnaireResult, error) {
	result := QuestionnaireResult{}
	correctAnswCounter, err := countCorrectAnswers(*input)
	if err != nil {
		return result, err
	}

	nOfQuestions := len(input.SelectedAnswers)

	mockStats.updateStats(nOfQuestions, correctAnswCounter)
	currentSuccessRatio := evalSuccessRatio(nOfQuestions, correctAnswCounter)

	return QuestionnaireResult{
		AnsweredQuestions:   nOfQuestions,
		CorrectAnswers:      correctAnswCounter,
		CurrentSuccessRatio: currentSuccessRatio,
		GlobalSuccessRatio:  mockStats.successRatio,
	}, nil
}

func evalSuccessRatio(nOfQuestions, correctAnswCounter int) float32 {
	return float32(correctAnswCounter) / float32(nOfQuestions)
}

func (ms *mockStatistics) updateStats(nOfQuestions, correctAnswCounter int) {
	ms.numOfExecutions++
	ms.totalCorrectAnswers += correctAnswCounter
	ms.totalQuestions += nOfQuestions
	ms.successRatio = float32(ms.totalCorrectAnswers) / float32(ms.totalQuestions)
}

func countCorrectAnswers(input QuestionnaireAnswers) (int, error) {
	count := 0
	for qID, aID := range input.SelectedAnswers {
		question, err := findMockQuestion(qID)
		if err != nil {
			return count, err
		}
		correctAnsw := question.getCorrectAnswer()
		if correctAnsw.ID == aID {
			count++
		}
	}
	return count, nil
}

func getQuestion(nAnsw int) Question {
	result := Question{}

	// get a random question
	mockQuestion := mockData[rnd.Intn(len(mockData))]

	// get question attributes
	result.ID = mockQuestion.ID
	result.Text = mockQuestion.Text
	result.Answers = mockQuestion.getAnswers(nAnsw)

	return result
}

func findMockQuestion(id uint) (mockQuestion, error) {
	result := mockQuestion{}
	for _, q := range mockData {
		if id == q.ID {
			result.ID = q.ID
			result.Text = q.Text
			result.Answers = q.Answers
			return result, nil
		}
	}
	return result, fmt.Errorf("can't find question id [%d]", id)
}

func (mq mockQuestion) getAnswers(n int) []Answer {
	// add correct answer
	result := []Answer{mq.getCorrectAnswer()}

	// find incorrect answers without duplicates using a map
	incorrectAnswers := map[uint]Answer{}
	for len(incorrectAnswers) < n-1 {
		incorrectAnswer := mq.getIncorrectAnswer()
		if _, exist := incorrectAnswers[incorrectAnswer.ID]; !exist {
			incorrectAnswers[incorrectAnswer.ID] = incorrectAnswer
		}
	}

	// add incorrect answers
	for _, incorrectAnswer := range incorrectAnswers {
		result = append(result, incorrectAnswer)
	}

	// shuffle answers to avoid the first one being always the correct one
	shuffle(result)
	return result
}

func (mq mockQuestion) getCorrectAnswer() Answer {
	var result Answer
	for _, answer := range mq.Answers {
		if answer.correct {
			result = answer
			break
		}
	}
	return result
}

func (mq mockQuestion) getIncorrectAnswer() Answer {
	var result Answer

	for {
		result = mq.Answers[rnd.Intn(len(mq.Answers))]
		if !result.correct {
			break
		}
	}
	return result
}

func Init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	initMockData()
}

func initMockData() {
	mockData = []mockQuestion{
		mockQuestion{
			ID:   1,
			Text: `What is the capital city of France?`,
			Answers: []Answer{
				Answer{
					ID:      1,
					Text:    `Paris`,
					correct: true,
				},
				Answer{
					ID:      2,
					Text:    `London`,
					correct: false,
				},
				Answer{
					ID:      3,
					Text:    `Rome`,
					correct: false,
				},
				Answer{
					ID:      4,
					Text:    `Berlin`,
					correct: false,
				},
				Answer{
					ID:      5,
					Text:    `Madrid`,
					correct: false,
				},
				Answer{
					ID:      6,
					Text:    `Amsterdam`,
					correct: false,
				},
			},
		},
		mockQuestion{
			ID:   2,
			Text: `Which number is not even?`,
			Answers: []Answer{
				Answer{
					ID:      7,
					Text:    `1`,
					correct: true,
				},
				Answer{
					ID:      8,
					Text:    `2`,
					correct: false,
				},
				Answer{
					ID:      9,
					Text:    `4`,
					correct: false,
				},
				Answer{
					ID:      10,
					Text:    `6`,
					correct: false,
				},
				Answer{
					ID:      11,
					Text:    `8`,
					correct: false,
				},
				Answer{
					ID:      12,
					Text:    `10`,
					correct: false,
				},
			},
		},
		mockQuestion{
			ID:   3,
			Text: `How many bullets has a revolver?`,
			Answers: []Answer{
				Answer{
					ID:      13,
					Text:    `6`,
					correct: true,
				},
				Answer{
					ID:      14,
					Text:    `5`,
					correct: false,
				},
				Answer{
					ID:      15,
					Text:    `4`,
					correct: false,
				},
				Answer{
					ID:      16,
					Text:    `3`,
					correct: false,
				},
				Answer{
					ID:      17,
					Text:    `2`,
					correct: false,
				},
				Answer{
					ID:      18,
					Text:    `1`,
					correct: false,
				},
			},
		},
		mockQuestion{
			ID:   4,
			Text: `In which year America has been discovered?`,
			Answers: []Answer{
				Answer{
					ID:      19,
					Text:    `1492`,
					correct: true,
				},
				Answer{
					ID:      20,
					Text:    `1999`,
					correct: false,
				},
				Answer{
					ID:      21,
					Text:    `1365`,
					correct: false,
				},
				Answer{
					ID:      22,
					Text:    `1564`,
					correct: false,
				},
				Answer{
					ID:      23,
					Text:    `2019`,
					correct: false,
				},
				Answer{
					ID:      24,
					Text:    `1001`,
					correct: false,
				},
			},
		},
		mockQuestion{
			ID:   5,
			Text: `Which of these algorithms is the slowest in average-case?`,
			Answers: []Answer{
				Answer{
					ID:      25,
					Text:    `Bubble sort`,
					correct: true,
				},
				Answer{
					ID:      26,
					Text:    `Merge sort`,
					correct: false,
				},
				Answer{
					ID:      27,
					Text:    `Quicksort`,
					correct: false,
				},
				Answer{
					ID:      28,
					Text:    `Heapsort`,
					correct: false,
				},
				Answer{
					ID:      29,
					Text:    `Insertion sort`,
					correct: false,
				},
			},
		},
		mockQuestion{
			ID:   6,
			Text: `Which of these programming languages is compiled?`,
			Answers: []Answer{
				Answer{
					ID:      30,
					Text:    `Golang`,
					correct: true,
				},
				Answer{
					ID:      31,
					Text:    `Ruby`,
					correct: false,
				},
				Answer{
					ID:      32,
					Text:    `Python`,
					correct: false,
				},
				Answer{
					ID:      33,
					Text:    `PHP`,
					correct: false,
				},
				Answer{
					ID:      34,
					Text:    `JavaScript`,
					correct: false,
				},
				Answer{
					ID:      35,
					Text:    `Perl`,
					correct: false,
				},
			},
		},
		mockQuestion{
			ID:   7,
			Text: `Which of these programming languages is interpreted?`,
			Answers: []Answer{
				Answer{
					ID:      36,
					Text:    `PHP`,
					correct: true,
				},
				Answer{
					ID:      37,
					Text:    `Java`,
					correct: false,
				},
				Answer{
					ID:      38,
					Text:    `Golang`,
					correct: false,
				},
				Answer{
					ID:      39,
					Text:    `C`,
					correct: false,
				},
				Answer{
					ID:      40,
					Text:    `C++`,
					correct: false,
				},
				Answer{
					ID:      41,
					Text:    `Swift`,
					correct: false,
				},
			},
		},
		mockQuestion{
			ID:   8,
			Text: `"I have no special talent. I am only passionately curious". Whom is this quote from?`,
			Answers: []Answer{
				Answer{
					ID:      42,
					Text:    `Albert Einstein`,
					correct: true,
				},
				Answer{
					ID:      43,
					Text:    `Mahatma Gandhi`,
					correct: false,
				},
				Answer{
					ID:      44,
					Text:    `Napoleon Bonaparte`,
					correct: false,
				},
				Answer{
					ID:      45,
					Text:    `Donald Trump`,
					correct: false,
				},
				Answer{
					ID:      46,
					Text:    `Angela Merkel`,
					correct: false,
				},
				Answer{
					ID:      47,
					Text:    `Plato`,
					correct: false,
				},
			},
		},
	}
	mockStats = mockStatistics{
		numOfExecutions:     0,
		totalCorrectAnswers: 0,
		totalQuestions:      0,
		successRatio:        0,
	}
}

// shuffle slice of vals inplace, swapping values to the end of the slice
func shuffle(vals []Answer) {
	// We start at the end of the slice, inserting our random
	// values one at a time.
	for n := len(vals); n > 0; n-- {
		rndIndex := rnd.Intn(n)
		vals[n-1], vals[rndIndex] = vals[rndIndex], vals[n-1]
	}
}
