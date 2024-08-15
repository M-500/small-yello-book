package utils

import "testing"

func TestGenerateRandomNumber(t *testing.T) {
	t.Log(GenerateRandomNumber())
	t.Log(GenerateRandomNumberStr())
}
