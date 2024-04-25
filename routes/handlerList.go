package routes

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	_ "strconv"

	list "github.com/mvitta/server/linkedList"
)

type BodyValue struct {
	ValueLast  int `json:"valueLast"`
	ValueFirst int `json:"valueFirst"`
}

type ResponseList struct {
	TheList list.LinkedList `json:"theList"`
}

func (theList *ResponseList) Json() []byte {
	json, errorJson := json.Marshal(theList)
	if errorJson != nil {
		slog.Error(errorJson.Error())
		return nil
	}
	return json
}

var l list.LinkedList

func HandleTestList(w http.ResponseWriter, r *http.Request) {
	body, error := io.ReadAll(r.Body)
	if error != nil {
		http.Error(w, "error al leer el cuerpo de la peticion", http.StatusBadRequest)
		return
	}

	var bodyValue BodyValue
	error_json := json.Unmarshal(body, &bodyValue)
	if error_json != nil {
		http.Error(w, "Error en los datos ingresados", http.StatusBadRequest)
		return
	}

	// value, error_conver := strconv.Atoi(bodyValue.Value)
	// if error_conver != nil {
	// 	http.Error(w, "el valor no es un numero valido", http.StatusBadRequest)
	// 	return
	// }

	// if the numbers are negatives, it does not execute the corresponding method

	if bodyValue.ValueFirst > 0 {
		l.AddToEnd(bodyValue.ValueLast)
	}

	if bodyValue.ValueLast > 0 {
		l.AddToBeginning(bodyValue.ValueFirst)
	}

	myResponse := ResponseList{
		TheList: l,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(myResponse.Json()))
}
