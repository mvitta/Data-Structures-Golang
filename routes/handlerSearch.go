package routes

import (
	"encoding/json"
	"io"
	"net/http"
)

type BodySearchElement struct {
	Value int `json:"value"`
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {

	body, error := io.ReadAll(r.Body)
	if error != nil {
		http.Error(w, "error al leer el cuerpo de la peticion", http.StatusBadRequest)
		return
	}

	var value BodySearchElement
	error_json := json.Unmarshal(body, &value)
	if error_json != nil {
		http.Error(w, "Error en los datos ingresados", http.StatusBadRequest)
		return
	}

	node, _ := l.SearchByValue(value.Value)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(node.NodeToJson())
}
