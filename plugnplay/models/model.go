package models

// Step 1 - Define the Structure of Input Data
// You got mail!
type Email struct {
	Author string
	Body   string
	Flag   string //Spam/Ham
}

// Step 2 - Define a ML Classifier Contract
// Binary Classifier Interface - Examples SVM, NN, NB
type Classifier interface {
	Learn(emails []Email)
	Predict(email Email) string
}

// Step 3 - Create a Model with a "Plug & Play" Classifer Field
// You got mail! - Is it Spam or Ham? (Model Example)
type SpamHamModel struct {
	Classifier Classifier
}

func (model *SpamHamModel) Learn(emails []Email) {
	model.Classifier.Learn(emails)
}

func (model *SpamHamModel) Predict(email Email) string {
	return model.Classifier.Predict(email)
}

// Step 5 - (Optional)
// Create a Model Interface - Examples: Spam/Ham, Sentiment Analysis
type Model interface {
	Learn(emails []Email)
	Predict(email Email) string
}
