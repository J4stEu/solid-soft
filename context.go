package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Use context for request
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	// Imitate request time and cancel request context after time expired
	go func() {
		if err := cancelRequest(); err != nil {
			cancel()
		}
	}()
	// Request with context
	// The Rick and Morty API: https://rickandmortyapi.com
	rmCharacter, err := rmCharacterRequest(ctx, "https://rickandmortyapi.com/api/character/1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rmCharacter)
	}
}

func cancelRequest() error {
	time.Sleep(100 * time.Millisecond)
	return errors.New("cancel request")
}

// Rick and Morty request
func rmCharacterRequest(ctx context.Context, request string) (*RMCharacter, error) {
	req, _ := http.NewRequest(http.MethodGet, request, nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	//select {
	//case <-ctx.Done():
	//	return nil, errors.New("request time out")
	//}
	var character RMCharacter
	err = json.NewDecoder(res.Body).Decode(&character)
	if err != nil {
		return nil, err
	}
	return &character, nil
}

type RMCharacter struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Species string `json:"species"`
	Type    string `json:"type"`
	Gender  string `json:"gender"`
	Origin  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"origin"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Image   string    `json:"image"`
	Episode []string  `json:"episode"`
	URL     string    `json:"url"`
	Created time.Time `json:"created"`
}
