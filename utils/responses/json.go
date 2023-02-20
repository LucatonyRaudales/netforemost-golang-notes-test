package responses
import (
	"net/http"
	"encoding/json"
	"fmt"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}){
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}