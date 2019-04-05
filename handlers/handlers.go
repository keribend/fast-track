package handlers

import (
	"fast-track/models"
	"net/http"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func GetQuestionnaire() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetQuestionnaire(5))
	}
}

func AnswerQuestionnaire() echo.HandlerFunc {
	return func(c echo.Context) error {
		selectedAnswers := new(models.QuestionnaireAnswers)

		c.Bind(selectedAnswers)

		questionnaireResult, err := models.AnswerQuestionnaire(selectedAnswers)

		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"questionnaireResult": questionnaireResult,
			})
		}

		return err
	}
}
