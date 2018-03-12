package classifiers

import (
	"strings"
	"testing"

	"resources/plugnplay/models"
)

func CreateTrainingEmails() []models.Email {
	return []models.Email{
		models.Email{Body: "opportunity to earn extra money", Flag: "Spam"},
		models.Email{Body: "druggists blame classy gentry Aladdin", Flag: "Spam"},
		models.Email{Body: "please take a look at this report", Flag: "Ham"},
		models.Email{Body: "lunch at noon?", Flag: "Ham"},
	}
}

func CreateValidationEmails() []models.Email {
	return []models.Email{
		models.Email{Body: "opportunity to earn extra money", Flag: "Spam"},
		models.Email{Body: "druggists blame classy gentry Aladdin", Flag: "Spam"},
		models.Email{Body: "please take a look at this report", Flag: "Ham"},
		models.Email{Body: "lunch at noon?", Flag: "Ham"},
	}
}

func TestLearn(t *testing.T) {
	nbModel := models.SpamHamModel{Classifier: &NBClassifier{}}
	trainingSet := CreateTrainingEmails()
	validationSet := CreateValidationEmails()

	nbModel.Learn(trainingSet)

	for i := 0; i < len(validationSet); i++ {
		input := validationSet[i].Body
		expected := validationSet[i].Flag
		actual := nbModel.Predict(validationSet[i])
		Assert(t, expected, actual, input)
	}
}

func Assert(t *testing.T, expected string, actual string, input string) {
	if actual != expected {
		t.Error(
			"\nFOR:       ", input,
			"\nEXPECTED:  ", expected,
			"\nACTUAL:    ", actual,
		)
	}
}

func AssertList(t *testing.T, expected string, actual []string, input string) {
	for i := 0; i < len(actual); i++ {
		if actual[i] == expected {
			return
		}
	}
	t.Error(
		"\nFOR:       ", input,
		"\nEXPECTED:  ", expected,
		"\nACTUAL:    ", strings.Join(actual, ","),
	)
}
