package utils

import (
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func ToDoRequest(url string) ([]byte, error) {

	req, errReq := http.NewRequest(http.MethodGet, url, nil)
	if errReq != nil {
		log.Println("Error para crear la preticion: ", errReq)
		return nil, errReq
	}
	req.Header.Add("X-RapidAPI-Key", os.Getenv("RAPIDAPI_KEY"))
	req.Header.Add("X-RapidAPI-Host", os.Getenv("RAPIDAPI_HOST"))

	response, errClient := http.DefaultClient.Do(req)
	if errClient != nil {
		slog.Error("Faild Request", errClient)
		return nil, errClient
	}

	defer response.Body.Close()
	body, errBody := io.ReadAll(response.Body)
	if errBody != nil {
		slog.Error("Error reading the body of the request", errBody)
		return nil, errBody
	}

	return body, nil
}
