package base

import (
	"errors"

	"golang.org/x/exp/slices"
)

func ValidateImageSize(size string) error {
	if size != "" && !slices.Contains([]string{"256x256", "512x512", "1024x1024"}, size) {
		return errors.New("size must be one of 256x256, 512x512 or 1024x1024")
	}
	return nil
}

func ValidateImageFormat(format string) error {
	if format != "" && !slices.Contains([]string{"url", "b64_json"}, format) {
		return errors.New("response_format must be one of url or b64_json")
	}
	return nil
}

func ValidatePrompt(prompt interface{}) error {
	switch prompt.(type) {
	case string, []string, []int64, [][]int64, []int, [][]int, nil:
		return nil
	}
	return errors.New("invalid prompt type, must be string, []string, []int64 or [][]int64, check their api documentation for more details.")
}

func ValidateInput(prompt interface{}) error {
	switch prompt.(type) {
	case string, []string, []int64, [][]int64, []int, [][]int, nil:
		return nil
	}
	return errors.New("invalid input type, must be string, []string, []int64 or [][]int64, check their api documentation for more details.")
}
