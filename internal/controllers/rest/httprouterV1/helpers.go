package httprouterv1

import (
	"encoding/json"
	"net/http"
)

type wrapJSON map[string]any

type errorJSON struct {
	Errors []string
}

func wrapErrorJSON(w http.ResponseWriter, data []error) error {
	var errors errorJSON

	for _, err := range data {
		if err != nil {
			errors.Errors = append(errors.Errors, err.Error())
		}
	}

	js, err := json.Marshal(errors)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return nil
}

func writeJSON(w http.ResponseWriter, data any) error {
	_, ok := data.(wrapJSON)
	if !ok {
	}

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
