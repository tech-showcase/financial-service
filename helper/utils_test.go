package helper

import (
	"testing"
)

func TestJoinURL(t *testing.T) {
	dummyBaseURL := "http://localhost:8080/"
	dummyPath := "/v1/resources"
	expectedURL := "http://localhost:8080/v1/resources"

	url, err := JoinURL(dummyBaseURL, dummyPath)

	if err != nil {
		t.Fatal("an error has occurred")
	} else if url != expectedURL {
		t.Fatal("unexpected output")
	}
}

func TestGetFrame(t *testing.T) {
	currentFrame := GetFrame(0)

	if currentFrame.Function != "github.com/tech-showcase/financial-service/helper.TestGetFrame" {
		t.Fatal("unexpected output")
	}
}
