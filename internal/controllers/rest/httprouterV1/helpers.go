package httprouterv1

import (
	"encoding/json"
	"net/http"
)

type wrapJSON map[string]any

func writeJSON(w http.ResponseWriter, data any) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return nil
}

func internalServerError(w http.ResponseWriter, err error) {
	httpError(w, err, http.StatusInternalServerError)
}

func httpError(w http.ResponseWriter, err error, code int) {
	http.Error(w, err.Error(), code)
}
