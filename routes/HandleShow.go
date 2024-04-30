package routes

import (
	"encoding/json"
	"log/slog"
	"net/http"

	myList "github.com/mvitta/server/linkedList"
)

type ResponseList struct {
	TheList myList.LinkedList `json:"theList"`
}

func (theList *ResponseList) Json() []byte {
	json, errorJson := json.Marshal(theList)
	if errorJson != nil {
		slog.Error(errorJson.Error())
		return nil
	}
	return json
}

func HandleShowElements(w http.ResponseWriter, r *http.Request) {

	response := ResponseList{
		TheList: l,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response.Json())

}
