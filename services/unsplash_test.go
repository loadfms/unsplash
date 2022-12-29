package services

import (
	"testing"
)

// TODO: Fix Tests.
func TestGetPictureUrl(t *testing.T) {
	unsplashResponse, err := getPictureUrl()

	if len(unsplashResponse.Results) == 0 || err != nil {
		t.Errorf("Error getting url from image")
	}
}

func TestRun(t *testing.T) {
	err := Run()

	if err != nil {
		t.Errorf(err.Error())
	}
}
