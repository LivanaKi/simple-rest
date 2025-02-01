package helper

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Users/natza/simple-rest/data/response"
)

func ReadRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)

	if err != nil {
		log.Fatal(err)
	}
}

func WriteResponse(writer http.ResponseWriter, response response.WebResponse) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(response.Code)
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}
