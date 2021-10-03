package service

import (
	"testing"
)

func TestSolveMdFile(t *testing.T) {
	InitConfig()
	filepath := "./../README.md"
	SolveMdFile(filepath)
}

func TestIninConfig(t *testing.T) {
	InitConfig()
}
