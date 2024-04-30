package routes

import (
	"encoding/json"
	"io"
	"net/http"
	_ "strconv"

	list "github.com/mvitta/server/linkedList"
)

type BodyValue struct {
	Data int `json:"data"`
}

var l list.LinkedList

func HandleTestAdd(w http.ResponseWriter, r *http.Request) {
	body, error := io.ReadAll(r.Body)
	if error != nil {
		http.Error(w, "error al leer el cuerpo de la peticion", http.StatusBadRequest)
		return
	}

	var dataBody BodyValue
	error_json := json.Unmarshal(body, &dataBody)
	if error_json != nil {
		http.Error(w, "Error en los datos ingresados", http.StatusBadRequest)
		return
	}

	node := l.AddToBeginning(dataBody.Data)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(node.NodeToJson())
}
