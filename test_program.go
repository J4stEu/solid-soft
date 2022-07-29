package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func main() {
	// Request users posts
	posts, err := getPosts([]int{42, 43, 44, 45})
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println(posts)
}

// getPosts - users posts request
func getPosts(requestPull []int) ([]Post, error) {
	posts := make([]Post, len(requestPull))
	var wg sync.WaitGroup
	c := make(chan error)
	go func() {
		wg.Wait()
		close(c)
	}()
	for k, v := range requestPull {
		wg.Add(1)
		go func(key, value int) {
			defer wg.Done()
			request := fmt.Sprintf("https://winry.khashaev.ru/posts/%v", value)
			// fmt.Println(request)
			postResp, err := postRequest("GET", request)
			if err != nil {
				c <- err
				return
			}
			posts[key] = Post{
				UserId: postResp.UserId,
				Id:     postResp.Id,
				Title:  postResp.Title,
				Body:   postResp.Body,
			}
		}(k, v)
	}

	for err := range c {
		if err != nil {
			return []Post{}, err
		}
	}

	return posts, nil
}

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// postRequest - request a single user post
func postRequest(method, request string) (*Post, error) {
	req, err := http.NewRequest(method, request, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respStruct := &Post{}
	if err = json.Unmarshal(bodyBytes, &respStruct); err != nil {
		return nil, err
	}

	return respStruct, nil
}
