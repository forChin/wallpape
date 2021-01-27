package main

import (
	"fmt"
	"strings"
)

type photo struct {
	Width  int               `json:"width"`
	Height int               `json:"height"`
	Src    map[string]string `json:"src"`
}

type searchResult struct {
	Error        string  `json:"error"`
	Status       int     `json:"status"`
	Code         string  `json:"code"`
	TotalResults int     `json:"total_results"`
	Photos       []photo `json:"photos"`
}

func (sr *searchResult) Err() error {
	if sr.Status != 0 && sr.Code != "" {
		return fmt.Errorf("could not find image: %s", strings.ToLower(sr.Code))
	}

	if sr.Error != "" {
		return fmt.Errorf("please, try again later: %v", strings.ToLower(sr.Error))
	}

	return nil
}
