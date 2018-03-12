package classifiers

import (
	"sort"
	"strings"

	"resources/plugnplay/models"

	"github.com/jbrukh/bayesian"
)

// Step 4 - Implement Classifier(s). Per the Contracts Terms.
// You got mail! - Is it Spam or Ham? (Model's Brain/Classifier)
type NBClassifier struct {
	classifier *bayesian.Classifier
	output     []bayesian.Class
}

func (c *NBClassifier) Learn(emails []models.Email) {
	c.output = distinctFlags(emails)
	c.classifier = bayesian.NewClassifierTfIdf(c.output...)
	for i := 0; i < len(emails); i++ {
		c.classifier.Learn(strings.Split(emails[i].Body, " "), bayesian.Class(emails[i].Flag))
	}
	c.classifier.ConvertTermsFreqToTfIdf()
}

func (c *NBClassifier) Predict(email models.Email) string {
	scores, _, _ := c.classifier.LogScores(strings.Split(email.Body, " "))
	results := models.Results{}
	for i := 0; i < len(scores); i++ {
		results = append(results, models.Result{ID: i, Score: scores[i]})
	}

	sort.Sort(sort.Reverse(results))

	flags := []string{}
	for i := 0; i < len(results); i++ {
		flags = append(flags, string(c.output[results[i].ID]))
	}
	return flags[0]
}

func distinctFlags(emails []models.Email) []bayesian.Class {
	result := []bayesian.Class{}
	j := 0
	for i := 0; i < len(emails); i++ {
		for j = 0; j < len(result); j++ {
			if emails[i].Flag == string(result[j]) {
				break
			}
		}
		if j == len(result) {
			result = append(result, bayesian.Class(emails[i].Flag))
		}
	}
	return result
}
