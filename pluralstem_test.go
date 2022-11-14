package pluralstem

import "testing"

func TestInvalidLanguage(t *testing.T) {
	_, err := Stem("hello", "engrish")
	if err == nil {
		t.Errorf("Expected error to be reutnred.")
	}
}

func TestEnglish(t *testing.T) {
	data := []struct {
		In, Out string
	}{
		{"dresses", "dress"}, {"dress", "dress"},
		{"axes", "axe"},
	}

	for _, datum := range data {
		stemmed, err := Stem(datum.In, "english")

		if err != nil {
			t.Errorf("unexpected error occurred: %v", err)
		}

		if stemmed != datum.Out {
			t.Errorf("expected %v but got %v", datum.Out, stemmed)
		}
	}
}
