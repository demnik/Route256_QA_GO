package testhelpers

import (
	"github.com/brianvoe/gofakeit"
	"testing"
)

func GetRandomPlatform(t *testing.T) string {
	t.Helper()

	return gofakeit.RandString([]string{"Ios",
		"Android",
		"Symbian",
		"BlackBerry",
		"Windows Phone",
		"Windows Mobile",
		"Ubuntu Touch",
		"AMD",
		"FreeDOS"},
	)
}
