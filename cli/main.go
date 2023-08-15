package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func get() {

	fmt.Println("Performing GET request...!")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		log.Fatal("Error ", err)
	}
	defer resp.Body.Close()

	Body, _ := io.ReadAll(resp.Body)

	fmt.Println("API Response as String:\n" + string(Body))

	var todoStruct Todo
	json.Unmarshal(Body, &todoStruct)
	fmt.Printf("API Response as struct %+v\n", todoStruct)

}

func post() {
	fmt.Println("Performing POST request/....!")

	PostReq := Todo{UserID: 2, ID: 3445, Title: "Hello master", Completed: true}

	req, _ := json.Marshal(PostReq)

	resp, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json; charset=utf-8", bytes.NewBuffer(req))
	if err != nil {
		log.Fatal("Error")
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("API response for POSt call %+v\n", string(body))
	defer resp.Body.Close()
}

func put() {
	fmt.Println("Performing Put request/....!")

	PostReq := Todo{UserID: 2, ID: 5220, Title: "Hello master", Completed: true}

	req, _ := json.Marshal(PostReq)

	reqst, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos", bytes.NewBuffer(req))
	if err != nil {
		log.Fatal("Error")
	}
	client := &http.Client{}
	resp, err := client.Do(reqst)

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("API response for Put call %+v\n", string(body))
	defer resp.Body.Close()
}

func delete() {
	fmt.Println("Performing Delete request/....!")

	PostReq := Todo{UserID: 2, ID: 3445, Title: "Hello master", Completed: true}

	req, _ := json.Marshal(PostReq)

	reqst, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/todos", bytes.NewBuffer(req))
	if err != nil {
		log.Fatal("Error")
	}

	client := &http.Client{}
	resp, err := client.Do(reqst)
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("API response for Delete call %+v\n", string(body))
	defer resp.Body.Close()
}

func main() {
	get()
	post()
	put()
	post()
	delete()
	get()
}
